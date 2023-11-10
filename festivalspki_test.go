package festivalspki_test

import (
	"crypto/tls"
	"testing"

	festivalspki "github.com/Festivals-App/festivals-pki"
)

func TestLoadServerCertificateHandler(t *testing.T) {

	handler := festivalspki.LoadServerCertificateHandler("certificates/*.festivalsapp.dev.crt", "certificates/*.festivalsapp.dev.key", "certificates/festivalsapp-development-root-ca.crt")
	_, err := handler(&tls.ClientHelloInfo{})
	if err != nil {
		t.Errorf("Handler failed to load server certificates.")
	}
}

func TestLoadServerCertificates(t *testing.T) {
	_, err := festivalspki.LoadServerCertificates("certificates/*.festivalsapp.dev.crt", "certificates/*.festivalsapp.dev.key", "certificates/festivalsapp-development-root-ca.crt")
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
