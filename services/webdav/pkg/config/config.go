package config

import (
	"context"

	"github.com/opencloud-eu/opencloud/pkg/shared"
	"go-micro.dev/v4/client"
)

// Config combines all available configuration parts.
type Config struct {
	Commons *shared.Commons `yaml:"-"` // don't use this directly as configuration for a service

	Service Service `yaml:"-"`

	Tracing *Tracing `yaml:"tracing"`
	Log     *Log     `yaml:"log"`
	Debug   Debug    `yaml:"debug"`

	GRPCClientTLS *shared.GRPCClientTLS `yaml:"grpc_client_tls"`
	GrpcClient    client.Client         `yaml:"-"`

	HTTP HTTP `yaml:"http"`

	DisablePreviews      bool            `yaml:"disablePreviews" env:"OC_DISABLE_PREVIEWS;WEBDAV_DISABLE_PREVIEWS" desc:"Set this option to 'true' to disable rendering of thumbnails triggered via webdav access. Note that when disabled, all access to preview related webdav paths will return a 404." introductionVersion:"1.0.0"`
	OpenCloudPublicURL   string          `yaml:"opencloud_public_url" env:"OC_URL;OC_PUBLIC_URL" desc:"URL, where OpenCloud is reachable for users." introductionVersion:"1.0.0"`
	WebdavNamespace      string          `yaml:"webdav_namespace" env:"WEBDAV_WEBDAV_NAMESPACE" desc:"CS3 path layout to use when forwarding /webdav requests" introductionVersion:"1.0.0"`
	RevaGateway          string          `yaml:"reva_gateway" env:"OC_REVA_GATEWAY" desc:"CS3 gateway used to look up user metadata" introductionVersion:"1.0.0"`
	RevaGatewayTLSMode   string          `yaml:"reva_gateway_tls_mode" env:"OC_REVA_GATEWAY_TLS_MODE" desc:"TLS mode for grpc connection to the CS3 gateway endpoint. Possible values are 'off', 'insecure' and 'on'. 'off': disables transport security for the clients. 'insecure' allows using transport security, but disables certificate verification (to be used with the autogenerated self-signed certificates). 'on' enables transport security, including server certificate verification." introductionVersion:"1.0.0"`
	RevaGatewayTLSCACert string          `yaml:"reva_gateway_tls_cacert" env:"OC_REVA_GATEWAY_TLS_CACERT" desc:"The root CA certificate used to validate the gateway's TLS certificate." introductionVersion:"1.0.0"`
	Context              context.Context `yaml:"-"`
}
