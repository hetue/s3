package internal

import (
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hetue/s3/internal/internal/config"
)

type Client = s3.Client

func newClient(server *config.Server, secret *config.Secret) *Client {
	options := s3.Options{
		Credentials:      credentials.NewStaticCredentialsProvider(secret.Ak, secret.Sk, secret.Session),
		Region:           server.Region,
		EndpointResolver: s3.EndpointResolverFromURL(server.Endpoint),
		UsePathStyle:     false,
	}

	return s3.New(options)
}
