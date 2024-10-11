package git

import (
	"fmt"
	"github.com/thomas-illiet/terrapi-runner/pkg/helper"
	"time"
)

type CloneConfig struct {
	URL     string `json:"url" validate:"required,url"`
	Address string `json:"address" validate:"required"`
	Branch  string `json:"branch" validate:"required"`
	Path    string `json:"path" validate:"required"`
}

type Clone struct {
	CloneConfig
	helper.RetryHelper

	client *Client
}

// Creates a new instance of gitClone.
func NewGitExtractor(cfg CloneConfig) *Clone {
	return &Clone{
		RetryHelper: helper.RetryHelper{Retry: 3, RetryTime: 2 * time.Second},
		CloneConfig: cfg,
		client:      &Client{},
	}
}

// Attempts to execute clone action with retries.
func (c *Clone) Execute() error {
	err := c.DoRetry(c.clone)
	if err != nil {
		return fmt.Errorf("failed to clone repository from URL %s to path %s: %w", c.URL, c.Path, err)
	}
	return nil
}

// Initializes and returns a gitClient instance.
func (c *Clone) getClient() *Client {
	if c.client == nil {
		c.client = NewGitClient()
		c.client.PrepareClient(c.Address, c.Branch)
	}

	return c.client
}

// Attempts to clone a Git repository to the specified path.
// If the cloning process encounters an error, it returns a retryableErr
// containing the original error. Otherwise, it returns nil.
func (c *Clone) clone(_ int) error {
	_, err := c.getClient().Clone(c.Path)
	if err != nil {
		return helper.RetryableErr{Err: err}
	}
	return nil
}
