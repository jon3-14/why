/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/jon-314/why/internal/dns"
	httpcheck "github.com/jon-314/why/internal/http"
	"github.com/jon-314/why/internal/summary"
	"github.com/jon-314/why/internal/tcp"
	"github.com/jon-314/why/internal/tls"
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

		fmt.Println()

		tlsResult := tls.Check(host, 443)

		fmt.Println("TLS")

		if tlsResult.Success {
			fmt.Printf("✓ Handshake successful in %v\n", tlsResult.Duration)
			fmt.Printf("✓ Version: %s\n", tlsResult.Version)
			fmt.Printf("✓ Certificate CN: %s\n", tlsResult.CommonName)

			if tlsResult.Expired {
				fmt.Println("✗ Certificate is expired")
			} else {
				fmt.Printf("✓ Certificate valid for %v\n", tlsResult.ExpiresIn.Round(time.Hour*24))
			}
		} else {
			fmt.Printf("✗ TLS failed: %s\n", tlsResult.Error)
		}

		fmt.Println()

		httpResult := httpcheck.Check(target)

		fmt.Println("HTTP")

		if httpResult.Success {
			fmt.Printf("✓ %s\n", httpResult.Status)
			fmt.Printf("✓ Response time: %v\n", httpResult.Duration)

			if httpResult.Redirects > 0 {
				fmt.Printf("✓ Redirects: %d\n", httpResult.Redirects)
			}
		} else {
			fmt.Printf("✗ HTTP failed: %s\n", httpResult.Error)
		}

		fmt.Println()
		fmt.Println("Summary")
		fmt.Println()

		summary.PrintTiming("DNS", result.Duration)
		summary.PrintTiming("TCP", tcpResult.Duration)
		summary.PrintTiming("TLS", tlsResult.Duration)
		summary.PrintTiming("HTTP", httpResult.Duration)
		fmt.Println()
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
