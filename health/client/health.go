package client

import (
	"context"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/health/proto"
)

func (c *Client) Health(ctx context.Context) (*pb.HealthResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.HealthClient.Health(ctx, nil)
}
