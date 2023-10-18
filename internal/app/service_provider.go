package app

import "github.com/sornick01/auth/internal/config"

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
}
