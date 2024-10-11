package common

type WrapperCli struct {
	// path represents the file system path used within the wrapper package.
	path string

	// Represents the directory path where the application will perform its operations.
	workingDirectory string
}

// Creates a new instance of WrapperCli with the specified binary path.
func NewWrapper(binPath string) *WrapperCli {
	return &WrapperCli{
		path: binPath,
	}
}

// Initializes a new Wrapper action with the "init" command.
func (t *WrapperCli) Init(params *WrapperInitParams) *WrapperAction {
	return &WrapperAction{
		action: "init",
		bin:    t,
		params: params,
		Dir:    t.workingDirectory,
	}
}

// Sets the working directory for the Wrapper CLI client.
func (client *WrapperCli) SetWorkingDirectory(workingDirectory string) *WrapperCli {
	client.workingDirectory = workingDirectory
	return client
}
