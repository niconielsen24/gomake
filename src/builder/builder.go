package builder

type GitRepo int

const (
	INIT_GIT_REPO = iota
	NO_GIT_REPO
)

type BuilderConfig struct {
	Lang        string
	ProjectName string
	GitRepo     GitRepo
	DirTree     DirTree
	Techs       []string
}

func NewBuilderConfig() BuilderConfig {
	return BuilderConfig{}
}

type GoMakeBuilder interface {
	// This method just handle all the project building
	InitProj(bc *BuilderConfig) (string, error)
	GetDependencies(bc *BuilderConfig) (string, error)
	BuildDirs(bc *BuilderConfig) error
}
