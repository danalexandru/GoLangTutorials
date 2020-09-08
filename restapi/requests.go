package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
)

// GetAcceleratorTypesList - this api request returns all the accelerator types
func GetAcceleratorTypesList(ctx context.Context, computeService *compute.Service, project string) string {
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
