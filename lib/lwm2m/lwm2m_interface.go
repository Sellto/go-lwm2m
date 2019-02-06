package lwm2m

import (
  "log"
  "net"
  "../coap"
)

func Register(t taskArg) {
  log.Println("Register Task")
  msg := coap.Message {
    Type:      coap.Acknowledgement,
    Code:      coap.Created,
    MessageID: t.Msg.MessageID,
    Token:     t.Msg.Token,
  }

  msg.SetOption(coap.ContentFormat, coap.TextPlain)
  msg.SetOption(coap.LocationPath, "rd/test01")
  err := coap.Transmit(t.Conn, t.From, msg)
  if err != nil {
    log.Printf("Error")
  }
}

func Update(l *net.UDPConn) {
  log.Println("Update Task")
}

func Read(l *net.UDPConn) {
  log.Println("Read Task")
}

func Write(l *net.UDPConn) {
  log.Println("Write Task")
}
