package framework

import (
	"crypto/tls"
	"time"
)

// defaultTLSConfig is the default TLS config
// that will be used if no TLS config is provided but
// use TLS is true.
var defaultTLSConfig = tls.Config{
	PreferServerCipherSuites: true,
	CurvePreferences: []tls.CurveID{
		tls.CurveP256,
		tls.X25519,
	},
	MinVersion: tls.VersionTLS12,
	CipherSuites: []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	},
}

var (
	defaultReadHeaderTimeout = time.Second * 5
	defaultReadTimeout       = time.Second * 5
	defaultWriteTimeout      = time.Second * 10
	defaultIdleTimeout       = time.Second * 60
)
