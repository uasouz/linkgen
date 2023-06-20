package tests

import (
	"context"
	"fmt"
	"linkgen/pkg/collection"
	"log"
	"strings"

	"github.com/containerd/containerd/platforms"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type ContainerPort struct {
	Port     string
	Protocol string
}

type Image struct {
	Name  string
	Ports []ContainerPort

	Env          map[string]string
	Mounts       testcontainers.ContainerMounts
	Cmd          []string
	Entrypoint   []string
	WaitStrategy wait.Strategy
}

type Container struct {
	testcontainers.Container
	URIS []string
}

func Setup(ctx context.Context, image Image) (*Container, error) {
	cont, uris, err := prepareContainer(ctx, image)
	if err != nil {
		return nil, err
	}
	return &Container{Container: cont, URIS: uris}, nil
}

func prepareContainer(ctx context.Context, image Image) (testcontainers.Container, []string, error) {

	containerPorts := collection.Map(image.Ports, func(port ContainerPort) string {
		return port.Port + "/" + port.Protocol
	})

	req := testcontainers.ContainerRequest{
		Image:         image.Name,
		Env:           image.Env,
		Mounts:        image.Mounts,
		Cmd:           image.Cmd,
		Entrypoint:    image.Entrypoint,
		ExposedPorts:  containerPorts,
		WaitingFor:    image.WaitStrategy,
		ImagePlatform: strings.Replace(platforms.DefaultString(), "darwin", "linux", 1),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		if container != nil {
			reader, _ := container.Logs(ctx)
			var buf []byte
			reader.Read(buf)
			fmt.Println(string(buf))
		}
		return nil, nil, err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, nil, err
	}

	var uris []string

	for i := range image.Ports {
		mappedPort, err := container.MappedPort(ctx, nat.Port(image.Ports[i].Port))
		if err != nil {
			return nil, nil, err
		}

		var uri string
		uri = fmt.Sprintf("%s:%s", hostIP, mappedPort.Port())
		uris = append(uris, uri)
	}

	log.Printf("TestContainers: Container %s is now running at %s\n", req.Image, uris)
	return container, uris, nil
}
