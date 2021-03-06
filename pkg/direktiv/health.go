package direktiv

import (
	"context"

	"google.golang.org/grpc"

	"github.com/vorteil/direktiv/pkg/health"
)

type healthServer struct {
	health.UnimplementedHealthServer

	config *Config
	grpc   *grpc.Server
}

func newHealthServer(config *Config) *healthServer {
	return &healthServer{
		config: config,
	}
}

func (hs *healthServer) start(s *WorkflowServer) error {
	return s.grpcStart(&hs.grpc, "health", s.config.HealthAPI.Bind, func(srv *grpc.Server) {
		health.RegisterHealthServer(srv, hs)
	})
}

func (hs *healthServer) stop() {

	if hs.grpc != nil {
		hs.grpc.GracefulStop()
	}

}

func (hs *healthServer) name() string {
	return "health"
}

func (hs *healthServer) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {

	var resp health.HealthCheckResponse

	resp.Status = health.HealthCheckResponse_SERVING

	return &resp, nil

}

func (hs *healthServer) Watch(in *health.HealthCheckRequest, srv health.Health_WatchServer) error {

	var resp health.HealthCheckResponse

	resp.Status = health.HealthCheckResponse_SERVING

	err := srv.Send(&resp)
	if err != nil {
		return err
	}

	return nil

}
