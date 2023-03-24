package main

import (
	"log"
	"net"

	"github.com/Prameesh-P/calculator/calcpb"
	"google.golang.org/grpc"
)

const(

	port=":9000"
)

type CalcServer struct{

	calcpb.ServicesServer
}

func main() {

	lis,err:=net.Listen("tcp",port)
	if err !=nil{
		log.Fatalf("Error while get listening on port %s, and err :%v",port,err)
	}
	grpcServer:=grpc.NewServer()

	calcpb.RegisterServicesServer(grpcServer,&CalcServer{})

	if err:=grpcServer.Serve(lis);err!=nil{
		log.Fatalf("Error while get Serving listener into grpce server %v",err)

	}

	log.Printf("Server is started on %s",port)
	
}