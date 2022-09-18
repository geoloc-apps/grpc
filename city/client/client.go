package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/city/proto"
)

type CityClientInterface interface {
	GetAllCities(ctx context.Context) (*pb.GetAllCitiesResponse, error)
	GetCityByID(ctx context.Context, req *pb.GetCityByIDRequest) (*pb.GetCityByIDResponse, error)
	GetCitiesByIDs(ctx context.Context, req *pb.GetCitiesByIDsRequest) (*pb.GetCitiesByIDsResponse, error)
	InsertCity(ctx context.Context, req *pb.InsertCityRequest) (*pb.InsertCityResponse, error)
	InsertCities(ctx context.Context, req *pb.InsertCitiesRequest) (*pb.InsertCitiesResponse, error)
	UpdateCity(ctx context.Context, req *pb.UpdateCityRequest) (*pb.UpdateCityResponse, error)
	UpdateCities(ctx context.Context, req *pb.UpdateCitiesRequest) (*pb.UpdateCitiesResponse, error)
	DeleteCity(ctx context.Context, req *pb.DeleteCityRequest) (*pb.DeleteCityResponse, error)
	DeleteCities(ctx context.Context, req *pb.DeleteCitiesRequest) (*pb.DeleteCitiesResponse, error)
}

type Client struct {
	CityClient pb.CityServiceClient
}

type Options struct {
	Address string
}

func NewClient(ctx context.Context, o Options, opts ...grpc.DialOption) (*Client, error) {
	if ctx == nil {
		newCtx, cancel := context.WithTimeout(context.Background(), grpcClient.DefaultConnectGRPCTimeout)
		defer cancel()
		ctx = newCtx
	}

	client, err := getClientWithContext(ctx, o)
	if err != nil {
		log.Fatalf("grpc connection failed, err: %v", err)
	}

	return client, nil
}

func getClientWithContext(ctx context.Context, o Options) (*Client, error) {
	defer time.Sleep(grpcClient.DefaultSleepTime)

	if ctx == context.Background() || ctx == context.TODO() {
		newCtx, cancel := context.WithTimeout(ctx, grpcClient.DefaultConnectGRPCTimeout)
		defer cancel()
		ctx = newCtx
	}

	dialOpts := []grpc.DialOption{}
	dialOpts = append(dialOpts, grpc.WithInsecure())

	conn, err := grpc.DialContext(ctx, o.Address, dialOpts...)
	if err != nil {
		log.Printf("grpc error dialing to: %s, err: %v", o.Address, err)
		if conn != nil {
			conn.Close()
		}
		return nil, err
	}

	client := &Client{
		CityClient: pb.NewCityServiceClient(conn),
	}

	return client, nil
}
