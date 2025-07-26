package resource

type Type string
type definedType []string

var (
	Workspace       = definedType{"WORKSPACE", "workspace", "workspaces", "ws"}
	Application     = definedType{"APPLICATION", "application", "applications", "app"}
	Env             = definedType{"ENV", "environment", "environments", "env"}
	GlobalEnv       = definedType{"GLOBAL_ENV", "global_env", "global_envs", "ge"}
	Domain          = definedType{"DOMAIN", "domain", "domains", "dom"}
	ApplicationType = definedType{"TYPE", "type", "types", "ts"}
)

var allResourceTypes = []definedType{Workspace, Application, Env, GlobalEnv, Domain, ApplicationType}

func (rt Type) IsValid() bool {
	for _, definedResource := range allResourceTypes {
		for _, availableStr := range definedResource {
			if rt == Type(availableStr) {
				return true
			}
		}
	}

	return false
}

func (rt Type) IsEqual(other definedType) bool {
	for _, availableStr := range other {
		if rt == Type(availableStr) {
			return true
		}
	}

	return false
}
