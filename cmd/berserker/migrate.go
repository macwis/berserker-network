package main

import (
	"berserker-network/database"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var migrateCmd = func() *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrates the blockchain database according to new business rules.",
		Run: func(cmd *cobra.Command, args []string) {
			state, err := database.NewStateFromDisk(getDataDirFromCmd(cmd))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			block0 := database.NewBlock(
				database.Hash{},
				state.NextBlockNumber(),
				uint64(time.Now().Unix()),
				[]database.Tx{
					database.NewTx("odin", "odin", 3, ""),
					database.NewTx("odin", "odin", 700, "reward"),
				},
			)

			block0hash, err := state.AddBlock(block0)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			block1 := database.NewBlock(
				block0hash,
				state.NextBlockNumber(),
				uint64(time.Now().Unix()),
				[]database.Tx{
					database.NewTx("odin", "loki", 2000, ""),
					database.NewTx("odin", "odin", 100, "reward"),
					database.NewTx("loki", "odin", 1, ""),
					database.NewTx("loki", "balder", 1000, ""),
					database.NewTx("loki", "odin", 50, ""),
					database.NewTx("odin", "odin", 600, "reward"),
				},
			)

			block1hash, err := state.AddBlock(block1)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			block2 := database.NewBlock(
				block1hash,
				2,
				uint64(time.Now().Unix()),
				[]database.Tx{
					database.NewTx("odin", "odin", 24700, "reward"),
				},
			)

			_, err = state.AddBlock(block2)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	addDefaultRequiredFlags(migrateCmd)

	return migrateCmd
}
