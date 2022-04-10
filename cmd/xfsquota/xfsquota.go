package main

import (
	"fmt"
	"github.com/silenceper/xfsquota/pkg/xfsquota"
	"github.com/spf13/cobra"
)

const (
	version = "v0.0.2"
)

var (
	xfsQuota *xfsquota.XfsQuota
	rootCmd  = &cobra.Command{
		Use:   "xfsquota",
		Short: "xfsquota is a tool for managing XFS quotas",
	}
)

var getQuotaCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get quota information",
	Example: "xfsquota get /home/user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		quotaRes, err := xfsQuota.GetQuota(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("quota Size(bytes):", quotaRes.Quota)
		fmt.Println("quota Inodes:", quotaRes.Inodes)
		fmt.Println("diskUsage Size(bytes):", quotaRes.QuotaUsed)
		fmt.Println("diskUsage Inodes:", quotaRes.InodesUsed)
	},
}

var size string
var inodes string
var setQuotaCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set quota information",
	Example: "xfsquota set /home/user -s 100MiB -i 100",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		err := xfsQuota.SetQuota(args[0], size, inodes)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("set quota success, path: %s, size:%s, inodes:%s\n", args[0], size, inodes)
	},
}

var cleanQuotaCmd = &cobra.Command{
	Use:     "clean",
	Short:   "clean quota information",
	Example: "xfsquota clean /home/user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		err := xfsQuota.CleanQuota(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("clean quota success, path:", args[0])
	},
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "get version",
	Example: "xfsquota version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("xfsquota version:", version)
	},
}

func init() {
	setQuotaCmd.Flags().StringVarP(&size, "size", "s", "0", "quota size")
	setQuotaCmd.Flags().StringVarP(&inodes, "inodes", "i", "0", "quota inodes")
	rootCmd.AddCommand(getQuotaCmd)
	rootCmd.AddCommand(setQuotaCmd)
	rootCmd.AddCommand(cleanQuotaCmd)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	xfsQuota = xfsquota.NewXfsQuota()
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
