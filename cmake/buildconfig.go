package cmake

// BuildConfig describe the configuration of a cmake build
type BuildConfig struct {
	SourceDir string
	BuildDir  string
	Generator string
	Configure []string
	Env       []string
}

func (bc *BuildConfig) SocketName() string {
	return bc.BuildDir + ".sock"
}
