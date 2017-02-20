package snc

import (
	"crypto/tls"
	"fmt"
)

type Client struct {
	host string
	port int
	conn *tls.Conn
}

func NewClient(host string, port int) *Client {
	return &Client{
		host: host,
		port: port,
	}
}

func (c *Client) Dial() error {
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port), &tls.Config{
		InsecureSkipVerify: true,
		// MinVersion:         tls.VersionSSL30,
	})
	if err != nil {
		return err
	}
	fmt.Println(conn)

	return nil
}
