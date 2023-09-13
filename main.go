package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"Hermes/performance"
)

var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "Welcome!",
	Long:  "Welcome! :)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the program! :)")
	},
}

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Get Cpu information",
	Long:  "Get information about CPU usage of your system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(performance.CpuInfo())
	},
}

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Get memory information",
	Long:  "Get information about memory usage of your system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(performance.MemInfo())
	},
}

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get disk information",
	Long:  "Get information about disk usage of your system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(performance.DiskInfo())
	},
}

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Get network information",
	Long:  "Get information about network usage of your system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(performance.NetInfo())
	},
}

var sysCmd = &cobra.Command{
	Use:   "sys",
	Short: "Get system information",
	Long:  "Get general information about your system such as OS version and architecture",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(performance.SysInfo())
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
	rootCmd.AddCommand(memCmd)
	rootCmd.AddCommand(diskCmd)
	rootCmd.AddCommand(netCmd)
	rootCmd.AddCommand(sysCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
