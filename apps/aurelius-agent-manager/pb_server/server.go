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

//func (s *PbServer) RegisterJob(ctx context.Context, request *aurelius_protobuf.RegisterJobRequest) (*aurelius_protobuf.RegisterJobResponse, error) {
//  log.Println("Start RegisterJob")
//
//  res := aurelius_protobuf.RegisterJobResponse{
//    JobId:   "123132",
//    Message: "Job registered",
//  }
//
//  return &res, nil
//}

//func (s *PbServer) AgentChat(srv aurelius_protobuf.AureliusAgentManager_AgentChatServer) error {
//	log.Println("Start AgentChat")
//	ctx := srv.Context()
//
//	for {
//		select {
//		case <-ctx.Done():
//			log.Println("AgentChat Done")
//			return ctx.Err()
//		default:
//		}
//
//		req, err := srv.Recv()
//
//		if err != nil {
//			log.Println("AgentChat Recv Error")
//			return err
//		}
//
//		log.Println("AgentChat Recv", req)
//
//		response := ""
//
//		switch payload := req.Payload.(type) {
//		case *aurelius_protobuf.ChatRequest_RegisterAgentRequest:
//			registerAgentRequest := payload.RegisterAgentRequest
//			log.Println("Register agent request", registerAgentRequest)
//			response = "Agent registered"
//
//		case *aurelius_protobuf.ChatRequest_RegisterJobRequest:
//			registerJobRequest := payload.RegisterJobRequest
//			log.Println("Register job request", registerJobRequest)
//			response = "Job registered"
//		}
//
//		res := aurelius_protobuf.AgentChatResponse{
//			Message: response,
//		}
//
//		if err := srv.Send(&res); err != nil {
//			log.Println("AgentChat Send Error")
//		}
//
//		log.Println("AgentChat Send", res)
//	}
//}

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
