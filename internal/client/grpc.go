package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/protocol/activity"
	"github.com/terrapi-solution/protocol/deployment"
	"github.com/terrapi-solution/protocol/health"
	"github.com/terrapi-solution/runner/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"strconv"
)

type Client struct {
	Activity   activity.ActivityServiceClient
	Deployment deployment.DeploymentServiceClient
	Health     health.HealthServiceClient
}

// NewClient initializes and returns a gRPC client
func NewClient() *Client {
	// Load the configuration
	cfg := config.Get()

	// Load the TLS certificate
	tlsConfig, err := loadTLSConfig(
		cfg.Controller.Certificates.CertFile,
		cfg.Controller.Certificates.KeyFile,
		cfg.Controller.Certificates.CaFile)
	if err != nil {
		log.Panic().Err(err).Msg("failed to load TLS configuration")
	}

	address := net.JoinHostPort(cfg.Controller.Host, strconv.Itoa(cfg.Controller.Port))
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(tlsConfig))
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to gRPC server")
	}

	client := &Client{
		Activity:   activity.NewActivityServiceClient(conn),
		Deployment: deployment.NewDeploymentServiceClient(conn),
		Health:     health.NewHealthServiceClient(conn),
	}

	return client
}

// loadTLSConfig loads the TLS configuration
func loadTLSConfig(certFile, keyFile, caFile string) (credentials.TransportCredentials, error) {
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load client certification: %w", err)
	}

	ca, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("faild to read CA certificate: %w", err)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(ca) {
		return nil, fmt.Errorf("faild to append the CA certificate to CA pool")
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: true,
		RootCAs:            pool,
	}

	return credentials.NewTLS(tlsConfig), nil
}
