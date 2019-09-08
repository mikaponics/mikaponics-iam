package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-iam/api"
)

func (s *MikaponicsIAMServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Token: "Instrument was updated",
		Status: true,
	}, nil
}


func (s *MikaponicsIAMServer) VerifyAccessToken(ctx context.Context, in *pb.VerifyAccessTokenRequest) (*pb.VerifyAccessTokenResponse, error) {
	return &pb.VerifyAccessTokenResponse{
		Status: true,
		UserId: 0,
		ThingId: 0,
		ExpiresAt: nil,
	}, nil
}


func (s *MikaponicsIAMServer) RefreshAccessToken(ctx context.Context, in *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	return &pb.RefreshAccessTokenResponse{
		Status: true,
		UserId: 0,
		ThingId: 0,
		ExpiresAt: nil,
	}, nil
}
