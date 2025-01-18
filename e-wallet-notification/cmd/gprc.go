package cmd

import (
	"e-wallet-notification/cmd/proto/notification"
	"e-wallet-notification/helpers"
	"e-wallet-notification/internal/api"
	"e-wallet-notification/internal/repository"
	"e-wallet-notification/internal/services"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func ServeGRPC() {
	d := dependencyInject()
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "5000"))
	if err != nil {
		log.Fatal("Failed to listen grpc port: ", err)
	}
	s := grpc.NewServer()

	// list method
	notification.RegisterNotificationServiceServer(s, d.EmailAPI)

	logrus.Info("start grpc server on port: ", helpers.GetEnv("GRPC_PORT", "5000"), "")
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve grpc port: ", err)
	}
}

type Dependency struct {
	EmailAPI *api.EmailAPI
}

func dependencyInject() *Dependency {
	emailRepo := &repository.EmailRepo{DB: helpers.DB}
	emailSvc := &services.EmailService{EmailRepo: emailRepo}
	emailApi := &api.EmailAPI{
		EmailService: emailSvc,
	}

	return &Dependency{
		EmailAPI: emailApi,
	}
}
