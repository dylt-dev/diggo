package cmd

import (
	"fmt"
	"encoding/json"

	"github.com/spf13/cobra"

	dylt "github.com/dylt-dev/dylt/lib"
)

func CreateCnameCommand () *cobra.Command {
	command := cobra.Command {
		Use: "cname $host",
		Short: "DNS CNAME record",
		Long: "Return the canonical name for the host, possibly after recursivingly doing multiple CNAME lookups", 
		Args: cobra.MinimumNArgs(1),
		RunE: runCnameCommand,
	}
	// command.Flags().BoolP("include-ips", "i", false, "Lookup IP Addresses of SRV records")

	return &command
}


func runCnameCommand (cmd *cobra.Command, args []string) error {
	host := args[0]
	data := dylt.GetCname(host)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
	return nil
}