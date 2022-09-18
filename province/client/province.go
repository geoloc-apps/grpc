package client

import (
	"context"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/province/proto"
)

func (c *Client) GetAllProvinces(ctx context.Context) (*pb.GetAllProvincesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.GetAllProvinces(ctx, nil)
}

func (c *Client) GetProvinceByID(ctx context.Context, req *pb.GetProvinceByIDRequest) (*pb.GetProvinceByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.GetProvinceByID(ctx, req)
}

func (c *Client) GetProvincesByIDs(ctx context.Context, req *pb.GetProvincesByIDsRequest) (*pb.GetProvincesByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.GetProvincesByIDs(ctx, req)
}

func (c *Client) InsertProvince(ctx context.Context, req *pb.InsertProvinceRequest) (*pb.InsertProvinceResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.InsertProvince(ctx, req)
}

func (c *Client) InsertProvinces(ctx context.Context, req *pb.InsertProvincesRequest) (*pb.InsertProvincesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.InsertProvinces(ctx, req)
}

func (c *Client) UpdateProvince(ctx context.Context, req *pb.UpdateProvinceRequest) (*pb.UpdateProvinceResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.UpdateProvince(ctx, req)
}

func (c *Client) UpdateProvinces(ctx context.Context, req *pb.UpdateProvincesRequest) (*pb.UpdateProvincesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.UpdateProvinces(ctx, req)
}

func (c *Client) DeleteProvince(ctx context.Context, req *pb.DeleteProvinceRequest) (*pb.DeleteProvinceResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.DeleteProvince(ctx, req)
}

func (c *Client) DeleteProvinces(ctx context.Context, req *pb.DeleteProvincesRequest) (*pb.DeleteProvincesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.ProvinceClient.DeleteProvinces(ctx, req)
}
