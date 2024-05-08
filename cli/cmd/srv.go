package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	dylt "github.com/dylt-dev/dylt/lib"
)

func CreateSrvCommand () *cobra.Command {
	command := cobra.Command {
		Use: "srv $domain",
		Short: "DNS SRV records",
		Long: "Operations pertaining to DNS SRV records (https://www.ietf.org/rfc/rfc2782.txt)", 
		Args: cobra.MinimumNArgs(1),
		RunE: runSrvCommand,
	}
	command.Flags().BoolP("include-ips", "i", false, "Lookup IP Addresses of SRV records")

	return &command
}


func runSrvCommand (cmd *cobra.Command, args []string) error {
	domain := args[0]
	includeIps, _ := cmd.Flags().GetBool("include-ips")

	data := dylt.GetSrvs(domain, includeIps)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(string(jsonData))
	return nil
}