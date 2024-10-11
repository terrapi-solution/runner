package wrapper

type InitParams struct {
	// Backend indicates whether the backend service should be initialized.
	Backend *bool

	// Specifies the configuration settings for the backend service.
	BackendConfig string

	// Indicates whether the output should be formatted as JSON.
	Json bool
}

// Creates and returns a new instance of InitParams.
func NewInitParams() *InitParams {
	return &InitParams{
		Json: true,
	}
}

// This function generates a map of  initialization options based on the
// parameters provided in the InitParams struct.
func (p *InitParams) Opts() map[string][]string {
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
	if p.Json {
		opts["json"] = []string{""}
	}

	// Disable color output, because product is not a terminal
	opts["no-color"] = []string{""}

	return opts
}

// Returns a string representation of the  initialization options
// based on the parameters provided in the InitParams struct.
func (p *InitParams) OptsString() string {
	return extractOptsString(p)
}

// Converts the InitParams options into a slice of strings.
// This function extracts the options from the InitParams struct and returns them
// as a slice of strings, which can be used for further processing or passing to other functions.
func (p *InitParams) OptsStringSlice() []string {
	return extractOptsStringSlice(p)
}
