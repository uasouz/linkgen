package tests

import (
	"linkgen/pkg/tests"

	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ContainerMySQLImage = tests.Image{
		Name: tests.SelectMySQLImageByArch(),
		Ports: []tests.ContainerPort{{
			Port:     "3306",
			Protocol: "tcp",
		}},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "secret",
			"MYSQL_DATABASE":      "linkgen",
		},
		WaitStrategy: wait.ForListeningPort("3306/tcp"),
	}
)
