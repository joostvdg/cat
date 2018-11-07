package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ServerPort string
var PersistenceBackendType string

func init() {
	rootCmd.AddCommand(versionCmd)

    serveGrpcCmd.Flags().StringVarP(&ServerPort, "port", "p", "9090", "Port to run the gRPC API server on")
    serveGrpcCmd.Flags().StringVarP(&PersistenceBackendType, "persistenceBackend", "b", "mem", "The persistence backend to use [mem,etcd]")
    rootCmd.AddCommand(serveGrpcCmd)

	serveHttpCmd.Flags().StringVarP(&ServerPort, "port", "p", "7777", "Port to run the HTTP API server on")
	serveHttpCmd.Flags().StringVarP(&PersistenceBackendType, "persistenceBackend", "b", "mem", "The persistence backend to use [mem,etcd]")
	rootCmd.AddCommand(serveHttpCmd)
}

var serveHttpCmd = &cobra.Command{
	Use:   "http",
	Short: "Runs CAT's API server (HTTP)",
	Long:  `This will run the http API server of CAT, it's main function`,
	Run: func(cmd *cobra.Command, args []string) {
        StartWebserver(ServerPort, PersistenceBackendType)
	},
}

var serveGrpcCmd = &cobra.Command{
    Use:   "grpc",
    Short: "Runs CAT's API server (GRPC)",
    Long:  `This will run the gRPC API server of CAT, it's main function`,
    Run: func(cmd *cobra.Command, args []string) {
        StartGRPCServer(PersistenceBackendType, ServerPort)
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
