package grpcserver

import (
	"context"
	tpb "edgecom.ai/timeseries/internal/proto/pb"
	"edgecom.ai/timeseries/internal/services"
	"edgecom.ai/timeseries/pkg/models"
	"edgecom.ai/timeseries/pkg/validation"
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
	var server *grpc.Server

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	server = grpc.NewServer()
	tpb.RegisterTimeSeriesServiceServer(server, s)

	log.Printf("gRPC server running on port %d", s.port)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	return nil
}

func (s *grpcServer) QueryData(ctx context.Context, req *tpb.QueryRequest) (*tpb.QueryResponse, error) {

	m := models.TimeSeriesQuery{
		Start:       req.Start,
		End:         req.End,
		Window:      req.Window,
		Aggregation: req.Aggregation,
	}
	err := validation.ValidateQueryRequest(m)
	if err != nil {
		log.Printf("Validation error: %v", err)
		return &tpb.QueryResponse{
			Meta: &tpb.QueryMetadata{
				Message: fmt.Sprintf("Validation error: %v", err),
				Status:  tpb.QueryStatus_ERROR,
			},
		}, err
	}
	result, err := s.service.GetByQuery(ctx, m)

	if err != nil {
		log.Printf("Error fetching data with query: %v", err)
		return &tpb.QueryResponse{
			Meta: &tpb.QueryMetadata{
				Message: fmt.Sprintf("Error retrieving data: %v", err),
				Status:  tpb.QueryStatus_ERROR,
			},
		}, err
	}
	return result.ToProtoResponse(req.Aggregation, req.Window), nil
}
