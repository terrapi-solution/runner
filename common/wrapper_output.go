package common

import (
	"fmt"
	"io"
)

const (
	STDERR = "stderr"
	STDOUT = "stdout"
)

type WrapperOutput struct {
	Stderr io.ReadCloser
	Stdout io.ReadCloser
}

type WrapperOutputLog struct {
	Entries []*WrapperOutputLogEntry `json:"entries"`
}

type WrapperOutputLogEntry struct {
	Type    string   `json:"type"`
	Content string   `json:"content"`
	Tags    []string `json:"prefix"`
}

// Creates and returns a new instance of OutputLog.
func NewOutputLogs() *WrapperOutputLog {
	return &WrapperOutputLog{
		Entries: make([]*WrapperOutputLogEntry, 0),
	}
}

// Creates a new OutputLogEntry with the specified message.
func (ol *WrapperOutputLog) Stdout(message string) *WrapperOutputLogEntry {
	return ol.StdoutWithTags(message, []string{"tf"})
}

// Creates a new OutputLogEntry with the specified message and tags.
func (ol *WrapperOutputLog) StdoutWithTags(message string, tags []string) *WrapperOutputLogEntry {
	return ol.Append(&WrapperOutputLogEntry{
		Type:    STDOUT,
		Content: message,
		Tags:    tags,
	})
}

// Creates a new OutputLogEntry with the specified error message.
func (ol *WrapperOutputLog) Stderr(message string) *WrapperOutputLogEntry {
	return ol.StderrWithTags(message, []string{"tf"})
}

// Creates a new OutputLogEntry with the specified error message and tags.
func (ol *WrapperOutputLog) StderrWithTags(message string, tags []string) *WrapperOutputLogEntry {
	return ol.Append(&WrapperOutputLogEntry{
		Type:    STDERR,
		Content: message,
		Tags:    tags,
	})
}

// Appends the specified OutputLogEntry to the OutputLog.
func (ol *WrapperOutputLog) Append(ole *WrapperOutputLogEntry) *WrapperOutputLogEntry {
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
