package festivalspki

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// LoadServerCertificatesHandler will return a function that loads the server certificate chain based on the given ClientHelloInfo.
func LoadServerCertificateHandler(serverCert string, serverKey string, rootCACert string) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		// ClientHelloInfo is not used, just try to load the local certificates
		certificate, err := tls.LoadX509KeyPair(serverCert, serverKey)
		if err != nil {
			return nil, errors.New("Failed to load server certificate and key with error: " + err.Error())
		}
		rootCACert, err := LoadX509Certificate(rootCACert)
		if err != nil {
			return nil, errors.New("Failed to load FestivalsApp Root CA certificate with error: " + err.Error())

		}
		certificate.Certificate = append(certificate.Certificate, rootCACert.Raw)
		return &certificate, err
	}
}

// LoadServerCertificates will attempt to load the server certificate chain.
func LoadServerCertificates(serverCert string, serverKey string, rootCACert string) (*tls.Certificate, error) {
	certificate, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		return nil, errors.New("Failed to load server certificate and key with error: " + err.Error())
	}
	rootCACertificate, rootErr := LoadX509Certificate(rootCACert)
	if rootErr != nil {
		return nil, errors.New("Failed to load FestivalsApp Root CA certificate with error: " + err.Error())

	}
	certificate.Certificate = append(certificate.Certificate, rootCACertificate.Raw)
	return &certificate, err
}

// LoadX509Certificate reads and parses a certificate from a .crt file.
// The file must contain PEM encoded data. The certificate file may only contain one certificate.
func LoadX509Certificate(certFile string) (*x509.Certificate, error) {

	certContent, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(certContent)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

// Creates and returns a certificate pool with the given certificate added to it.
func LoadCertificatePool(certFile string) (*x509.CertPool, error) {

	rootCertPool := x509.NewCertPool()
	certContent, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	if ok := rootCertPool.AppendCertsFromPEM(certContent); !ok {
		return nil, errors.New("Failed to append certificate to certificate pool.")
	}
	return rootCertPool, nil
}
