package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	ts "github.com/Subasri-V/application-new/netxd_customer/netxd"
	"github.com/Subasri-V/application-new/netxd_customer_controller/constants"
)

func main() {

	conn, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := ts.NewCustomerDetailsClient(conn)
	// response, err := client.CreateCustomer(context.Background(), &ts.CustomerRequest{
	// 	Customerid: 2,
	// 	Firstname:  "Gayatri",
	// 	Lastname:   "v",
	// 	Bankid:     1256,
	// 	Balance:    5000.0,
	// 	CreatedAt:  "",
	// 	UpdatedAt:  "",
	// 	IsActive:   true,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to call CreateProfile: %v", err)
	// }

	// fmt.Printf("Response: %s\n", response)

	// getRes,err:=client.GetCustomerById(context.Background(),&ts.IdReq{
	// 	Customerid: 1,
	// })
	// if err!=nil{
	// 	log.Fatalf("failed to get customer: %v",err)
	// }

	// fmt.Printf("Response: %s\n", getRes)

	updateRes,err:=client.DeleteCustomerById(context.Background(),&ts.DeleteReq{
		Customerid: 1,
	})
	if err!=nil{
		log.Fatalf("failed to get customer: %v",err)
	}

	fmt.Printf("Response: %s\n", updateRes)

	
}
