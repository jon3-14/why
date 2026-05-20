package tls

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type Result struct {
	Success    bool
	Duration   time.Duration
	Version    string
	Error      string
	ExpiresIn  time.Duration
	Expired    bool
	CommonName string
}

func Check(host string, port int) Result {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	start := time.Now()

	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	conn, err := tls.DialWithDialer(
		dialer,
		"tcp",
		address,
		&tls.Config{
			ServerName: host,
		},
	)

	duration := time.Since(start)

	if err != nil {
		return Result{
			Success:  false,
			Error:    err.Error(),
			Duration: duration,
		}
	}

	defer func() {
		_ = conn.Close()
	}()

	state := conn.ConnectionState()

	cert := state.PeerCertificates[0]

	expiresIn := time.Until(cert.NotAfter)

	return Result{
		Success:    true,
		Duration:   duration,
		Version:    tlsVersion(state.Version),
		ExpiresIn:  expiresIn,
		Expired:    expiresIn < 0,
		CommonName: cert.Subject.CommonName,
	}
}

func tlsVersion(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown"
	}
}
