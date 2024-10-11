package common

import (
	"fmt"
	"time"
)

type GitCloneConfig struct {
	URL     string `json:"url" validate:"required,url"`
	Address string `json:"address" validate:"required"`
	Branch  string `json:"branch" validate:"required"`
	Path    string `json:"path" validate:"required"`
}

type gitClone struct {
	GitCloneConfig
	retryHelper

	client *gitClient
}

// Creates a new instance of gitClone.
func NewGitExtractor(cfg GitCloneConfig) *gitClone {
	return &gitClone{
		retryHelper:    retryHelper{Retry: 3, RetryTime: 2 * time.Second},
		GitCloneConfig: cfg,
		client:         &gitClient{},
	}
}

// Attempts to execute clone action with retries.
func (c *gitClone) Execute() error {
	err := c.doRetry(c.clone)
	if err != nil {
		return fmt.Errorf("failed to clone repository from URL %s to path %s: %w", c.URL, c.Path, err)
	}
	return nil
}

// Initializes and returns a gitClient instance.
func (c *gitClone) getClient() *gitClient {
	if c.client == nil {
		c.client = NewGitClient()
		c.client.PrepareClient(c.Address, c.Branch)
	}

	return c.client
}

// Attempts to clone a Git repository to the specified path.
// If the cloning process encounters an error, it returns a retryableErr
// containing the original error. Otherwise, it returns nil.
func (c *gitClone) clone(_ int) error {
	_, err := c.getClient().Clone(c.Path)
	if err != nil {
		return retryableErr{err: err}
	}
	return nil
}
