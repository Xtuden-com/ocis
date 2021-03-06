// +build !simple

package command

import (
	"github.com/micro/cli/v2"
	"github.com/owncloud/ocis/accounts/pkg/command"
	svcconfig "github.com/owncloud/ocis/accounts/pkg/config"
	"github.com/owncloud/ocis/accounts/pkg/flagset"
	"github.com/owncloud/ocis/ocis/pkg/config"
	"github.com/owncloud/ocis/ocis/pkg/register"
	"github.com/owncloud/ocis/ocis/pkg/version"
)

// AccountsCommand is the entrypoint for the accounts command.
func AccountsCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "accounts",
		Usage:    "Start accounts server",
		Category: "Extensions",
		Flags:    flagset.ServerWithConfig(cfg.Accounts),
		Subcommands: []*cli.Command{
			command.ListAccounts(cfg.Accounts),
			command.AddAccount(cfg.Accounts),
			command.UpdateAccount(cfg.Accounts),
			command.RemoveAccount(cfg.Accounts),
			command.InspectAccount(cfg.Accounts),
			command.PrintVersion(cfg.Accounts),
		},
		Action: func(c *cli.Context) error {
			accountsCommand := command.Server(configureAccounts(cfg))
			if err := accountsCommand.Before(c); err != nil {
				return err
			}

			return cli.HandleAction(accountsCommand.Action, c)
		},
	}
}

func configureAccounts(cfg *config.Config) *svcconfig.Config {
	cfg.Accounts.Log.Level = cfg.Log.Level
	cfg.Accounts.Log.Pretty = cfg.Log.Pretty
	cfg.Accounts.Log.Color = cfg.Log.Color
	cfg.Accounts.Server.Version = version.String

	// TODO: we need tracing on the accounts service as well. when we have it, apply default config from OCIS here.

	if cfg.TokenManager.JWTSecret != "" {
		cfg.Accounts.TokenManager.JWTSecret = cfg.TokenManager.JWTSecret
		cfg.Accounts.Repo.CS3.JWTSecret = cfg.TokenManager.JWTSecret
	}

	return cfg.Accounts
}

func init() {
	register.AddCommand(AccountsCommand)
}
