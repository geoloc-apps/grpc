package client

import (
	"context"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/city/proto"
)

func (c *Client) GetAllCities(ctx context.Context) (*pb.GetAllCitiesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.GetAllCities(ctx, nil)
}

func (c *Client) GetCityByID(ctx context.Context, req *pb.GetCityByIDRequest) (*pb.GetCityByIDResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.GetCityByID(ctx, req)
}

func (c *Client) GetCitiesByIDs(ctx context.Context, req *pb.GetCitiesByIDsRequest) (*pb.GetCitiesByIDsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.GetCitiesByIDs(ctx, req)
}

func (c *Client) InsertCity(ctx context.Context, req *pb.InsertCityRequest) (*pb.InsertCityResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.InsertCity(ctx, req)
}

func (c *Client) InsertCities(ctx context.Context, req *pb.InsertCitiesRequest) (*pb.InsertCitiesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.InsertCities(ctx, req)
}

func (c *Client) UpdateCity(ctx context.Context, req *pb.UpdateCityRequest) (*pb.UpdateCityResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.UpdateCity(ctx, req)
}

func (c *Client) UpdateCities(ctx context.Context, req *pb.UpdateCitiesRequest) (*pb.UpdateCitiesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.UpdateCities(ctx, req)
}

func (c *Client) DeleteCity(ctx context.Context, req *pb.DeleteCityRequest) (*pb.DeleteCityResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.DeleteCity(ctx, req)
}

func (c *Client) DeleteCities(ctx context.Context, req *pb.DeleteCitiesRequest) (*pb.DeleteCitiesResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, haveTimeout := ctx.Deadline(); !haveTimeout {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultTimeout)
		defer cancel()
		ctx = newCtx
	}

	return c.CityClient.DeleteCities(ctx, req)
}
