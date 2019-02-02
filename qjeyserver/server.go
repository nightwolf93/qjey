package qjeyserver

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"os"
	"qjey/auth"
	qjeyserverprob "qjey/qjeyserver/qjeyserverprob"
)

type QjeyServer struct {
	Config QjeyServerConfig
	Tokens []*auth.Token
	GrpcServer *grpc.Server
}

// NewQjeyServer .. create a new qjay server instance
func NewQjeyServer(config QjeyServerConfig) *QjeyServer {
	qjeyServer := &QjeyServer{
		Config: config,
	}

	if os.Getenv("API_KEY") == "changemenow" { // warn about the default api key
		logrus.Warn("please change the default API_KEY env var because is not secure")
	}
	logrus.Info(fmt.Sprintf("starting the qjey server on %s:%d", qjeyServer.Config.Host, qjeyServer.Config.Port))
	go qjeyServer.init()

	logrus.Info("qjey server started on ", qjeyServer.Config.Host, ":", qjeyServer.Config.Port)

	return qjeyServer
}

// init .. initialize the qjeyserver and start the grpc server listener
func (s *QjeyServer) init() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	if err != nil {
		logrus.Fatal(fmt.Sprintf("can't listen on %s:%d: %v", s.Config.Host, s.Config.Port, err))
	}
	s.GrpcServer = grpc.NewServer()
	qjeyserverprob.RegisterQjeyServerServer(s.GrpcServer, s)
	if err := s.GrpcServer.Serve(lis); err != nil {
		logrus.Fatal("can't server grpc server: ", err)
	}
}

// findToken .. find a registered token
func (s *QjeyServer) findToken(token string) (*auth.Token) {
	//TODO: Find the token
	return nil
}

// RequestToken .. client request a token for interact with other functions of the server
func (s *QjeyServer) RequestToken(context context.Context, req *qjeyserverprob.RequestTokenMessage) (*qjeyserverprob.RequestTokenResponse, error) {
	peer, _ := peer.FromContext(context)

	if req.ApiKey != os.Getenv("API_KEY") {
		logrus.Warn(peer.Addr.String(), " auth failed, wrong API_KEY given by the client")
		return &qjeyserverprob.RequestTokenResponse{
			Status: -1,
			Token: "",
			Nodes: nil,
		}, nil
	}

	// Create the token and append it to the server token list
	token := auth.NewToken()
	s.Tokens = append(s.Tokens, token)

	logrus.Info("auth success, new token given to the client")

	return &qjeyserverprob.RequestTokenResponse{
		Status: 1,
		Token: token.TokenString,
		Nodes: nil,
	}, nil
}

func (s *QjeyServer) CheckTokenValidity(context context.Context, req *qjeyserverprob.CheckTokenValidityMessage) (*qjeyserverprob.CheckTokenValidityResponse, error) {
	token := s.findToken(req.Token)
	if token == nil {
		return &qjeyserverprob.CheckTokenValidityResponse{
			Valid: false,
		}, nil
	}

	return &qjeyserverprob.CheckTokenValidityResponse{
		Valid: true,
	}, nil
}