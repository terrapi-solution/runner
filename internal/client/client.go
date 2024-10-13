package client

import (
	"fmt"
	"github.com/thomas-illiet/terrapi-runner/internal/client/model"
	"github.com/thomas-illiet/terrapi-runner/internal/config"
	"io"
	"net/http"
)

type Client struct {
	config *config.Config
}

func NewClient(config *config.Config) *Client {
	return &Client{config: config}
}

func (c *Client) GetDeployment(identifier uint) (model.Deployment, error) {
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

func (c *Client) SetDeploymentStatus(identifier uint, status string) error {
	return nil
}

func (c *Client) AddDeploymentStdoutLog(identifier uint, content string) error {
	return nil
}

func (c *Client) AddDeploymentStderrLog(identifier uint, content string) error {
	return nil
}
