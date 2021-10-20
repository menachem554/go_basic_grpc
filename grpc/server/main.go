package main
import (
  "net"
  "os"
  mysvccore "github.com/menachem554/go_basic_grpc"
  mysvcgrpc "github.com/menachem554/go_basic_grpc"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
)
func main() {
  // configure our core service
  userService := mysvccore.NewService()
  // configure our gRPC service controller
  userServiceController := NewUserServiceController(userService)
  // start a gRPC server
  server := grpc.NewServer()
  mysvcgrpc.RegisterUserServiceServer(server, userServiceController)
  reflection.Register(server)
  con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
  if err != nil {
    panic(err)
  }
  err = server.Serve(con)
  if err != nil {
    panic(err)
  }
}
