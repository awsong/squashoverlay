package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/awsong/squashoverlay"
	snapshotsapi "github.com/containerd/containerd/api/services/snapshots/v1"
	"github.com/containerd/containerd/contrib/snapshotservice"
)

func main() {
	// Provide a unix address to listen to, this will be the `address`
	// in the `proxy_plugin` configuration.
	// The root will be used to store the snapshots.
	if len(os.Args) < 3 {
		fmt.Printf("invalid args: usage: %s <unix addr> <root>\n", os.Args[0])
		os.Exit(1)
	}

	rpc := grpc.NewServer()

	sn, err := squashoverlay.NewSnapshotter(os.Args[2], squashoverlay.AsynchronousRemove)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	// Convert the snapshotter to a gRPC service,
	service := snapshotservice.FromSnapshotter(sn)

	// Register the service with the gRPC server
	snapshotsapi.RegisterSnapshotsServer(rpc, service)

	l, err := net.Listen("unix", os.Args[1])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	if err := rpc.Serve(l); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
