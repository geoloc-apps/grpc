package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	grpcClient "github.com/geoloc-apps/grpc"
	pb "github.com/geoloc-apps/grpc/geoloc/proto"
)

type GeolocClientInterface interface {
	GetAllGeolocs(ctx context.Context) (*pb.GetAllGeolocsResponse, error)
	GetGeolocByID(ctx context.Context, req *pb.GetGeolocByIDRequest) (*pb.GetGeolocByIDResponse, error)
	GetGeolocsByIDs(ctx context.Context, req *pb.GetGeolocsByIDsRequest) (*pb.GetGeolocsByIDsResponse, error)
	InsertGeoloc(ctx context.Context, req *pb.InsertGeolocRequest) (*pb.InsertGeolocResponse, error)
	InsertGeolocs(ctx context.Context, req *pb.InsertGeolocsRequest) (*pb.InsertGeolocsResponse, error)
	UpdateGeoloc(ctx context.Context, req *pb.UpdateGeolocRequest) (*pb.UpdateGeolocResponse, error)
	UpdateGeolocs(ctx context.Context, req *pb.UpdateGeolocsRequest) (*pb.UpdateGeolocsResponse, error)
	DeleteGeoloc(ctx context.Context, req *pb.DeleteGeolocRequest) (*pb.DeleteGeolocResponse, error)
	DeleteGeolocs(ctx context.Context, req *pb.DeleteGeolocsRequest) (*pb.DeleteGeolocsResponse, error)
}

type Client struct {
	GeolocClient pb.GeolocServiceClient
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
		GeolocClient: pb.NewGeolocServiceClient(conn),
	}

	return client, nil
}
