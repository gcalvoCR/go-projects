package main

import (
	"context"
	"fmt"

	"github.com/gcalvoCR/go-learn/cmd/004-grpc/001basic/model"
)

type server struct {
	model.UnimplementedMathServiceServer
}

func (s *server) Add(ctx context.Context, in *model.MathRequest) (*model.MathResponse, error) {
	if s == nil {
		return nil, fmt.Errorf("add called on nil value")
	}
	if in == nil {
		return nil, fmt.Errorf("Add called with invalid parameter value nil")
	}

	result := &model.MathResponse{}
	result.Result = in.Operand1 + in.Operant2
	return result, nil
}
