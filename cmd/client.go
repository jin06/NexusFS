package cmd

import (
	"fmt"
	"github.com/jin06/NexusFS/config"
	"github.com/jin06/NexusFS/nfs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run as a client.",
	Long:  "",
	Run:   runClient,
}

func runClient(cmd *cobra.Command, args []string) {
	clusterName := viper.GetString("cluster_name")
	fmt.Println(clusterName)
	fmt.Println("hello")
	mt := viper.GetString("mountpoint")
	ld := viper.GetString("lookback_dir")
	sType := viper.GetString("storage_type")
	fuseName := config.Def.ClusterName + "_" + config.Def.NodeName

	sfs, err := nfs.New(
		nfs.MountPoint(mt),
		nfs.LBPath(ld),
		nfs.FSType(sType),
		nfs.FuseName(fuseName),
	)
	if err != nil {
		panic(err)
	}
	err = sfs.Mount()
	panic(err)
}
