package builder

// nginx
const (
	NginxVersion           = "1.11.1"
	NginxDownloadURLPrefix = "http://nginx.org/download"
)

// pcre
const (
	PcreVersion           = "8.38"
	PcreDownloadURLPrefix = "http://ftp.csx.cam.ac.uk/pub/software/programming/pcre"
)

// openssl
const (
	OpenSSLVersion           = "1.0.2h"
	OpenSSLDownloadURLPrefix = "http://www.openssl.org/source"
)

// zlib
const (
	ZlibVersion           = "1.2.8"
	ZlibDownloadURLPrefix = "http://zlib.net"
)

// openResty
const (
	OpenRestyVersion           = "1.9.15.1"
	OpenRestyDownloadURLPrefix = "https://openresty.org/download"
)

// tengine
const (
	TengineVersion           = "2.1.2"
	TengineDownloadURLPrefix = "http://tengine.taobao.org/download"
)

// component enumerations
const (
	ComponentNginx = iota
	ComponentOpenResty
	ComponentTengine
	ComponentPcre
	ComponentOpenSSL
	ComponentZlib
	ComponentMax
)
