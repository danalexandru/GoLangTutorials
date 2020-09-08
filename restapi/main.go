package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

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

func initCustomTerminal() *bufio.Scanner {
	fmt.Println("Custom terminal v0.1 started...")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	return scanner
}

func runCustomTerminal(scanner *bufio.Scanner) string {
	fmt.Printf("\n>> ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	command := scanner.Text()

	return command
}

func main() {

	ctx, computeService := getComputeService()

	// Project ID for this request.
	project := "tutorialproject-200903" // TODO: Update placeholder value.

	scanner := initCustomTerminal()
	var command string
	for {
		command = runCustomTerminal(scanner)

		if function, ok := Execute[command]; ok {
			result := function(ctx, computeService, project)
			fmt.Println(result)
		} else {
			fmt.Printf("Command not found: \"%s\".\n", command)
		}

	}

}
