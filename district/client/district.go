package client

import (
	"context"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/district/proto"
)

func (c *Client) GetAllDistricts(ctx context.Context) (*pb.GetAllDistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.GetAllDistricts(ctx, nil)
}

func (c *Client) GetDistrictByID(ctx context.Context, req *pb.GetDistrictByIDRequest) (*pb.GetDistrictByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.GetDistrictByID(ctx, req)
}

func (c *Client) GetDistrictsByIDs(ctx context.Context, req *pb.GetDistrictsByIDsRequest) (*pb.GetDistrictsByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.GetDistrictsByIDs(ctx, req)
}

func (c *Client) InsertDistrict(ctx context.Context, req *pb.InsertDistrictRequest) (*pb.InsertDistrictResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.InsertDistrict(ctx, req)
}

func (c *Client) InsertDistricts(ctx context.Context, req *pb.InsertDistrictsRequest) (*pb.InsertDistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.InsertDistricts(ctx, req)
}

func (c *Client) UpdateDistrict(ctx context.Context, req *pb.UpdateDistrictRequest) (*pb.UpdateDistrictResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.UpdateDistrict(ctx, req)
}

func (c *Client) UpdateDistricts(ctx context.Context, req *pb.UpdateDistrictsRequest) (*pb.UpdateDistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.UpdateDistricts(ctx, req)
}

func (c *Client) DeleteDistrict(ctx context.Context, req *pb.DeleteDistrictRequest) (*pb.DeleteDistrictResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.DeleteDistrict(ctx, req)
}

func (c *Client) DeleteDistricts(ctx context.Context, req *pb.DeleteDistrictsRequest) (*pb.DeleteDistrictsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.DistrictClient.DeleteDistricts(ctx, req)
}
