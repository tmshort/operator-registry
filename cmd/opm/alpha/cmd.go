package alpha

import (
	"github.com/spf13/cobra"

	"github.com/operator-framework/operator-registry/cmd/opm/alpha/bundle"
	"github.com/operator-framework/operator-registry/cmd/opm/alpha/list"
	rendergraph "github.com/operator-framework/operator-registry/cmd/opm/alpha/render-graph"
	"github.com/operator-framework/operator-registry/cmd/opm/alpha/template"
)

func NewCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Hidden: true,
		Use:    "alpha",
		Short:  "Run an alpha subcommand",
		Args:   cobra.NoArgs,
		Run:    func(_ *cobra.Command, _ []string) {}, // adding an empty function here to preserve non-zero exit status for misstated subcommands/flags for the command hierarchy
	}

	runCmd.AddCommand(
		bundle.NewCmd(),
		list.NewCmd(),
		rendergraph.NewCmd(),
		template.NewCmd(),
	)
	return runCmd
}
