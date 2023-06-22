package cmd

import (
	"context"
	"fmt"
	"github.com/NubeIO/nrule/config"
	"github.com/NubeIO/nrule/logger"
	"github.com/NubeIO/nrule/server/constants"
	"github.com/NubeIO/nrule/server/router"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starting rubix-edge-wires",
	Long:  "it starts a server for edge-wires flow-engine",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	if err := config.Setup(RootCmd); err != nil {
		fmt.Errorf("error: %s", err) // here log is not setup yet...
	}
	logger.Init()
	if err := os.MkdirAll(config.Config.GetAbsDataDir(), os.FileMode(constants.Permission)); err != nil {
		panic(err)
	}

	logger.Logger.Infoln("starting rubix-edge-wires...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := router.Setup(ctx)

	host := "0.0.0.0"
	port := config.Config.GetPort()
	logger.Logger.Infof("server is starting at %s:%s", host, port)
	logger.Logger.Fatalf("%v", r.Run(fmt.Sprintf("%s:%s", host, port)))

}

func init() {
	RootCmd.AddCommand(serverCmd)
}
