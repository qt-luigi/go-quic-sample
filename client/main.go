package main

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"os"

	quic "github.com/lucas-clemente/quic-go"
)

func main() {
	dAddr := "localhost:10443"
	session, err := quic.DialAddr(dAddr,
		&tls.Config{InsecureSkipVerify: true,
			NextProtos: []string{"echo"},
		}, nil)
	if err != nil {
		log.Println(err)
		return
	}
	stream, err := session.
		OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	stream.Write([]byte{0})
	io.Copy(os.Stdout, stream)
}
