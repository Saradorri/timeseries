package grpcserver

import (
	"context"
	tpb "edgecom.ai/timeseries/internal/proto/pb"
	"edgecom.ai/timeseries/internal/services"
	"edgecom.ai/timeseries/pkg/models"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer interface {
	StartServer() error
	QueryData(ctx context.Context, req *tpb.QueryRequest) (*tpb.QueryResponse, error)
}

type grpcServer struct {
	port    int
	service services.TimeSeriesService
	tpb.UnimplementedTimeSeriesServiceServer
}

func NewServer(p int, tsService services.TimeSeriesService) GrpcServer {
	return &grpcServer{
		port:                                 p,
		service:                              tsService,
		UnimplementedTimeSeriesServiceServer: tpb.UnimplementedTimeSeriesServiceServer{},
	}
}

func (s *grpcServer) StartServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	tpb.RegisterTimeSeriesServiceServer(server, s)

	log.Printf("gRPC grpcServer running on port %d", s.port)
	return server.Serve(lis)
}

func (s *grpcServer) QueryData(ctx context.Context, req *tpb.QueryRequest) (*tpb.QueryResponse, error) {
	response := &tpb.QueryResponse{
		Meta: &tpb.QueryMetadata{
			Aggregation: req.Aggregation,
			Window:      req.Window,
			Status:      tpb.QueryStatus_ERROR,
		},
	}

	result, err := s.service.GetByQuery(ctx, models.TimeSeriesQuery{
		Start:       req.Start,
		End:         req.End,
		Window:      req.Window,
		Aggregation: req.Aggregation,
	})

	if err != nil {
		response.Meta.Message = err.Error()
		return response, err
	}

	for _, result := range result {
		tsData := &tpb.TimeSeriesData{
			Time:  result.Time,
			Value: result.Value,
		}
		response.Data = append(response.Data, tsData)
	}

	if len(result) > 0 {
		response.Meta.Status = tpb.QueryStatus_SUCCESS
		response.Meta.Message = "Query executed successfully"
	} else {
		response.Meta.Status = tpb.QueryStatus_ERROR
		response.Meta.Message = "No data found for the given parameters"
	}

	return response, nil
}
