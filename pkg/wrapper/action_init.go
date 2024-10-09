package wrapper

type TerraformInitParams struct {
	// Backend indicates whether the backend service should be initialized.
	Backend *bool

	// Specifies the configuration settings for the backend service.
	BackendConfig string

	// Indicates whether the output should be formatted as JSON.
	Json bool
}

// Creates and returns a new instance of TerraformInitParams.
func NewInitParams() *TerraformInitParams {
	return &TerraformInitParams{
		Json: true,
	}
}

// This function generates a map of Terraform initialization options based on the
// parameters provided in the TerraformInitParams struct.
func (p *TerraformInitParams) Opts() map[string][]string {
	opts := make(map[string][]string)

	// Backend indicates whether the backend service should be initialized.
	if p.Backend != nil && !*p.Backend {
		opts["backend"] = []string{"false"}
	}

	// Specifies the configuration settings for the backend service.
	if p.BackendConfig != "" {
		opts["backend-config"] = []string{p.BackendConfig}
	}

	// Indicates whether the output should be formatted as JSON.
	if p.Backend != nil && *p.Backend {
		opts["json"] = []string{"true"}
	}

	return opts
}

// Returns a string representation of the Terraform initialization options
// based on the parameters provided in the TerraformInitParams struct.
func (p *TerraformInitParams) OptsString() string {
	return extractOptsString(p)
}

// Converts the TerraformInitParams options into a slice of strings.
// This function extracts the options from the TerraformInitParams struct and returns them
// as a slice of strings, which can be used for further processing or passing to other functions.
func (p *TerraformInitParams) OptsStringSlice() []string {
	return extractOptsStringSlice(p)
}
