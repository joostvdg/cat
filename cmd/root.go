package cmd

import (
	"fmt"
	"github.com/joostvdg/cat/api"
	"github.com/spf13/cobra"
	"os"
)

var ServerPort string
var PersistenceBackend string

func init() {
	rootCmd.AddCommand(versionCmd)
	serveCmd.Flags().StringVarP(&ServerPort, "port", "p", "7777", "Port to run the API server on")
	serveCmd.Flags().StringVarP(&PersistenceBackend, "persistenceBackend", "b", "mem", "The persistence backend to use [mem,etcd]")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs CAT's API server",
	Long:  `This will run the web API server of CAT, it's main function`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Serve(ServerPort, PersistenceBackend)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CAT",
	Long:  `All software has versions. This is CAT's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CAT 0.1.0")
	},
}

var rootCmd = &cobra.Command{
	Use:   "cat",
	Short: "cat is a small application tracker",
	Long:  `Yada yada yada`,
	Run: func(cmd *cobra.Command, args []string) {
		// return "0.1.0"
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
