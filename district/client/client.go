package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/district/proto"
)

type DistrictClientInterface interface {
	GetAllDistricts(ctx context.Context) (*pb.GetAllDistrictsResponse, error)
	GetDistrictByID(ctx context.Context, req *pb.GetDistrictByIDRequest) (*pb.GetDistrictByIDResponse, error)
	GetDistrictsByIDs(ctx context.Context, req *pb.GetDistrictsByIDsRequest) (*pb.GetDistrictsByIDsResponse, error)
	InsertDistrict(ctx context.Context, req *pb.InsertDistrictRequest) (*pb.InsertDistrictResponse, error)
	InsertDistricts(ctx context.Context, req *pb.InsertDistrictsRequest) (*pb.InsertDistrictsResponse, error)
	UpdateDistrict(ctx context.Context, req *pb.UpdateDistrictRequest) (*pb.UpdateDistrictResponse, error)
	UpdateDistricts(ctx context.Context, req *pb.UpdateDistrictsRequest) (*pb.UpdateDistrictsResponse, error)
	DeleteDistrict(ctx context.Context, req *pb.DeleteDistrictRequest) (*pb.DeleteDistrictResponse, error)
	DeleteDistricts(ctx context.Context, req *pb.DeleteDistrictsRequest) (*pb.DeleteDistrictsResponse, error)
}

type Client struct {
	DistrictClient pb.DistrictServiceClient
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
		DistrictClient: pb.NewDistrictServiceClient(conn),
	}

	return client, nil
}
