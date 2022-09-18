package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/province/proto"
)

type ProvinceClientInterface interface {
	GetAllProvinces(ctx context.Context) (*pb.GetAllProvincesResponse, error)
	GetProvinceByID(ctx context.Context, req *pb.GetProvinceByIDRequest) (*pb.GetProvinceByIDResponse, error)
	GetProvincesByIDs(ctx context.Context, req *pb.GetProvincesByIDsRequest) (*pb.GetProvincesByIDsResponse, error)
	InsertProvince(ctx context.Context, req *pb.InsertProvinceRequest) (*pb.InsertProvinceResponse, error)
	InsertProvinces(ctx context.Context, req *pb.InsertProvincesRequest) (*pb.InsertProvincesResponse, error)
	UpdateProvince(ctx context.Context, req *pb.UpdateProvinceRequest) (*pb.UpdateProvinceResponse, error)
	UpdateProvinces(ctx context.Context, req *pb.UpdateProvincesRequest) (*pb.UpdateProvincesResponse, error)
	DeleteProvince(ctx context.Context, req *pb.DeleteProvinceRequest) (*pb.DeleteProvinceResponse, error)
	DeleteProvinces(ctx context.Context, req *pb.DeleteProvincesRequest) (*pb.DeleteProvincesResponse, error)
}

type Client struct {
	ProvinceClient pb.ProvinceServiceClient
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
		ProvinceClient: pb.NewProvinceServiceClient(conn),
	}

	return client, nil
}
