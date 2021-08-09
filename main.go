package main

import (
	"context"
	trfmain "git-connector/pkg/rule/TRF_Main"
	util "git-connector/pkg/util"
)

import (
	_ "git-connector/docs"
)
import (
	"git-connector/pkg/middleware"
)
import (
	"git-connector/pkg/controller"
	"git-connector/pkg/xiLogger"
	"git-connector/sidecar"
	"log"
	"net"
)

func main() {

	//controller.InitGormCrypto() // GormCrypto

	cntxt := util.CustomContext{}
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	trfmain.TRF_Main(&cntxt, config)

	RunServer()
}

func RunServer() {
	port := middleware.GetGrpcPort()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := middleware.GetGrpcServerConfigs()

	xiLogger.Log.Info("\n Server listening on port %v \n", ":"+port)
	sidecar.RegisterServiceServer(s, &controller.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
