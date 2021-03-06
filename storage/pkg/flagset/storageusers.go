package flagset

import (
	"github.com/micro/cli/v2"
	"github.com/owncloud/ocis/storage/pkg/config"
)

// StorageUsersWithConfig applies cfg to the root flagset
func StorageUsersWithConfig(cfg *config.Config) []cli.Flag {
	flags := []cli.Flag{

		// debug ports are the odd ports
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9159",
			Usage:       "Address to bind debug server",
			EnvVars:     []string{"STORAGE_USERS_DEBUG_ADDR"},
			Destination: &cfg.Reva.StorageUsers.DebugAddr,
		},

		// Services

		// Storage home

		&cli.StringFlag{
			Name:        "grpc-network",
			Value:       "tcp",
			Usage:       "Network to use for the users storage, can be 'tcp', 'udp' or 'unix'",
			EnvVars:     []string{"STORAGE_USERS_GRPC_NETWORK"},
			Destination: &cfg.Reva.StorageUsers.GRPCNetwork,
		},
		&cli.StringFlag{
			Name:        "grpc-addr",
			Value:       "0.0.0.0:9157",
			Usage:       "GRPC Address to bind users storage",
			EnvVars:     []string{"STORAGE_USERS_GRPC_ADDR"},
			Destination: &cfg.Reva.StorageUsers.GRPCAddr,
		},
		&cli.StringFlag{
			Name:        "http-network",
			Value:       "tcp",
			Usage:       "Network to use for the storage service, can be 'tcp', 'udp' or 'unix'",
			EnvVars:     []string{"STORAGE_USERS_HTTP_NETWORK"},
			Destination: &cfg.Reva.StorageUsers.HTTPNetwork,
		},
		&cli.StringFlag{
			Name:        "http-addr",
			Value:       "0.0.0.0:9158",
			Usage:       "HTTP Address to bind users storage",
			EnvVars:     []string{"STORAGE_USERS_HTTP_ADDR"},
			Destination: &cfg.Reva.StorageUsers.HTTPAddr,
		},
		// TODO allow disabling grpc / http services
		/*
			&cli.StringSliceFlag{
				Name:    "grpc-service",
				Value:   cli.NewStringSlice("storageprovider"),
				Usage:   "--service storageprovider [--service otherservice]",
				EnvVars: []string{"STORAGE_USERS_GRPC_SERVICES"},
			},
			&cli.StringSliceFlag{
				Name:    "http-service",
				Value:   cli.NewStringSlice("dataprovider"),
				Usage:   "--service dataprovider [--service otherservice]",
				EnvVars: []string{"STORAGE_USERS_HTTP_SERVICES"},
			},
		*/

		&cli.StringFlag{
			Name:        "driver",
			Value:       "ocis",
			Usage:       "storage driver for users mount: eg. local, eos, owncloud, ocis or s3",
			EnvVars:     []string{"STORAGE_USERS_DRIVER"},
			Destination: &cfg.Reva.StorageUsers.Driver,
		},
		&cli.StringFlag{
			Name:        "mount-path",
			Value:       "/users",
			Usage:       "mount path",
			EnvVars:     []string{"STORAGE_USERS_MOUNT_PATH"},
			Destination: &cfg.Reva.StorageUsers.MountPath,
		},
		&cli.StringFlag{
			Name:        "mount-id",
			Value:       "1284d238-aa92-42ce-bdc4-0b0000009157", // /users
			Usage:       "mount id",
			EnvVars:     []string{"STORAGE_USERS_MOUNT_ID"},
			Destination: &cfg.Reva.StorageUsers.MountID,
		},
		&cli.BoolFlag{
			Name:        "expose-data-server",
			Value:       false,
			Usage:       "exposes a dedicated data server",
			EnvVars:     []string{"STORAGE_USERS_EXPOSE_DATA_SERVER"},
			Destination: &cfg.Reva.StorageUsers.ExposeDataServer,
		},
		&cli.StringFlag{
			Name:        "data-server-url",
			Value:       "http://localhost:9158/data",
			Usage:       "data server url",
			EnvVars:     []string{"STORAGE_USERS_DATA_SERVER_URL"},
			Destination: &cfg.Reva.StorageUsers.DataServerURL,
		},
		&cli.StringFlag{
			Name:        "http-prefix",
			Value:       "data",
			Usage:       "prefix for the http endpoint, without leading slash",
			EnvVars:     []string{"STORAGE_USERS_HTTP_PREFIX"},
			Destination: &cfg.Reva.StorageUsers.HTTPPrefix,
		},

		// some drivers need to look up users at the gateway

		// Gateway

		&cli.StringFlag{
			Name:        "gateway-endpoint",
			Value:       "localhost:9142",
			Usage:       "endpoint to use for the storage gateway service",
			EnvVars:     []string{"STORAGE_GATEWAY_ENDPOINT"},
			Destination: &cfg.Reva.Gateway.Endpoint,
		},
		// User provider

		&cli.StringFlag{
			Name:        "users-endpoint",
			Value:       "localhost:9144",
			Usage:       "endpoint to use for the storage service",
			EnvVars:     []string{"STORAGE_USERPROVIDER_ENDPOINT"},
			Destination: &cfg.Reva.Users.Endpoint,
		},
	}

	flags = append(flags, TracingWithConfig(cfg)...)
	flags = append(flags, DebugWithConfig(cfg)...)
	flags = append(flags, SecretWithConfig(cfg)...)
	flags = append(flags, DriverEOSWithConfig(cfg)...)
	flags = append(flags, DriverLocalWithConfig(cfg)...)
	flags = append(flags, DriverOwnCloudWithConfig(cfg)...)
	flags = append(flags, DriverOCISWithConfig(cfg)...)

	return flags
}
