package wrapper

import (
	"bufio"
	"context"
	"fmt"
	activity "github.com/terrapi-solution/protocol/activity/v1"
	"github.com/terrapi-solution/runner/internal/client"
	"os/exec"
	"sort"
	"strings"
)

type ActionParams interface {
	Opts() map[string][]string
	OptsString() string
	OptsStringSlice() []string
}

type Action struct {
	Cmd    *exec.Cmd
	Dir    string
	out    *Output
	logs   *OutputLog
	action string
	bin    *Cli
	params ActionParams
}

// Initialise prepares and initializes a Action by setting up the cmd
// arguments, creating the cmd, setting the working directory if specified,
// and initializing the output.
func (a *Action) Initialise() *Action {
	// Prepare the cmd arguments
	args := append([]string{a.action}, a.params.OptsStringSlice()...)

	// Create the cmd
	a.Cmd = exec.Command(a.bin.path, args...)

	// Set the working directory if specified
	if a.Dir != "" {
		a.Cmd.Dir = a.Dir
	}

	// Initialize the output
	a.out = &Output{}

	return a
}

// Run executes the  cmd associated with the Action instance.
// It starts the cmd and returns any error encountered during the start process.
func (a *Action) Run() (err error) {
	return a.Cmd.Run()
}

// Initializes the logging mechanism for the Action.
// It sets up the stdout and stderr pipes for the cmd execution and
// starts goroutines to capture and log the output.
func (a *Action) InitLogger(log *OutputLog) (err error) {
	a.logs = log

	rpcClient := client.NewClient()
	// Configure stdout capture
	if a.out.Stdout, err = a.Cmd.StdoutPipe(); err != nil {
		return
	}
	scannerStdout := bufio.NewScanner(a.out.Stdout)
	go func() {
		for scannerStdout.Scan() {
			request := &activity.InsertRequest{
				Deployment: int32(1),
				Pointer:    activity.Pointer_POINTER_STDOUT,
				Message:    scannerStdout.Text(),
			}
			rpcClient.Activity.Insert(context.Background(), request)
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

// Processes the options from ActionParams and returns them as a sorted slice of strings.
// It sorts the keys of the options map, then iterates over each key and its associated values.
// For each key-value pair, it constructs a string in the format "-key=value" and appends it to the output slice.
// If the key is "var", it appends the key and value as separate elements to handle  variable syntax.
// The resulting slice contains all options in a consistent and sorted order.
func extractOptsStringSlice(p ActionParams) []string {
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

// Takes a ActionParams object and returns a string
// representation of its options. It does this by joining the elements of the
// slice returned by extractOptsStringSlice with a space character.
func extractOptsString(p ActionParams) string {
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
