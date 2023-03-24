package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Prameesh-P/calculator/calcpb"
)

func OperationForNumberss(c calcpb.ServicesClient,req *calcpb.AdditionRequestFromUser){

	stream,err:=c.ChatService(context.Background(),req)
	if err!=nil {
		log.Fatalf("Could not send request to server streaming>> : %v",err)
	}
	for {
			msg,err:=stream.Recv()
			if err ==io.EOF{
				break
			}
			if err !=nil{
				log.Fatalf("error while streaming %v",err)
			}
			log.Println(msg)
	}

	log.Println("Streaming is finished")	

}

func EnterNumber()int32{
	var num int32
	var err error
	log.Println("Enter the number for operation..\n 1) Addition \n 2) subtraction \n 3) multiplication \n 4) division")

	_,err=fmt.Scanln(&num)
	if err !=nil{
		log.Fatalf("err %v",err)
	}
	if num>4 || num <1{
		log.Fatalf("Foooooooooooollllllll")
	}
	return num
}
