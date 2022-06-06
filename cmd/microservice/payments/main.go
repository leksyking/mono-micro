package main

import "log"

func main() {
	log.Println("Starting the payments microservice")

	defer log.Println("Closing payments microservice")

	ctx := cmd.Context()

	paymentsInterface := createPaymentsMicroservice()
	if err := paymentsInterface.Run(ctx); err != nil {
		panic(err)
	}
}
