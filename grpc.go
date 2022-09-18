package grpc

import "time"

var (
	DefaultTimeout            = 2 * time.Second
	DefaultConnectGRPCTimeout = 5 * time.Second
	DefaultSleepTime          = 100 * time.Millisecond
)
