package flagset

import (
	"github.com/micro/cli"
	"github.com/owncloud/ocis-devldap/pkg/config"
)

// RootWithConfig applies cfg to the root flagset
func RootWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "config-file",
			Value:       "",
			Usage:       "Path to config file",
			EnvVar:      "DEVLDAP_CONFIG_FILE",
			Destination: &cfg.File,
		},
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "Set logging level",
			EnvVar:      "DEVLDAP_LOG_LEVEL",
			Destination: &cfg.Log.Level,
		},
	}
}

// HealthWithConfig applies cfg to the root flagset
func HealthWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:8390",
			Usage:       "Address to debug endpoint",
			EnvVar:      "DEVLDAP_DEBUG_ADDR",
			Destination: &cfg.Debug.Addr,
		},
	}
}

// ServerWithConfig applies cfg to the root flagset
func ServerWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Usage:       "Enable sending traces",
			EnvVar:      "DEVLDAP_TRACING_ENABLED",
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-type",
			Value:       "jaeger",
			Usage:       "Tracing backend type",
			EnvVar:      "DEVLDAP_TRACING_TYPE",
			Destination: &cfg.Tracing.Type,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "Endpoint for the agent",
			EnvVar:      "DEVLDAP_TRACING_ENDPOINT",
			Destination: &cfg.Tracing.Endpoint,
		},
		&cli.StringFlag{
			Name:        "tracing-collector",
			Value:       "",
			Usage:       "Endpoint for the collector",
			EnvVar:      "DEVLDAP_TRACING_COLLECTOR",
			Destination: &cfg.Tracing.Collector,
		},
		&cli.StringFlag{
			Name:        "tracing-service",
			Value:       "devldap",
			Usage:       "Service name for tracing",
			EnvVar:      "DEVLDAP_TRACING_SERVICE",
			Destination: &cfg.Tracing.Service,
		},
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:8390",
			Usage:       "Address to bind debug server",
			EnvVar:      "DEVLDAP_DEBUG_ADDR",
			Destination: &cfg.Debug.Addr,
		},
		&cli.StringFlag{
			Name:        "debug-token",
			Value:       "",
			Usage:       "Token to grant metrics access",
			EnvVar:      "DEVLDAP_DEBUG_TOKEN",
			Destination: &cfg.Debug.Token,
		},
		&cli.BoolFlag{
			Name:        "debug-pprof",
			Usage:       "Enable pprof debugging",
			EnvVar:      "DEVLDAP_DEBUG_PPROF",
			Destination: &cfg.Debug.Pprof,
		},
		&cli.BoolFlag{
			Name:        "debug-zpages",
			Usage:       "Enable zpages debugging",
			EnvVar:      "DEVLDAP_DEBUG_ZPAGES",
			Destination: &cfg.Debug.Zpages,
		},
		&cli.StringFlag{
			Name:        "ldap-addr",
			Value:       "0.0.0.0:10389",
			Usage:       "Address to bind ldap server",
			EnvVar:      "DEVLDAP_LDAP_ADDR",
			Destination: &cfg.LDAP.Addr,
		},
		&cli.StringFlag{
			Name:        "data-file",
			Value:       "data.json",
			Usage:       "Path to ldap data file",
			EnvVar:      "DEVLDAP_DATA_FILE",
			Destination: &cfg.LDAP.Data,
		},
		&cli.StringFlag{
			Name:        "asset-path",
			Value:       "",
			Usage:       "Path to custom assets",
			EnvVar:      "DEVLDAP_ASSET_PATH",
			Destination: &cfg.Asset.Path,
		},
	}
}
