package main

import (
	"log"
	"strings"

	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// go run example/grpc_server/app.go
// Visit: https://127.0.0.1:3000/ or go run example/grpc_client/app.go
func main() {

	app := gear.New()
	app.UseHandler(logging.Default())
	app.Use(func(ctx *gear.Context) error {
		if !strings.HasPrefix(ctx.Path, "/helloworld.Greeter/SayHello") {
			// HTTP request/response
			return ctx.HTML(200, "<h1>gRPC</h1>")
		}
		return nil
	})

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	app.Use(gear.WrapHandler(s))
	log.Fatalln(app.ListenTLS(":3000", "./testdata/out/test.crt", "./testdata/out/test.key"))
}