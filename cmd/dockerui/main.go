package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	fmt.Println("hello, world!!")

	os.Setenv("DOCKER_API_VERSION", "1.35")

	cl, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	containers, err := cl.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})

	for _, c := range containers {
		println(c.ID)
	}

	println(cl.ClientVersion())

	images, err := cl.ImageList(ctx, types.ImageListOptions{
		All: true,
	})

	for _, img := range images {
		println(img.ID)
	}
}
