package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/dylt-dev/diggo/cli/cmd"
)

func Run () int {
	rootCmd := createRootCommand()
	rootCmd.AddCommand(cmd.CreateSrvCommand())
	rootCmd.AddCommand(cmd.CreateCnameCommand())
	rootCmd.AddCommand(cmd.CreateTxtCommand())
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

func createRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "diggo",
		Short: "dig-like CLI",
		Long: "A dig-like CLI that does useful things dig doesn't. Like return JSON.",
	}
		
	return cmd
}
