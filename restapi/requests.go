package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
)

// GetAcceleratorTypesList - this api request returns all the accelerator types
func GetAcceleratorTypesList(ctx context.Context, computeService *compute.Service, project string, params map[string]string) string {
	result := ""
	req := computeService.AcceleratorTypes.AggregatedList(project)
	if err := req.Pages(ctx, func(page *compute.AcceleratorTypeAggregatedList) error {
		for name, acceleratorTypesScopedList := range page.Items {
			// TODO: Change code below to process each (name: acceleratorTypesScopedList) element:
			result = fmt.Sprintf("%s%v: %#v\n", result, name, acceleratorTypesScopedList)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("{\n%s\n}", result)
}

// GetVMInstacesList - this api request reutrns all the virtual-machine instances
func GetVMInstacesList(ctx context.Context, computeService *compute.Service, project string, params map[string]string) string {
	result := ""
	req := computeService.Instances.List(project, params["zone"])
	if err := req.Pages(ctx, func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			// TODO: Change code below to process each `instance` resource:
			result = fmt.Sprintf("%s%#v\n", result, instance)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return result
}
