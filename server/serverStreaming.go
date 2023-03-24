package main

import (
	"log"
	"time"

	"github.com/Prameesh-P/calculator/calcpb"
)

func (s *CalcServer) ChatService(req *calcpb.AdditionRequestFromUser, stm calcpb.Services_ChatServiceServer) error {

	var result []int32
	numberForOperation:=req.GetOperation()
	res:=checkOperation(numberForOperation,req)
	result = append(result, req.FirstNumber, req.SecondNumber, res)
	for i, v := range result {
		if i == 1 {
			log.Printf("Your First Number is >>>>> %d " ,req.FirstNumber )
			
		}
		time.Sleep(500*time.Millisecond)
		if i==1 {
			log.Printf("Your second Number is >>>>> %d ",req.SecondNumber)
		}
		if i==2{
			log.Printf("The sum is >>>>> %d ",v)
		}
		time.Sleep(1*time.Second)
	}
	return nil
}
func checkOperation(n int32,req *calcpb.AdditionRequestFromUser)int32{
	var ress int32
	if n==1{
		ress=req.FirstNumber+req.SecondNumber
	}else if n==2{
		ress=req.FirstNumber-req.SecondNumber
	}else if n==3{
		ress=req.FirstNumber*req.SecondNumber
	}else if n==4{
		ress=req.FirstNumber/req.SecondNumber
	}
	return ress
}