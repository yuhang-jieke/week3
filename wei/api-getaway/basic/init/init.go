package init

import (
	"flag"
	"log"
	"week3/wei/api-getaway/basic/config"
	__ "week3/wei/api-getaway/basic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	GrpcInit()
}
func GrpcInit() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	config.GoodsClient = __.NewGoodsClient(conn)
}
