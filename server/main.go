package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"

	quic "github.com/lucas-clemente/quic-go"
)

func main() {
	listenPort := "0.0.0.0:10443"
	protos := []string{
		"echo",
	}
	listener, err := quic.ListenAddr(listenPort,
		generateTLSConfig(protos), nil)
	if err != nil {
		return
	}
	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			return
		}
		proto := sess.ConnectionState().NegotiatedProtocol
		log.Println("Proto: " + proto)
		stream, err := sess.AcceptStream(context.Background())
		if err != nil {
			return
		}
		defer stream.Close()
		stream.Write([]byte("ボーッとGoを書いてんじゃねーよ！"))
	}
}

func generateTLSConfig(protos []string) *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader,
		&template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   protos,
	}
}
