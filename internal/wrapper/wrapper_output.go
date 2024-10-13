package wrapper

import (
	"fmt"
	"io"
)

const (
	STDERR = "stderr"
	STDOUT = "stdout"
)

type Output struct {
	Stderr io.ReadCloser
	Stdout io.ReadCloser
}

type OutputLog struct {
	Entries []*OutputLogEntry `json:"entries"`
}

type OutputLogEntry struct {
	Type    string   `json:"type"`
	Content string   `json:"content"`
	Tags    []string `json:"prefix"`
}

// Creates and returns a new instance of OutputLog.
func NewOutputLogs() *OutputLog {
	return &OutputLog{
		Entries: make([]*OutputLogEntry, 0),
	}
}

// Creates a new OutputLogEntry with the specified message.
func (ol *OutputLog) Stdout(message string) *OutputLogEntry {
	return ol.StdoutWithTags(message, []string{"tf"})
}

// Creates a new OutputLogEntry with the specified message and tags.
func (ol *OutputLog) StdoutWithTags(message string, tags []string) *OutputLogEntry {
	return ol.Append(&OutputLogEntry{
		Type:    STDOUT,
		Content: message,
		Tags:    tags,
	})
}

// Creates a new OutputLogEntry with the specified error message.
func (ol *OutputLog) Stderr(message string) *OutputLogEntry {
	return ol.StderrWithTags(message, []string{"tf"})
}

// Creates a new OutputLogEntry with the specified error message and tags.
func (ol *OutputLog) StderrWithTags(message string, tags []string) *OutputLogEntry {
	return ol.Append(&OutputLogEntry{
		Type:    STDERR,
		Content: message,
		Tags:    tags,
	})
}

// Appends the specified OutputLogEntry to the OutputLog.
func (ol *OutputLog) Append(ole *OutputLogEntry) *OutputLogEntry {
	ol.Entries = append(ol.Entries, ole)
	return ole
}

// Returns a string representation of the OutputLog.
func (ol *OutputLog) String() (output string) {
	output = ""
	for _, entry := range ol.Entries {
		output = output + entry.String() + "\n"
	}
	return
}

// Returns a string representation of the OutputLogEntry.
func (ole *OutputLogEntry) String() string {
	tags := ""
	for _, tag := range ole.Tags {
		tags = tags + "[" + tag + "]"
	}

	var prefix string
	if ole.Type == STDERR {
		prefix = "[stderr]"
	} else if ole.Type == STDOUT {
		prefix = "[stdout]"
	}

	return fmt.Sprintf("%s%s %s", prefix, tags, ole.Content)
}
