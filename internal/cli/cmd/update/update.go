package update

import (
	"github.com/spf13/cobra"

	"github.com/akuity/kargo/internal/cli/config"
	"github.com/akuity/kargo/internal/cli/option"
)

func NewCommand(cfg config.CLIConfig, opt *option.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a freight alias",
	}
	cmd.AddCommand(newUpdateFreightAliasCommand(cfg, opt))
	return cmd
}
