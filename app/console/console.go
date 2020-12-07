package console

import (
	"github.com/spf13/cobra"
	"log"
)

func InitConsole() {
	app := &cobra.Command{
		Use:   "tools",
		Short: "tools",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}

	app.AddCommand(ServerCmd)
	app.AddCommand(MigrateCmd)

	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
