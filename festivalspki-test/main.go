package main

import (
	"fmt"

	festivalspki "github.com/Festivals-App/festivals-pki"
)

func main() {

	fmt.Println("TESTED festivals-pki package...")

	return

	serverCertificateChain, err := festivalspki.LoadServerCertificates("cert.crt", "cert.key", "rootCA.crt")
	if err != nil {
		fmt.Println("Failed to load server certificates.")
	}
	fmt.Println("Did load server certificate: " + serverCertificateChain.Leaf.IssuingCertificateURL[0])

	cert, err := festivalspki.LoadX509Certificate("cert.crt")
	if err != nil {
		fmt.Println("Failed to load single certificate.")
	}
	fmt.Println("Did load single certificate: " + cert.DNSNames[0])

}

/*
package main

import "github.com/Festivals-App/festivals-pki"
 go run main.go

func main() {

	cert, err := LoadServerCertificates("cert.crt", "cert.key")


	cert, err = LoadX509Certificate("cert.crt")
}

*/
