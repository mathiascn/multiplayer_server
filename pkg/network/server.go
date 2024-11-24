package network

import (
	"fmt"
	"net"
)

type UDPServer struct {
	conn    *net.UDPConn
	handler func(*net.UDPConn, *net.UDPAddr, []byte)
}

func NewUDPServer(addr *net.UDPAddr, handler func(*net.UDPConn, *net.UDPAddr, []byte)) (*UDPServer, error) {
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPServer{conn: conn, handler: handler}, nil
}

func (s *UDPServer) Run() {
	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := s.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		go s.handler(s.conn, clientAddr, buffer[:n])
	}
}

func (s *UDPServer) Close() {
	s.conn.Close()
}
