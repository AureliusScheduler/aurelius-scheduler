package pb_server

import (
	"aurelius/apps/aurelius-agent-manager/service"
	"aurelius/libs/aurelius-protobuf/generated"
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type PbServer struct {
	aurelius_protobuf.UnimplementedAureliusAgentManagerServer
	agentService *service.AgentService
}

func NewPbServer(
	agentService *service.AgentService,
) *PbServer {
	return &PbServer{
		agentService: agentService,
	}
}

func (s *PbServer) RegisterAgent(ctx context.Context, request *aurelius_protobuf.RegisterAgentRequest) (*aurelius_protobuf.RegisterAgentResponse, error) {
	log.Println("Start RegisterAgent")

	registerAgentParams := &service.RegisterAgentParams{
		Name:     request.Name,
		Version:  request.Version,
		Stack:    request.Stack,
		Instance: request.Instance,
	}

	var registerJobs []*service.RegisterJobParams

	for _, job := range request.Jobs {
		registerJobs = append(registerJobs, &service.RegisterJobParams{
			Name:    job.Name,
			Version: job.Version,
		})
	}

	result, err := s.agentService.RegisterAgent(registerAgentParams, registerJobs)

	if err != nil {
		return nil, err
	}

	res := aurelius_protobuf.RegisterAgentResponse{
		AgentId: result.AgentId,
	}

	return &res, nil
}

func (s *PbServer) Listen() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	aurelius_protobuf.RegisterAureliusAgentManagerServer(server, s)
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var Module = fx.Options(
	fx.Provide(NewPbServer),
)
