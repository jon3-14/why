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
	"github.com/jon-314/why/internal/httpcheck"
	"github.com/jon-314/why/internal/icmp"
	"github.com/jon-314/why/internal/output"
	"github.com/jon-314/why/internal/summary"
	"github.com/jon-314/why/internal/tcp"
	"github.com/jon-314/why/internal/tls"
	"github.com/jon-314/why/internal/udp"
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
		output.Section(fmt.Sprintf("WHY report for %s", host))

		dnsResult := dns.Check(host)

		output.Section("DNS")

		if dnsResult.Success {
			output.Success("Resolved in %v", dnsResult.Duration)
		} else {
			output.Failure("DNS failed: %s", dnsResult.Error)
		}

		tcpResult := tcp.Check(host, 443)

		output.Section("TCP")

		if tcpResult.Success {
			output.Success("Connected to port 443 in %v", tcpResult.Duration)
		} else {
			output.Failure("TCP failed: %s", tcpResult.Error)
		}

		udpResult := udp.Check(host, 53)

		output.Section("UDP")

		if udpResult.Success {
			output.Success("UDP packet sent in %v", udpResult.Duration)

			if udpResult.Response {
				output.Success("Response received")
			} else {
				output.Warning("No response received")
			}
		} else {
			output.Failure("UDP failed: %s", udpResult.Error)
		}

		tlsResult := tls.Check(host, 443)

		output.Section("TLS")

		if tlsResult.Success {
			output.Success("Handshake successful in %v", tlsResult.Duration)
			output.Success("Version: %s", tlsResult.Version)
			output.Success("Certificate CN: %s", tlsResult.CommonName)

			if tlsResult.Expired {
				output.Failure("Certificate is expired")
			} else {
				output.Success(
					"Certificate valid for %v",
					tlsResult.ExpiresIn.Round(time.Hour*24),
				)
			}
		} else {
			output.Failure("TLS failed: %s", tlsResult.Error)
		}

		httpResult := httpcheck.Check(target)

		output.Section("HTTP")

		if httpResult.Success {
			output.Success("%s", httpResult.Status)
			output.Success("Response time: %v", httpResult.Duration)

			if httpResult.Redirects > 0 {
				output.Success("Redirects: %d", httpResult.Redirects)
			}
		} else {
			output.Failure("HTTP failed: %s", httpResult.Error)
		}

		icmpResult := icmp.Check(host)

		output.Section("ICMP")

		if icmpResult.Success {
			output.Success("Ping successful in %v", icmpResult.Duration)
		} else {
			output.Failure("Ping failed: %s", icmpResult.Error)
		}

		output.Section("Summary")

		summary.PrintTiming("DNS", dnsResult.Duration)
		summary.PrintTiming("TCP", tcpResult.Duration)
		summary.PrintTiming("UDP", udpResult.Duration)
		summary.PrintTiming("TLS", tlsResult.Duration)
		summary.PrintTiming("HTTP", httpResult.Duration)
		summary.PrintTiming("ICMP", icmpResult.Duration)
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
