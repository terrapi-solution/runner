package client

import (
	"fmt"
	"github.com/thomas-illiet/terrapi-runner/pkg/config"
	"io"
	"net/http"
)

type Client struct {
	config config.Config
}

func NewClient(config config.Config) *Client {
	return &Client{config: config}
}

func (c *Client) GetDeploymentConfiguration(identifier uint) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/v1/deployments/%d", identifier))
	if err != nil {
		return "", fmt.Errorf("failed to get deployment: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	return string(body), nil
}

func (c *Client) SetDeploymentStatus(identifier uint, status string) {
}

func (c *Client) AddDeploymentStdoutLog(identifier uint, content string) {
}

func (c *Client) AddDeploymentStderrLog(identifier uint, content string) {
}
