/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/jon-314/why/internal/dns"
	"github.com/jon-314/why/internal/tcp"
	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
    Use:   "why <target>",
    Short: "Explain connectivity issues",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
	target := args[0]

	parsed, err := url.Parse(target)
	if err != nil {
		fmt.Println("Invalid URL:", err)
		return
	}

	host := parsed.Hostname()

	result := dns.Check(host)

	fmt.Println("DNS Check")

	if result.Success {
		fmt.Printf("✓ Resolved in %v\n", result.Duration)

		for _, ip := range result.IPs {
			fmt.Printf("  - %s\n", ip)
		}
	} else {
		fmt.Printf("✗ DNS failed: %s\n", result.Error)
	}

	fmt.Println()

tcpResult := tcp.Check(host, 443)

fmt.Println("TCP")

if tcpResult.Success {
	fmt.Printf("✓ Connected to port 443 in %v\n", tcpResult.Duration)
} else {
	fmt.Printf("✗ TCP failed: %s\n", tcpResult.Error)
}
},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.why.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


