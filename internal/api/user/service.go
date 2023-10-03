package user

import pb "github.com/sornick01/auth/pkg/user_v1"

// Service serves
type Service struct {
	pb.UnimplementedUserV1Server
}

// NewService aaa
func NewService() *Service {
	return &Service{}
}
