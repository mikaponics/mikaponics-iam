package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-iam/api"
)

func (s *MikaponicsRPC) AddTimeSeriesDatum(ctx context.Context, in *pb.TimeSeriesDatumRequest) (*pb.MikapodIAMResponse, error) {
	return &pb.MikapodIAMResponse{
		Message: "Instrument was updated",
		Status: true,
	}, nil
}
