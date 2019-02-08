package lwm2m

import (
  "log"
  "net"
  "../coap"
)


func Listen(t taskArg){
  msg,addr,_ := coap.Receive(t.Conn,buf,ResponseTimeout)
  if msg.Option(11) == "rd" {
    if len(msg.Options(11)) == 1 {
      task.Add(Task{toDo:Register,Conn: t.Conn,From: addr,Msg:msg})
    } else {
      task.Add(Task{toDo:Update,Conn: t.Conn,From: addr, Msg:msg})
    }
  }
}

func Register(t taskArg) {
  log.Println("Register Task")
  log.Println(string(t.Msg.Payload))
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

func Update(t taskArg) {
  log.Println("Update")
  msg := coap.Message {
    Type:      coap.Acknowledgement,
    Code:      coap.Changed,
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

func Read(l *net.UDPConn) {
  log.Println("Read Task")
}

func Write(l *net.UDPConn) {
  log.Println("Write Task")
}
