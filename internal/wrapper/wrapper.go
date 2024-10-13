package wrapper

type Cli struct {
	path             string
	workingDirectory string
}

// Creates a new instance of Cli with the specified binary path.
func New(binPath string) *Cli {
	return &Cli{
		path: binPath,
	}
}

// Initializes a new  action with the "init" command.
func (t *Cli) Init(params *InitParams) *Action {
	return &Action{
		action: "init",
		bin:    t,
		params: params,
		Dir:    t.workingDirectory,
	}
}

// Sets the working directory for the  CLI client.
func (client *Cli) SetWorkingDirectory(workingDirectory string) *Cli {
	client.workingDirectory = workingDirectory
	return client
}
