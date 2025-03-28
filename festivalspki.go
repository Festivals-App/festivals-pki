package festivalspki

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// LoadServerCertificatesHandler will return a function that loads the server certificate chain based on the given ClientHelloInfo.
//
// Deprecated: Don't include the root CA in the server certificate use LoadServerCertificate instead.
func LoadServerCertificateChainHandler(serverCert string, serverKey string, rootCACert string) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		// ClientHelloInfo is not used, just try to load the local certificates
		certificate, err := tls.LoadX509KeyPair(serverCert, serverKey)
		if err != nil {
			return nil, errors.New("failed to load server certificate and key with error: " + err.Error())
		}
		rootCACert, err := LoadX509Certificate(rootCACert)
		if err != nil {
			return nil, errors.New("failed to load server root certificate with error: " + err.Error())

		}
		certificate.Certificate = append(certificate.Certificate, rootCACert.Raw)
		return &certificate, nil
	}
}

// LoadServerCertificates will attempt to load the server certificate chain.
//
// Deprecated: Don't include the root CA in the server certificate use LoadServerCertificate instead.
func LoadServerCertificateChain(serverCert string, serverKey string, rootCACert string) (*tls.Certificate, error) {
	certificate, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		return nil, errors.New("failed to load server certificate and key with error: " + err.Error())
	}
	rootCACertificate, rootErr := LoadX509Certificate(rootCACert)
	if rootErr != nil {
		return nil, errors.New("failed to load server root certificate with error: " + rootErr.Error())

	}
	certificate.Certificate = append(certificate.Certificate, rootCACertificate.Raw)
	return &certificate, nil
}

// LoadServerCertificates will attempt to load the given server certificate and key.
func LoadServerCertificate(serverCert string, serverKey string) (*tls.Certificate, error) {
	certificate, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		return nil, errors.New("failed to load server certificate and key with error: " + err.Error())
	}
	return &certificate, nil
}

// LoadX509Certificate reads and parses a certificate from a .crt file.
// The file must contain PEM encoded data. The certificate file may only contain one certificate.
func LoadX509Certificate(certFile string) (*x509.Certificate, error) {

	certContent, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(certContent)
	return x509.ParseCertificate(block.Bytes)
}

// Creates and returns a certificate pool with the given certificate added to it.
func LoadCertificatePool(certFile string) (*x509.CertPool, error) {

	certContent, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	rootCertPool := x509.NewCertPool()
	if ok := rootCertPool.AppendCertsFromPEM(certContent); !ok {
		return nil, errors.New("failed to append certificate to certificate pool")
	}
	return rootCertPool, nil
}
