package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type requestModule struct {
	Address  string `json:"address"`
	Branch   string `json:"branch"`
	Path     string `json:"path"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type requestAction struct {
	Name      string            `json:"name"`
	Variables []requestVariable `json:"variables"`
}

type requestVariable struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Secret bool   `json:"secret"`
}

type deploymentConfig struct {
	Module requestModule `json:"module"`
	Action requestAction `json:"action"`
}

type deployment struct {
	retryHelper

	Config deploymentConfig `json:"config"`
	Runner RunnerConfig     `json:"runner"`
}

// Creates a new instance of deploymentConfig.
func NewDeployment() *deployment {
	return &deployment{
		retryHelper: retryHelper{Retry: 3, RetryTime: 2 * time.Second},
	}
}

func (c *deployment) Load(id uint) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/v1/deployments/%d", id))
	if err != nil {
		return fmt.Errorf("failed to get deployment: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	log.Printf("Response Body: %s", string(body))
	return nil
}
