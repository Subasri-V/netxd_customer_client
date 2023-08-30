package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Subasri-V/application-new/netxd_customer_controller/constants"
	"google.golang.org/grpc"

	ts "github.com/Subasri-V/application-new/netxd_customer/netxd"
)

func main(){

	conn, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client:=ts.NewCustomerDetailsClient(conn)
	response,err:=client.CreateCustomer(context.Background(),&ts.CustomerRequest{
		Customerid: 1,
		Firstname:  "Subasri",
		Lastname:   "v",
		Bankid:     125,
		Balance:    500.0,
		CreatedAt:  "",
		UpdatedAt:  "",
		IsActive:   true,
	})
	
	if err != nil {
		log.Fatalf("Failed to call CreateProfile: %v", err)
	}

	fmt.Printf("Response: %s\n", response)
}
