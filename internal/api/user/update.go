package user

import (
	"context"

	pb "github.com/sornick01/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update updates
func (s *Service) Update(_ context.Context, _ *pb.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
