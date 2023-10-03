package user

import (
	"context"

	pb "github.com/sornick01/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete deletes
func (s *Service) Delete(_ context.Context, _ *pb.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
