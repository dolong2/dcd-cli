package resource

type Type string
type definedType []string

var (
	Workspace       = definedType{"workspace", "workspaces", "ws"}
	Application     = definedType{"application", "applications", "app"}
	Env             = definedType{"environment", "environments", "env"}
	Domain          = definedType{"domain", "domains", "dom"}
	ApplicationType = definedType{"type", "types", "ts"}
	VOLUME          = definedType{"volume", "volumes", "vol"}
)

var allResourceTypes = []definedType{Workspace, Application, Env, Domain, ApplicationType, VOLUME}

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
