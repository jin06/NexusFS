package cmd

import (
	"fmt"
	"github.com/jin06/NexusFS/api"
	"github.com/jin06/NexusFS/api/handlers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run as a server.",
	Long: ` 
	todo
	`,
	Run: runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	fmt.Println("hello, server running")
	addr := viper.GetString("listen_address")
	logrus.Debug("ruuning addr: ", addr)
	err := runApiServer(addr)
	panic(err)
}

func runApiServer(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	srv := api.NewServer()
	path := viper.GetString("local_path")
	fsService := handlers.NewFS(path)
	gServer := grpc.NewServer()
	api.RegisterHelloServer(gServer, srv)
	api.RegisterFSServer(gServer, fsService)

	return gServer.Serve(listener)
}
