package snc

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"sync"
)

// Client holds all informations about an snc client
type Client struct {
	host string
	port int
	conn *tls.Conn
}

// NewClient creates a new client from hostname and target port
func NewClient(host string, port int) *Client {
	return &Client{
		host: host,
		port: port,
	}
}

// Dial attemps to open an TCP connection to the server
func (c *Client) Dial() error {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port), &tls.Config{
		InsecureSkipVerify: true,
		// MinVersion:         tls.VersionSSL30,
	})
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}()

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println(err)
		}
		os.Exit(0)
	}()
	wg.Wait()
	return nil
}
