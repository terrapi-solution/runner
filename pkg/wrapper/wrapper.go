package wrapper

type TerraformCli struct {
	// path represents the file system path used within the wrapper package.
	path string

	// Represents the directory path where the application will perform its operations.
	workingDirectory string
}

// Creates a new instance of TerraformCli using the default binary path "terraform".
func New() *TerraformCli {
	return NewWithBinPath("terraform")
}

// Creates a new instance of TerraformCli with the specified binary path.
func NewWithBinPath(binPath string) *TerraformCli {
	return &TerraformCli{
		path: binPath,
	}
}

// Initializes a new Terraform action with the "init" command.
func (t *TerraformCli) Init(params *TerraformInitParams) *TerraformAction {
	return &TerraformAction{
		action: "init",
		bin:    t,
		params: params,
		Dir:    t.workingDirectory,
	}
}

// Sets the working directory for the Terraform CLI client.
func (client *TerraformCli) SetWorkingDirectory(workingDirectory string) *TerraformCli {
	client.workingDirectory = workingDirectory
	return client
}
