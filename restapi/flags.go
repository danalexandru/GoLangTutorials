package main

import (
	"context"

	"google.golang.org/api/compute/v1"
)

var (
	// Execute - a map containing the equivalent functions that do api requests to gpc
	Execute = map[string]func(ctx context.Context, computeService *compute.Service, project string, params map[string]string) string{
		"gcloud compute accelerator-types list": GetAcceleratorTypesList,
		"gcloud compute instances list":         GetVMInstacesList,
		"gcloud compute networks list":          GetNetworksList,
	}
)
