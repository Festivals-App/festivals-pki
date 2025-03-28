package festivalspki_test

import (
	"crypto/tls"
	"testing"

	festivalspki "github.com/Festivals-App/festivals-pki"
)

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
