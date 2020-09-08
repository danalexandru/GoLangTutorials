package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

func getComputeService() (context.Context, *compute.Service) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	return ctx, computeService
}

func main() {
	ctx, computeService := getComputeService()

	// Project ID for this request.
	project := "tutorialproject-200903" // TODO: Update placeholder value.
	result := Execute["gcloud compute accelerator-types list"](ctx, computeService, project)

	fmt.Println(result)
}
