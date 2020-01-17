package main

import (
	"flag"
	"fmt"
	"github.com/ruben-zhi/consul-go"
	api "github.com/ruben-zhi/consul-go/example/server/api/echo"
	echo_pb "github.com/ruben-zhi/consul-go/example/server/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	health_pb "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := flag.String("p", "13831", "gRPC server port")
	flag.Parse()

	server := grpc.NewServer()
	echo_pb.RegisterEchoServer(server, &api.EchoServer{})

	hsrv := health.NewServer()
	hsrv.SetServingStatus("", health_pb.HealthCheckResponse_SERVING)
	health_pb.RegisterHealthServer(server, hsrv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalln(err)
	}

	register := consul_go.NewConsulRegister()
	err = register.Register("id-grpc-go", "grpc-go", 13831)
	if err != nil {
		log.Fatalln(err)
	}

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)
	errChan := make(chan error, 5)

	go func() {
		err := server.Serve(lis)
		errChan <- err
	}()
	defer close(errChan)

	select {
	case <-stopSignal:
		log.Println("graceful stop ...")
		register.DeregisterRegister("id-grpc-go")
		server.GracefulStop()
		log.Println("graceful stop gRPC server")
	case e := <-errChan:
		log.Println(e.Error())
	}
}
