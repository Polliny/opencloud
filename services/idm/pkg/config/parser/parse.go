package parser

import (
	"errors"

	occfg "github.com/opencloud-eu/opencloud/pkg/config"
	"github.com/opencloud-eu/opencloud/pkg/shared"
	"github.com/opencloud-eu/opencloud/services/idm/pkg/config"
	"github.com/opencloud-eu/opencloud/services/idm/pkg/config/defaults"

	"github.com/opencloud-eu/opencloud/pkg/config/envdecode"
)

// ParseConfig loads configuration from known paths.
func ParseConfig(cfg *config.Config) error {
	err := occfg.BindSourcesToStructs(cfg.Service.Name, cfg)
	if err != nil {
		return err
	}

	defaults.EnsureDefaults(cfg)
	// load all env variables relevant to the config in the current context.
	if err := envdecode.Decode(cfg); err != nil {
		// no environment variable set for this config is an expected "error"
		if !errors.Is(err, envdecode.ErrNoTargetFieldsAreSet) {
			return err
		}
	}

	defaults.Sanitize(cfg)

	return Validate(cfg)
}

func Validate(cfg *config.Config) error {
	if cfg.CreateDemoUsers && cfg.AdminUserID == "" {
		return shared.MissingAdminUserID(cfg.Service.Name)
	}

	if cfg.ServiceUserPasswords.Idm == "" {
		return shared.MissingServiceUserPassword(cfg.Service.Name, "IDM")
	}

	if cfg.AdminUserID != "" && cfg.ServiceUserPasswords.OCAdmin == "" {
		return shared.MissingServiceUserPassword(cfg.Service.Name, "admin")
	}

	if cfg.ServiceUserPasswords.Idp == "" {
		return shared.MissingServiceUserPassword(cfg.Service.Name, "IDP")
	}

	if cfg.ServiceUserPasswords.Reva == "" {
		return shared.MissingServiceUserPassword(cfg.Service.Name, "REVA")
	}

	return nil
}
