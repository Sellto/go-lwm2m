// Package coap provides a CoAP client and server.
package coap

import (
	"net"
	"time"
)



// Transmit a message.
func Transmit(l *net.UDPConn, a *net.UDPAddr, m Message) error {
	d, err := m.MarshalBinary()
	if err != nil {
		return err
	}

	if a == nil {
		_, err = l.Write(d)
	} else {
		_, err = l.WriteTo(d, a)
	}
	return err
}

// Receive a message.
func Receive(l *net.UDPConn, buf []byte, ResponseTimeout time.Duration) (Message, *net.UDPAddr, error) {
	l.SetReadDeadline(time.Now().Add(ResponseTimeout))
	nr, addr, err := l.ReadFromUDP(buf)
	if err != nil {
		return Message{}, addr, err
	}
	return ParseMessage(buf[:nr],addr)
}
