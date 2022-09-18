package client

import (
	"context"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/geoloc/proto"
)

func (c *Client) GetAllGeolocs(ctx context.Context) (*pb.GetAllGeolocsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.GetAllGeolocs(ctx, nil)
}

func (c *Client) GetGeolocByID(ctx context.Context, req *pb.GetGeolocByIDRequest) (*pb.GetGeolocByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.GetGeolocByID(ctx, req)
}

func (c *Client) GetGeolocsByIDs(ctx context.Context, req *pb.GetGeolocsByIDsRequest) (*pb.GetGeolocsByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.GetGeolocsByIDs(ctx, req)
}

func (c *Client) InsertGeoloc(ctx context.Context, req *pb.InsertGeolocRequest) (*pb.InsertGeolocResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.InsertGeoloc(ctx, req)
}

func (c *Client) InsertGeolocs(ctx context.Context, req *pb.InsertGeolocsRequest) (*pb.InsertGeolocsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.InsertGeolocs(ctx, req)
}

func (c *Client) UpdateGeoloc(ctx context.Context, req *pb.UpdateGeolocRequest) (*pb.UpdateGeolocResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.UpdateGeoloc(ctx, req)
}

func (c *Client) UpdateGeolocs(ctx context.Context, req *pb.UpdateGeolocsRequest) (*pb.UpdateGeolocsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.UpdateGeolocs(ctx, req)
}

func (c *Client) DeleteGeoloc(ctx context.Context, req *pb.DeleteGeolocRequest) (*pb.DeleteGeolocResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.DeleteGeoloc(ctx, req)
}

func (c *Client) DeleteGeolocs(ctx context.Context, req *pb.DeleteGeolocsRequest) (*pb.DeleteGeolocsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.GeolocClient.DeleteGeolocs(ctx, req)
}
