package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/island/proto"
)

type IslandClientInterface interface {
	GetAllIslands(ctx context.Context) (*pb.GetAllIslandsResponse, error)
	GetIslandByID(ctx context.Context, req *pb.GetIslandByIDRequest) (*pb.GetIslandByIDResponse, error)
	GetIslandsByIDs(ctx context.Context, req *pb.GetIslandsByIDsRequest) (*pb.GetIslandsByIDsResponse, error)
	InsertIsland(ctx context.Context, req *pb.InsertIslandRequest) (*pb.InsertIslandResponse, error)
	InsertIslands(ctx context.Context, req *pb.InsertIslandsRequest) (*pb.InsertIslandsResponse, error)
	UpdateIsland(ctx context.Context, req *pb.UpdateIslandRequest) (*pb.UpdateIslandResponse, error)
	UpdateIslands(ctx context.Context, req *pb.UpdateIslandsRequest) (*pb.UpdateIslandsResponse, error)
	DeleteIsland(ctx context.Context, req *pb.DeleteIslandRequest) (*pb.DeleteIslandResponse, error)
	DeleteIslands(ctx context.Context, req *pb.DeleteIslandsRequest) (*pb.DeleteIslandsResponse, error)
}

type Client struct {
	IslandClient pb.IslandServiceClient
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
		IslandClient: pb.NewIslandServiceClient(conn),
	}

	return client, nil
}
