package client

import (
	"context"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/island/proto"
)

func (c *Client) GetAllIslands(ctx context.Context) (*pb.GetAllIslandsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.GetAllIslands(ctx, nil)
}

func (c *Client) GetIslandByID(ctx context.Context, req *pb.GetIslandByIDRequest) (*pb.GetIslandByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.GetIslandByID(ctx, req)
}

func (c *Client) GetIslandsByIDs(ctx context.Context, req *pb.GetIslandsByIDsRequest) (*pb.GetIslandsByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.GetIslandsByIDs(ctx, req)
}

func (c *Client) InsertIsland(ctx context.Context, req *pb.InsertIslandRequest) (*pb.InsertIslandResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.InsertIsland(ctx, req)
}

func (c *Client) InsertIslands(ctx context.Context, req *pb.InsertIslandsRequest) (*pb.InsertIslandsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.InsertIslands(ctx, req)
}

func (c *Client) UpdateIsland(ctx context.Context, req *pb.UpdateIslandRequest) (*pb.UpdateIslandResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.UpdateIsland(ctx, req)
}

func (c *Client) UpdateIslands(ctx context.Context, req *pb.UpdateIslandsRequest) (*pb.UpdateIslandsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.UpdateIslands(ctx, req)
}

func (c *Client) DeleteIsland(ctx context.Context, req *pb.DeleteIslandRequest) (*pb.DeleteIslandResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.DeleteIsland(ctx, req)
}

func (c *Client) DeleteIslands(ctx context.Context, req *pb.DeleteIslandsRequest) (*pb.DeleteIslandsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.IslandClient.DeleteIslands(ctx, req)
}
