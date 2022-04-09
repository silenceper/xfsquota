package main

import (
	"fmt"
	"github.com/silenceper/xfsquota/pkg/xfsquota"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var (
	debug    bool
	basePath string
	xfsQuota *xfsquota.XfsQuota
	rootCmd  = &cobra.Command{
		Use:   "xfsquota",
		Short: "xfsquota is a tool for managing XFS quotas",
	}
)

var getQuotaCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get quota information",
	Example: "xfsquota -b /data get /home/user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		err := xfsQuota.Init(basePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		quotaRes, err := xfsQuota.GetQuota(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("quota Size(bytes):", quotaRes.Quota.Size)
		fmt.Println("quota Inodes:", quotaRes.Quota.Inodes)
		fmt.Println("diskUsage Size(bytes):", quotaRes.DiskUsage.Size)
		fmt.Println("diskUsage Inodes:", quotaRes.DiskUsage.InodeCount)
	},
}

var size string
var inodes string
var setQuotaCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set quota information",
	Example: "xfsquota -b /data/ set /home/user -s 100MiB -i 100",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		err := xfsQuota.Init(basePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = xfsQuota.SetQuota(args[0], size, inodes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("set quota success, path: %s, size:%s, inodes:%s\n", args[0], size, inodes)
	},
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	rootCmd.PersistentFlags().StringVarP(&basePath, "basePath", "b", "./", "base path for backing filesystem device")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
	setQuotaCmd.Flags().StringVarP(&size, "size", "s", "0", "quota size")
	setQuotaCmd.Flags().StringVarP(&inodes, "inodes", "i", "0", "quota inodes")
	rootCmd.AddCommand(getQuotaCmd)
	rootCmd.AddCommand(setQuotaCmd)
}

func main() {
	if debug {
		log.SetLevel(log.DebugLevel)
	}
	xfsQuota = xfsquota.NewXfsQuota()
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
