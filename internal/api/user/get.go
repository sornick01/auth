package user

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	pb "github.com/sornick01/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Get gets
func (s *Service) Get(_ context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	//fmt.Println(color.Green("requstes id: %d", req.Id))
	fmt.Println(color.GreenString("requested id: %d", req.Id))

	now := time.Now()

	return &pb.GetResponse{
		Id:        req.Id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      pb.Role_user,
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}, nil
}
