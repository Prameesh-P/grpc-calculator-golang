package main

import (
	"fmt"
	"log"
	"github.com/Prameesh-P/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const(

	port =":9000"
)

func main() {

	conn,err:=grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		log.Fatalf("error while dialing on the port %s , err: %v",port,err)
	}

	defer conn.Close()

	client:=calcpb.NewServicesClient(conn)

	firstNum:=int32(0)
	secondNum:=int32(0)

	log.Println("Enter First Number for addition : ")
	
	_,err=fmt.Scanln(&firstNum)
	if err !=nil{
		log.Fatalf("error %v",err)
	}

	fmt.Println("Enter Second Number for addition : ")

	_,err=fmt.Scanln(&secondNum)
	
	if err !=nil{
		log.Fatalf("error %v",err)
	}
	opNum:=EnterNumber()
	res:=&calcpb.AdditionRequestFromUser{
		FirstNumber: firstNum,
		SecondNumber: secondNum,
		Operation: opNum,
	}

	 OperationForNumberss(client,res)

}
