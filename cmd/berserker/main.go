package main

import (
	"berserker-network/fs"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const flagDataDir = "datadir"
const flagIP = "ip"
const flagPort = "port"

func main() {
	var bersCmd = &cobra.Command{
		Use:   "berserker",
		Short: "Berserker Network Ledger CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s.%s.%s-beta %s", Major, Minor, Fix, Verbal)
		},
	}

	bersCmd.AddCommand(migrateCmd())
	bersCmd.AddCommand(versionCmd)
	bersCmd.AddCommand(runCmd())
	bersCmd.AddCommand(balancesCmd())

	err := bersCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data dir where the DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func getDataDirFromCmd(cmd *cobra.Command) string {
	dataDir, _ := cmd.Flags().GetString(flagDataDir)

	return fs.ExpandPath(dataDir)
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
