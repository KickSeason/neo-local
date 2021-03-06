package stack

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

// NewNeoScanSync creates a new service for the
// registry.gitlab.com/cityofzion/neo-scan/sync:latest image.
func NewNeoScanSync() Service {
	return Service{
		Author: "cityofzion",
		ContainerConfig: &container.Config{
			Env: []string{
				"DB_DATABASE=neoscan_prodv",
				"DB_HOSTNAME=postgres",
				"DB_PASSWORD=postgres",
				"DB_USERNAME=postgres",
				"NEOSCAN=neo-scan-api:4000",
				"NEO_NOTIFICATIONS_SERVER='http://notifications-server:8080/v1'",
				"NEO_SEEDS='http://neo-privatenet:30333;http://neo-privatenet:30334;http://neo-privatenet:30335;http://neo-privatenet:30336'",
				"REPLACE_OS_VARS=true",
			},
			ExposedPorts: map[nat.Port]struct{}{
				"4002/tcp": {},
			},
		},
		DependsOn: []string{
			"neo-privatenet",
			"notifications-server",
			"postgres",
		},
		HostConfig: &container.HostConfig{
			Links: []string{
				"neo-privatenet:30333",
				"neo-privatenet:30334",
				"neo-privatenet:30335",
				"neo-privatenet:30336",
			},
			Privileged: false,
		},
		Image: "registry.gitlab.com/cityofzion/neo-scan/sync",
		Name:  "neo-scan-sync",
		Tag:   "latest",
	}
}
