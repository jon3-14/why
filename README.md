# why

`why` is a CLI troubleshooting assistant for diagnosing DNS, TCP, UDP, TLS, and HTTP connectivity issues.

Instead of manually running multiple tools like:

* `dig`
* `curl`
* `openssl`
* `nc`

`why` orchestrates diagnostics automatically and summarizes where communication is failing.

---

# Features

* DNS resolution checks
* TCP connectivity testing
* UDP diagnostics
* TLS handshake and certificate validation
* HTTP response analysis
* Timing summaries
* Cross-platform binaries for:

  * macOS
  * Linux
  * Windows

---

# Installation

Download the latest release from:

https://github.com/jon-314/why/releases

---

## macOS / Linux

Download the appropriate archive for your platform:

* `why_*_darwin_arm64.tar.gz`
* `why_*_darwin_amd64.tar.gz`
* `why_*_linux_amd64.tar.gz`

Extract the archive:

```bash
tar -xzf why_*.tar.gz
```

Make the binary executable:

```bash
chmod +x why
```

(Optional) Move it into your PATH:

```bash
sudo mv why /usr/local/bin/why
```

Then run:

```bash
why https://example.com
```

---

## Windows

Download:

```txt
why_*_windows_amd64.zip
```

Extract the ZIP archive and run:

```powershell
.\why.exe https://example.com
```

---

# Example

```bash
why https://example.com
```

Example output:

```txt
DNS
✓ Resolved in 14ms

TCP
✓ Connected to port 443 in 42ms

TLS
✓ Handshake successful
✓ TLS 1.3 negotiated

HTTP
✓ 200 OK

Summary
DNS   14ms
TCP   42ms
TLS   91ms
HTTP  201ms
```

# License

MIT
