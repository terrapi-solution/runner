package common

type WrapperInitParams struct {
	// Backend indicates whether the backend service should be initialized.
	Backend *bool

	// Specifies the configuration settings for the backend service.
	BackendConfig string

	// Indicates whether the output should be formatted as JSON.
	Json bool
}

// Creates and returns a new instance of WrapperInitParams.
func NewInitParams() *WrapperInitParams {
	return &WrapperInitParams{
		Json: true,
	}
}

// This function generates a map of Wrapper initialization options based on the
// parameters provided in the WrapperInitParams struct.
func (p *WrapperInitParams) Opts() map[string][]string {
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

// Returns a string representation of the Wrapper initialization options
// based on the parameters provided in the WrapperInitParams struct.
func (p *WrapperInitParams) OptsString() string {
	return extractOptsString(p)
}

// Converts the WrapperInitParams options into a slice of strings.
// This function extracts the options from the WrapperInitParams struct and returns them
// as a slice of strings, which can be used for further processing or passing to other functions.
func (p *WrapperInitParams) OptsStringSlice() []string {
	return extractOptsStringSlice(p)
}
