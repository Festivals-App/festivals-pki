<p align="center">
   <a href="https://github.com/festivals-app/festivals-pki/commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/festivals-app/festivals-identity-server?style=flat"></a>
   <a href="https://github.com/festivals-app/festivals-pki/issues" title="Open Issues"><img src="https://img.shields.io/github/issues/festivals-app/festivals-identity-server?style=flat"></a>
   <a href="./LICENSE" title="License"><img src="https://img.shields.io/github/license/festivals-app/festivals-pki.svg"></a>
</p>

<h1 align="center">
  <br/><br/>
    FestivalsApp PKI
  <br/><br/>
</h1>

The festivals pki repository contains descriptions, workflows and go modules to ensure secure communication between all components of the FestivalsApp.

![Figure 1: Architecture Overview Highlighted](https://github.com/Festivals-App/festivals-documentation/blob/main/images/architecture/export/architecture_overview_pki.svg "Figure 1: Architecture Overview Highlighted")

<hr/>
<p align="center">
  <a href="#development">Development</a> •
  <a href="#deployment">Deployment</a> •
  <a href="#engage">Engage</a>
</p>
<hr/>

To secure communication between components the `FestivalsApp` uses [mTLS](https://www.cloudflare.com/learning/access-management/what-is-mutual-tls/) with self signed certificates.
First we need to create a certificate authority (CA) to issue certificates, then we create a certificate for each service and client.
You can read more about the exact procedures in the [Certification Practice Statements](CERTIFICATIONPRACTICE.md) document.

## Development

I use [easy-rsa](https://github.com/OpenVPN/easy-rsa), which is maintained by the wonderfull community of [OpenVPN](https://openvpn.net/community/), to build and manage the FestivalsApp Root CA.
Even tho the details are quite complex, at the basis of a certificate authority stands a single root certificate. The root certificate is self-signed, meaning that we create it ourself.
The idea is that every party that is communicating with each other needs a certificate signed with this root certificate.

1. First we need to install `easy-rsa` and create the FestivalsApp Root CA.

   ```bash
   # installing the easy-rsa on macOS
   brew install easy-rsa
   
   # init the pki
   easyrsa init-pki
   # Create our root CA certificate (use at least a 40 character random password for the key file)
   easyrsa build-ca
   ```

   On macOS this will create all neccessary files at `/opt/homebrew/etc/pki`

2. To create a certificate/key pair for inter-service communication we first create a certificate request with the name of the service node and then sign the request.

   ```bash
   # create signing request
   easyrsa gen-req <UNIQUE_SERVER_NAME> nopass
   # Enter <UNIQUE_SERVER_DOMAIN_NAME>
   Common Name (eg: your user, host, or server name): <UNIQUE_SERVER_DOMAIN_NAME>
   # sign the service request
   easyrsa --subject-alt-name="DNS:<UNIQUE_SERVER_DOMAIN_NAME>" sign-req serverClient <UNIQUE_SERVER_NAME>
   # or sign a client request
   easyrsa sign-req client <UNIQUE_SERVER_NAME>
   ```

   2.1 Optionally convert certificates and keys to PEM format (for example for usage with mysql)

   ```bash
   openssl x509 -in cert.crt -out cert.pem -outform PEM
   openssl rsa -in cert.key -text > cert-key.pem
   ```

   2.2 Optionally convert certificates and keys to DER format and .p12 keystore file (for usage with swift)

   ```bash
   # Convert from .crt to .pem to .der
   openssl x509 -in cert.crt -out cert.pem -outform PEM
   openssl x509 -in cert.pem -out cert.der -outform der
   # Using -legacy for compability with macOS/iOS. Use at least a 20 character random password for the keystore file.
   openssl pkcs12 -export -legacy -in cert.crt -inkey cert.key -out cert.p12
   ```

   2.3 Optionally convert certificates and keys from CRT to PEM

   ```bash
   # Convert from .crt to .pem public key
   openssl x509 -pubkey -noout -in server.crt > pubkey.pem
   openssl rsa -in server.key -text > privkey.pem
   ```

3. Copy the certificate/key pair to server and move them to their designated location

   ```bash
   scp <path/to/cert/key> <user>@<server>:/home/<user>
   sudo mv </old/cert/location> <new/cert/key/location>
   ```

4. Make the files accessible to the processes and set proper access permissions for certificates and keys

   ```bash
   sudo chown <service-user> </cert/key/location>
   sudo chmod 640/600 <cert/key/location>
   ```

## Local Development

If you want to test on your local machine

sudo nano /etc/hosts

```bash
# local development on this machine
127.0.0.1       gateway.festivalsapp.dev
127.0.0.1       identity-0.festivalsapp.dev
127.0.0.1       festivals-0.festivalsapp.dev
127.0.0.1       database-0.festivalsapp.dev
127.0.0.1       fileserver-0.festivalsapp.dev
127.0.0.1       website-0.festivalsapp.dev

127.0.0.1       festivalsapp.dev
127.0.0.1       www.festivalsapp.dev
127.0.0.1       website.festivalsapp.dev
127.0.0.1       discovery.festivalsapp.dev
127.0.0.1       api.festivalsapp.dev
127.0.0.1       files.festivalsapp.dev
```

## Development on a test server

If you have an development server in your private network

```bash
# local development server for festivalsapp
<ip address>       gateway.festivalsapp.home
<ip address>       identity-0.festivalsapp.home
<ip address>       festivals-0.festivalsapp.home
<ip address>       festivals-1.festivalsapp.home
<ip address>       database-0.festivalsapp.home
<ip address>       fileserver-0.festivalsapp.home
<ip address>       website-0.festivalsapp.home

<gateway ip address>            festivalsapp.home
<gateway ip address>            www.festivalsapp.home
<gateway ip address>            website.festivalsapp.home
<gateway ip address>            discovery.festivalsapp.home
<gateway ip address>            api.festivalsapp.home
<gateway ip address>            files.festivalsapp.home
```

# Deployment

Add the FestivalsaApp Root CA certificate to the trusted root certificates:

## macOS

```bash
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ~/new-root-certificate.crt
```

## Linux (Ubuntu)

```bash
sudo cp new-root-certificate.crt /usr/local/share/ca-certificates/new-root-certificate.crt
sudo update-ca-certificates
```

## Engage

I welcome every contribution, whether it is a pull request or a fixed typo. The best place to discuss questions and suggestions regarding the festivals-pki is the [issues](https://github.com/festivals-app/festivals-pki/issues/) section. More general information and a good starting point if you want to get involved is the [festival-documentation](https://github.com/Festivals-App/festivals-documentation) repository.

The following channels are available for discussions, feedback, and support requests:

| Type                     | Channel                                                |
| ------------------------ | ------------------------------------------------------ |
| **General Discussion**   | <a href="https://github.com/festivals-app/festivals-documentation/issues/new/choose" title="General Discussion"><img src="https://img.shields.io/github/issues/festivals-app/festivals-documentation/question.svg?style=flat-square"></a> </a>   |
| **Other Requests**    | <a href="mailto:simon@festivalsapp.org" title="Email me"><img src="https://img.shields.io/badge/email-Simon-green?logo=mail.ru&style=flat-square&logoColor=white"></a>   |

### Licensing

Copyright (c) 2023-2025 Simon Gaus. Licensed under the [**GNU Lesser General Public License v3.0**](./LICENSE)

<https://docs.bigchaindb.com/projects/server/en/v1.1.0/production-deployment-template/easy-rsa.html#how-to-install-configure-easy-rsa>
openssl rsa -in /opt/homebrew/etc/pki/private/gateway-server.key -out /opt/homebrew/etc/pki/private/gateway-server-unencrypted.key
