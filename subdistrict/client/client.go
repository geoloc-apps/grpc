package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/subdistrict/proto"
)

type SubdistrictClientInterface interface {
	GetAllSubdistricts(ctx context.Context) (*pb.GetAllSubdistrictsResponse, error)
	GetSubdistrictByID(ctx context.Context, req *pb.GetSubdistrictByIDRequest) (*pb.GetSubdistrictByIDResponse, error)
	GetSubdistrictsByIDs(ctx context.Context, req *pb.GetSubdistrictsByIDsRequest) (*pb.GetSubdistrictsByIDsResponse, error)
	InsertSubdistrict(ctx context.Context, req *pb.InsertSubdistrictRequest) (*pb.InsertSubdistrictResponse, error)
	InsertSubdistricts(ctx context.Context, req *pb.InsertSubdistrictsRequest) (*pb.InsertSubdistrictsResponse, error)
	UpdateSubdistrict(ctx context.Context, req *pb.UpdateSubdistrictRequest) (*pb.UpdateSubdistrictResponse, error)
	UpdateSubdistricts(ctx context.Context, req *pb.UpdateSubdistrictsRequest) (*pb.UpdateSubdistrictsResponse, error)
	DeleteSubdistrict(ctx context.Context, req *pb.DeleteSubdistrictRequest) (*pb.DeleteSubdistrictResponse, error)
	DeleteSubdistricts(ctx context.Context, req *pb.DeleteSubdistrictsRequest) (*pb.DeleteSubdistrictsResponse, error)
}

type Client struct {
	SubdistrictClient pb.SubdistrictServiceClient
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
		SubdistrictClient: pb.NewSubdistrictServiceClient(conn),
	}

	return client, nil
}
