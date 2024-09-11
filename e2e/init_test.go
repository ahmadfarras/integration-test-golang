package e2e_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var network *dockertest.Network

func TestMain(m *testing.M) {
	// Init Docker pool.
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %v", err)
	}

	// Create a Docker network for the tests.
	network, err = pool.CreateNetwork("test-network")
	if err != nil {
		log.Fatalf("Could not create network: %v", err)
	}

}

func deployMysql(pool *dockertest.Pool) (*dockertest.Resource, error) {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Hostname:     "mysql-container",
		Repository:   "mysql",
		Tag:          "5.7",
		ExposedPorts: []string{"3306"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306/tcp": {
				{HostIP: "", HostPort: "3306"},
			},
		},
		Networks: []*dockertest.Network{
			network,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("could not start resource: %v", err)

	}

	if err := pool.Retry(func() error {
		log.Print
		return nil
	}); err != nil {
		return nil, fmt.Errorf("could not connect to Docker: %v", err)

	}
}
