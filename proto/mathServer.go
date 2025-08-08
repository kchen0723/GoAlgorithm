package main

import (
	"context"
	"log"
	// mathpb "./proto"
)

type MathService struct {
	UnimplementedMathServiceServer
}

func (m *MathService) Sum(c context.Context, req *MathRequest) (*MathResponse, error) {
	sum := req.Num1 + req.Num2
	log.Printf("num1: %d, num2: %d, sum: %d\n", req.Num1, req.Num2, sum)
	return &MathResponse{Reply: sum}, nil
}

func (m *MathService) Multi(c context.Context, req *MathRequest) (*MathResponse, error) {
	reply := req.Num1 * req.Num2
	log.Printf("num1: %d, num2: %d, reply: %d\n", req.Num1, req.Num2, reply)
	return &MathResponse{Reply: reply}, nil
}
