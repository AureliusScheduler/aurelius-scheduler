package pb_server

import (
	"aurelius/libs/aurelius-protobuf/generated"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type PbServer struct {
	aurelius_protobuf.UnimplementedAureliusAgentManagerServer
	//Server pb.UnimplementedAureliusAgentManagerServer
}

func NewPbServer() *PbServer {
	return &PbServer{}
}

func (s *PbServer) AgentChat(srv aurelius_protobuf.AureliusAgentManager_AgentChatServer) error {
	log.Println("Start AgentChat")
	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			log.Println("AgentChat Done")
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()

		if err != nil {
			log.Println("AgentChat Recv Error")
			return err
		}

		log.Println("AgentChat Recv", req)

		response := ""

		switch payload := req.Payload.(type) {
		case *aurelius_protobuf.ChatRequest_RegisterAgentRequest:
			registerAgentRequest := payload.RegisterAgentRequest
			log.Println("Register agent request", registerAgentRequest)
			response = "Agent registered"

		case *aurelius_protobuf.ChatRequest_RegisterJobRequest:
			registerJobRequest := payload.RegisterJobRequest
			log.Println("Register job request", registerJobRequest)
			response = "Job registered"
		}

		res := aurelius_protobuf.AgentChatResponse{
			Message: response,
		}

		if err := srv.Send(&res); err != nil {
			log.Println("AgentChat Send Error")
		}

		log.Println("AgentChat Send", res)
	}
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
