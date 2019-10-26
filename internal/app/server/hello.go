package server
import (
	pb "../../api/protobuf-spec"
	"golang.org/x/net/context"
)
type helloService struct{}
func NewHelloService() *helloService {
	return &helloService{}
}
func (h helloService) SayHelloWorld(ctx context.Context, r *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message : "test",
	}, nil
}