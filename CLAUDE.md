# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

`dcd` is a Cobra-based Go CLI for managing resources (workspace, application, domain, env, volume) on the DCD platform, talking to a REST API and, for interactive `exec`, a websocket endpoint.

## Commands

```bash
# Build the binary (module path: github.com/dolong2/dcd-cli)
go build -o dcd .

# Run without building
go run . <subcommand>

# Vet
go vet ./...
```

There are no `_test.go` files in the repo currently, so there is no test command to run.

### baseUrl is injected at build time, not hardcoded

`api.baseUrl` (in `api/api_client.go`) and `websocket.baseUrl` (in `websocket/websocket_client.go`) are unexported package-level vars that default to `""`. They are only ever set via `-ldflags -X`, as done in `install.sh`:

```bash
go build -ldflags="-X github.com/dolong2/dcd-cli/api.baseUrl=https://dcd-api.dolong2.co.kr -X github.com/dolong2/dcd-cli/websocket.baseUrl=wss://dcd-api.dolong2.co.kr" -o dcd
```

A plain `go build` produces a binary that cannot reach any backend — all HTTP/websocket calls will hit a relative/empty host. When testing against a real or local backend, pass matching `-X` flags for both vars.

### Cross-platform build gotcha (`api/exec/util/filekey_*.go`)

`getFileKey` has per-OS implementations selected by build tag. Go's implicit filename-based build constraint (`_linux`, `_darwin`, `_windows` suffixes) is ANDed with any explicit `//go:build` line in the same file — a file named `foo_linux.go` is restricted to `linux` even if its `//go:build` comment says `linux || darwin`. This is why the Unix implementation lives in `filekey_unix.go` (not `filekey_linux.go`) with `//go:build linux || darwin`; renaming it back to a `_linux.go`/`_darwin.go`-suffixed name will silently break the other OS.

## Architecture

**`cmd/`** — Cobra command definitions. Handlers are thin: parse flags/args, call into `api/exec`, and wrap errors as `cmdError.NewCmdError(code, msg)` (`cmd/err`). `root.go`'s `Execute()` unwraps a returned `*cmdError.CmdError` and calls `os.Exit(cmdErr.Code)`, or exits 1 for any other error. Resource-scoped subcommands (e.g. `application`) use `PersistentPreRun` to stash a positional arg (e.g. `applicationId`) into `cmd.Context()` via a private `contextKey`, which child commands (e.g. `exec.go`) read back out.

**`api/`** — `api_client.go` is a minimal generic HTTP client (`SendGet/Post/Patch/Put/Delete`) that prefixes `targetUrl` with the ldflags-injected `baseUrl`, sets `Authorization`/other headers, and turns non-2xx responses into `httpErr` (`api/err`).

**`api/exec/`** — one file per API operation (`create.go`, `deploy.go`, `get_application.go`, ...), all in `package exec`. Each builds a request, calls `api.Send*`, and unmarshals into a typed response. Sub-packages:
- `request/`, `response/` — wire DTOs per resource.
- `template/` — structs (`WorkspaceTemplate`, `ApplicationTemplate`, ...) that parse user-authored YAML/JSON resource definitions (see `example/`). Each has `validateMetadata()` and `ToRequest()` to convert into an `api/exec/request` DTO. `create.go`'s `create()` dispatches on `metadata.resourceType` (`WORKSPACE`/`APPLICATION`/`ENV`/`DOMAIN`/`VOLUME`) read via the shared `ParsingMetaData` struct, parsed once to sniff the type before parsing again into the concrete template.
- `util/` — `resource_mapper.go` maps a local template file path to its created `resourceId` in `dcd-info/resource-mapping-info.json`, keyed by an OS-level file identity (`getFileKey`, see the build-tag gotcha above), not the path string, so renamed/moved files still resolve. `get_workspace_info.go` and `resolve_file_ext.go` are similar local-state/format helpers.

**`cmd/util/`** — CLI-side helpers distinct from `api/exec/util`: `get_workspace_info.go` resolves the active workspace id (from `--workspace` flag or the saved `dcd-info/workspace-info.json`), `save_workspace_info.go` persists it (written by `cmd/use.go`), `get_resource_type.go`/`print_resource.go` back the generic `get` command.

**`cmd/resource/resource_type.go`** — central registry of resource type names/aliases (e.g. `workspace`/`workspaces`/`ws`) used by `dcd get <type>` to validate and normalize the CLI arg.

**`dcd-info/`** — gitignored local state directory created next to wherever `dcd` is run:
- `token-info.json` — access/refresh tokens; `api/exec/util.go`'s `GetAccessToken()` reads it, and auto-calls `ReissueToken` + retries (via `goto`) if the access token is expired, erroring out only if the refresh token has also expired.
- `workspace-info.json` — the workspace set by `dcd use <workspaceId>`.
- `resource-mapping-info.json` — file-path → resourceId map (see `resource_mapper.go` above).

**`websocket/`** — used only by `dcd application exec --ws`, for an interactive shell into a running application (`/application/exec?applicationId=...`). Has its own independently-injected `baseUrl` (see build note above). `cmd/exec.go` runs separate goroutines for the read loop, the stdin read (nested one level deeper since `bufio.Reader.ReadLine` blocks), and the send loop, coordinated over channels with an `interrupt`/`errChan` select to shut down cleanly on Ctrl-C or a connection error.

**`example/`** — sample `json`/`yml` resource definitions for every resource type, useful as reference for the `template/` structs' expected shape and for manual testing of `dcd create -f <path>`.
