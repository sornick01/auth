package user

import (
	"context"

	pb "github.com/sornick01/auth/pkg/user_v1"
)

// Create creates
func (s *Service) Create(_ context.Context, _ *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Id: 1231241412,
	}, nil
}
