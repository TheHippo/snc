package snc

import (
	"crypto/tls"
	"fmt"
	"net"
	"sync"
)

type Server struct {
	hasConnection     bool
	hasConnectionLock sync.RWMutex
	port              int
	ip                net.IP
	listener          net.Listener
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:            net.ParseIP(ip),
		port:          port,
		hasConnection: false,
	}
}

func (s *Server) Listen() error {
	certPem, keyPEM, err := generateCert(s.ip.String())
	if err != nil {
		return err
	}

	cert, err := tls.X509KeyPair(certPem.Bytes(), keyPEM.Bytes())
	if err != nil {
		return err
	}
	listener, err := tls.Listen("tcp", fmt.Sprintf("%s:%d", s.ip.String(), s.port), &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		// MinVersion:         tls.VersionSSL30,
	})
	if err != nil {
		return err
	}
	s.listener = listener
	for {
		con, err := s.listener.Accept()
		if err != nil {
			return err
		}
		if s.checkConnection(con) == true {
			go s.handleConnection(con)
		}
	}
}

func (s *Server) checkConnection(con net.Conn) bool {
	s.hasConnectionLock.RLock()
	defer s.hasConnectionLock.RUnlock()
	if s.hasConnection == true {
		fmt.Println("Already got active connections")
		con.Close()
		return false
	}
	return true
}

func (s *Server) handleConnection(con net.Conn) {
	s.hasConnectionLock.Lock()
	s.hasConnection = true
	s.hasConnectionLock.Unlock()
	defer func() {
		s.hasConnectionLock.Lock()
		s.hasConnection = false
		s.hasConnectionLock.Unlock()
	}()
	fmt.Println("Handle incoming connection", con.RemoteAddr().String())

}
