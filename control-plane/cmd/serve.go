package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"senkawa.moe/haa-chan/app/web"
)

var (
	listenAddr     string
	grpcListenAddr string
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Long:  `Serves up the web server`,
	Run: func(cmd *cobra.Command, args []string) {
		app := web.NewApplication()
		err := web.RunGRPCServer(app, grpcListenAddr)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}

		log.Fatal(app.Server.Listen(listenAddr))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&listenAddr, "address", "localhost:8000", "Listen address")
	serveCmd.Flags().StringVar(&grpcListenAddr, "grpc-address", "localhost:8001", "gRPC Listen address")
}
