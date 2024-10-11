package common

import (
	"bufio"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

type WrapperActionParams interface {
	Opts() map[string][]string
	OptsString() string
	OptsStringSlice() []string
}

type WrapperAction struct {
	Cmd    *exec.Cmd
	Dir    string
	out    *WrapperOutput
	logs   *OutputLog
	action string
	bin    *WrapperCli
	params WrapperActionParams
}

// Initialise prepares and initializes a WrapperAction by setting up the command
// arguments, creating the command, setting the working directory if specified,
// and initializing the output.
func (a *WrapperAction) Initialise() *WrapperAction {
	// Prepare the command arguments
	args := append([]string{a.action}, a.params.OptsStringSlice()...)

	// Create the command
	a.Cmd = exec.Command(a.bin.path, args...)

	// Set the working directory if specified
	if a.Dir != "" {
		a.Cmd.Dir = a.Dir
	}

	// Initialize the output
	a.out = &WrapperOutput{}

	return a
}

// Run executes the Wrapper command associated with the WrapperAction instance.
// It starts the command and returns any error encountered during the start process.
func (a *WrapperAction) Run() (err error) {
	return a.Cmd.Run()
}

// Initializes the logging mechanism for the WrapperAction.
// It sets up the stdout and stderr pipes for the command execution and
// starts goroutines to capture and log the output.
func (a *WrapperAction) InitLogger(log *OutputLog) (err error) {
	a.logs = log

	// Configure stdout capture
	if a.out.Stdout, err = a.Cmd.StdoutPipe(); err != nil {
		return
	}
	scannerStdout := bufio.NewScanner(a.out.Stdout)
	go func() {
		for scannerStdout.Scan() {
			// Print the stdout output to the console and store it in the log
			fmt.Print(
				a.logs.Stdout(scannerStdout.Text()).String() + "\n",
			)
		}
	}()

	// Configure stderr capture
	if a.out.Stderr, err = a.Cmd.StderrPipe(); err != nil {
		return
	}
	scannerStderr := bufio.NewScanner(a.out.Stderr)
	go func() {
		for scannerStderr.Scan() {
			// Print the stderr output to the console and store it in the log
			fmt.Print(
				a.logs.Stderr(scannerStderr.Text()).String() + "\n",
			)
		}
	}()

	return
}

// Processes the options from WrapperActionParams and returns them as a sorted slice of strings.
// It sorts the keys of the options map, then iterates over each key and its associated values.
// For each key-value pair, it constructs a string in the format "-key=value" and appends it to the output slice.
// If the key is "var", it appends the key and value as separate elements to handle Wrapper variable syntax.
// The resulting slice contains all options in a consistent and sorted order.
func extractOptsStringSlice(p WrapperActionParams) []string {
	opts := p.Opts()
	keys := mapStringSliceKeys(opts)
	sort.Strings(keys)

	var outputs []string
	for _, key := range keys {
		values := opts[key]
		sort.Strings(values)
		for _, val := range values {
			output := "-" + key
			if val != "" {
				if key == "var" {
					outputs = append(outputs, output, "'"+val+"'")
					continue
				}
				output += "=" + val
			}
			outputs = append(outputs, output)
		}
	}
	return outputs
}

// Takes a WrapperActionParams object and returns a string
// representation of its options. It does this by joining the elements of the
// slice returned by extractOptsStringSlice with a space character.
func extractOptsString(p WrapperActionParams) string {
	optionsSlice := extractOptsStringSlice(p)
	options := strings.Join(optionsSlice, " ")
	return options
}

// Takes a map with string keys and slice of strings as values,
// and returns a slice containing all the keys from the input map.
func mapStringSliceKeys(s map[string][]string) []string {
	keys := make([]string, 0, len(s)) // Initialize slice with capacity of map length

	for k := range s {
		keys = append(keys, k) // Append each key to the slice
	}

	return keys
}