package cmd

import (
	"fmt"
	"encoding/json"

	"github.com/spf13/cobra"

	dylt "github.com/dylt-dev/dylt/lib"
)

func CreateTxtCommand () *cobra.Command {
	command := cobra.Command {
		Use: "txt $host",
		Short: "DNS TXT record",
		Long: "Return the TXT records for the domain", 
		Args: cobra.MinimumNArgs(1),
		RunE: runTxtCommand,
	}
	// command.Flags().BoolP("include-ips", "i", false, "Lookup IP Addresses of SRV records")

	return &command
}


func runTxtCommand (cmd *cobra.Command, args []string) error {
	host := args[0]
	data := dylt.GetTxts(host)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
	return nil
}