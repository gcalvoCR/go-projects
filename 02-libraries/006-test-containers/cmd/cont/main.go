package main

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type Bread struct {
	num int
}

func (b Bread) Init() {
	fmt.Println("bread initialized")
}

func init() {
	fmt.Println("Miau, I'm the init func being called!")
}

func initContainer() {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Could not start redis: %s", err)
	}

	name, _ := redisC.Name(ctx)
	port, _ := redisC.MappedPort(ctx, "6379/tcp")

	log.Println("INFO", name, port)

	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			log.Fatalf("Could not stop redis: %s", err)
		}
	}()
}

func main() {
	fmt.Println("Hello world fro the server")
	b := Bread{}

	b.Init()
	initContainer()
}
