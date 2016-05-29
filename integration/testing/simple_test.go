package composeit

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	compose "github.com/libkermit/compose/testing"
	"github.com/libkermit/docker"
)

func TestSimpleProject(t *testing.T) {
	project := compose.CreateProject(t, "simple", "../assets/simple.yml")
	project.Start(t)

	container := project.Container(t, "hello")
	if container.Name != "/simple_hello_1" {
		t.Fatalf("expected name '/simple_hello_1', got %s", container.Name)
	}

	project.Stop(t)
}

func findContainersForProject(name string) ([]types.Container, error) {
	client, err := client.NewEnvClient()
	if err != nil {
		return []types.Container{}, err
	}
	filterArgs := filters.NewArgs()
	if filterArgs, err = filters.ParseFlag(docker.KermitLabelFilter, filterArgs); err != nil {
		return []types.Container{}, err
	}

	return client.ContainerList(context.Background(), types.ContainerListOptions{
		All:    true,
		Filter: filterArgs,
	})
}
