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

	log.Printf("gRPC server running on port %d", s.port)
	return server.Serve(lis)
}

func (s *grpcServer) QueryData(ctx context.Context, req *tpb.QueryRequest) (*tpb.QueryResponse, error) {
	result, err := s.service.GetByQuery(ctx, models.TimeSeriesQuery{
		Start:       req.Start,
		End:         req.End,
		Window:      req.Window,
		Aggregation: req.Aggregation,
	})

	if err != nil {
		log.Printf("Error fetching data with query: %v", err)
		return &tpb.QueryResponse{
			Meta: &tpb.QueryMetadata{
				Message: fmt.Sprintf("Error retrieving data: %v", err),
				Status:  tpb.QueryStatus_ERROR,
			},
		}, err
	}
	return toProtoResponse(result, req.Aggregation, req.Window), nil
}
