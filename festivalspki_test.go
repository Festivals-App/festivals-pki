package festivalspki_test

import (
	"crypto/tls"
	"fmt"
	"testing"

	festivalspki "github.com/Festivals-App/festivals-pki"
)

// TestNewServerTLSConfig tests the initialization of tls.Config with test certificates.
func TestNewServerTLSConfig(t *testing.T) {
	serverCert := "certificates/festivalsapp.dev.crt"
	serverKey := "certificates/festivalsapp.dev.key"
	clientCA := "certificates/festivalsapp-development-root-ca.crt"
	tlsConfig, err := festivalspki.NewServerTLSConfig(serverCert, serverKey, clientCA)

	if err != nil {
		t.Fatalf("❌ NewServerTLSConfig failed: %v", err)
	}
	if tlsConfig.ClientAuth != tls.RequireAndVerifyClientCert {
		t.Errorf("❌ Expected ClientAuth to be RequireAndVerifyClientCert, got %v", tlsConfig.ClientAuth)
	}
	if len(tlsConfig.Certificates) == 0 {
		t.Errorf("❌ Expected server certificate to be loaded, but found none")
	}
	if tlsConfig.ClientCAs == nil {
		t.Errorf("❌ Expected ClientCAs to be set, but it is nil")
	}
	if err == nil && len(tlsConfig.Certificates) > 0 && tlsConfig.ClientCAs != nil {
		fmt.Println("✅ NewServerTLSConfig successfully loaded certificates and initialized tls.Config")
	}
}

func TestLoadServerCertificateChainHandler(t *testing.T) {

	handler := festivalspki.LoadServerCertificateChainHandler("certificates/festivalsapp.dev.crt", "certificates/festivalsapp.dev.key", "certificates/festivalsapp-development-root-ca.crt")
	_, err := handler(&tls.ClientHelloInfo{})
	if err != nil {
		t.Errorf("Handler failed to load server certificates.")
	}
}

func TestLoadServerCertificateChain(t *testing.T) {
	_, err := festivalspki.LoadServerCertificateChain("certificates/festivalsapp.dev.crt", "certificates/festivalsapp.dev.key", "certificates/festivalsapp-development-root-ca.crt")
	if err != nil {
		t.Errorf("Failed to load server certificates.")
	}
}

func TestLoadServerCertificate(t *testing.T) {
	_, err := festivalspki.LoadServerCertificate("certificates/festivalsapp.dev.crt", "certificates/festivalsapp.dev.key")
	if err != nil {
		t.Errorf("Failed to load server certificates.")
	}
}

func TestLoadX509Certificate(t *testing.T) {
	_, err := festivalspki.LoadX509Certificate("certificates/festivalsapp-development-root-ca.crt")
	if err != nil {
		t.Errorf("Failed to load X509 certificate.")
	}
	//t.Log("running TestLoadX509Certificate")
}
