package client

import (
	"context"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/subdistrict/proto"
)

func (c *Client) GetAllSubdistricts(ctx context.Context) (*pb.GetAllSubdistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.GetAllSubdistricts(ctx, nil)
}

func (c *Client) GetSubdistrictByID(ctx context.Context, req *pb.GetSubdistrictByIDRequest) (*pb.GetSubdistrictByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.GetSubdistrictByID(ctx, req)
}

func (c *Client) GetSubdistrictsByIDs(ctx context.Context, req *pb.GetSubdistrictsByIDsRequest) (*pb.GetSubdistrictsByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.GetSubdistrictsByIDs(ctx, req)
}

func (c *Client) InsertSubdistrict(ctx context.Context, req *pb.InsertSubdistrictRequest) (*pb.InsertSubdistrictResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.InsertSubdistrict(ctx, req)
}

func (c *Client) InsertSubdistricts(ctx context.Context, req *pb.InsertSubdistrictsRequest) (*pb.InsertSubdistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.InsertSubdistricts(ctx, req)
}

func (c *Client) UpdateSubdistrict(ctx context.Context, req *pb.UpdateSubdistrictRequest) (*pb.UpdateSubdistrictResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.UpdateSubdistrict(ctx, req)
}

func (c *Client) UpdateSubdistricts(ctx context.Context, req *pb.UpdateSubdistrictsRequest) (*pb.UpdateSubdistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.UpdateSubdistricts(ctx, req)
}

func (c *Client) DeleteSubdistrict(ctx context.Context, req *pb.DeleteSubdistrictRequest) (*pb.DeleteSubdistrictResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.DeleteSubdistrict(ctx, req)
}

func (c *Client) DeleteSubdistricts(ctx context.Context, req *pb.DeleteSubdistrictsRequest) (*pb.DeleteSubdistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.SubdistrictClient.DeleteSubdistricts(ctx, req)
}
