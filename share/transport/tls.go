package transport

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
)

// Not used temporarily
func newCustomTLSKeyPair(certfile, keyfile string) (*tls.Certificate, error) {
	tlsCert, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		return nil, err
	}
	return &tlsCert, nil
}

func newRandomTLSKeyPair() *tls.Certificate {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		&template,
		&key.PublicKey,
		key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tlsCert
}

// Not used temporarily
// Only support one ca file to add
func newCertPool(caPath string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()

	caCrt, err := os.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	pool.AppendCertsFromPEM(caCrt)

	return pool, nil
}

func NewServerTLSConfig() (*tls.Config, error) {
	var base = &tls.Config{}

	// Not used temporarily

	// if certPath == "" || keyPath == "" {
	// 	// server will generate tls conf by itself
	// 	cert := newRandomTLSKeyPair()
	// 	base.Certificates = []tls.Certificate{*cert}
	// } else {
	// 	cert, err := newCustomTLSKeyPair(certPath, keyPath)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	base.Certificates = []tls.Certificate{*cert}
	// }

	// if caPath != "" {
	// 	pool, err := newCertPool(caPath)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	base.ClientAuth = tls.RequireAndVerifyClientCert
	// 	base.ClientCAs = pool
	// }

	cert := newRandomTLSKeyPair()
	base.Certificates = []tls.Certificate{*cert}

	return base, nil
}

func NewClientTLSConfig(serverName string) (*tls.Config, error) {
	var base = &tls.Config{}

	// Not used temporarily

	// if caPath != "" {
	// 	pool, err := newCertPool(caPath)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	base.RootCAs = pool
	// 	base.InsecureSkipVerify = false
	// } else {
	// 	base.InsecureSkipVerify = true
	// }

	base.InsecureSkipVerify = true
	base.ServerName = serverName

	return base, nil
}
