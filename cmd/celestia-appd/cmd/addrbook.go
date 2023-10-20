package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/celestiaorg/celestia-app/app"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/p2p/pex"
)

func addrbookCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "addrbook peers.txt addrbook.json",
		Short: "Convert a list of peers into an address book",
		Long: "Convert a list of peers into an address book.\n" +
			"The first argument (peers.txt) should contain a new line separated list of peers. The format for a peer is `id@ip:port` or `id@domain:port`.\n" +
			"The second argument (addrbook.json) should contain the complete file path, including both the directory path and the desired output file name, in the following format: `path/to/filename`. The address book is saved to the output file in JSON format.\n",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputFile := args[0]
			outputFile := args[1]

			data, err := os.ReadFile(inputFile)
			if err != nil {
				return err
			}
			lines := strings.Split(string(data), "\n")

			routabilityStrict := app.DefaultConsensusConfig().P2P.AddrBookStrict
			book := pex.NewAddrBook(outputFile, routabilityStrict)
			for _, line := range lines {
				if line == "" {
					continue
				}
				address, err := p2p.NewNetAddressString(line)
				if err != nil {
					return err
				}
				err = book.AddAddress(address, address)
				if err != nil {
					return err
				}
			}

			book.Save()
			fmt.Printf("Converted %s into %s\n", inputFile, outputFile)
			return nil
		},
	}

	return cmd
}
