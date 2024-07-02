package main

import (
	// "GoSha256/ssl"
	"GoSha256/generator"
	"GoSha256/ssl"
	// "context"
	pb "GoSha256/generator/proto"
	"log/slog"
	"net"

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHashedPasswordGeneratorServer
}

func (s *server) GeneratePassword(in *emptypb.Empty, stream pb.HashedPasswordGenerator_GeneratePasswordServer) error {
	pass_stream := generator.GeneratePasswords()
	var count int32 = 0
	for {
		pass := <-pass_stream
		count = count +1 
		if err := stream.Send(&pb.HashedPassword{
			Index:    count,
			Password: pass,
			Hash:     ssl.Sha256(pass),
		}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHashedPasswordGeneratorServer(s, &server{})
	slog.Info("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve: %v", err)
	}
}
