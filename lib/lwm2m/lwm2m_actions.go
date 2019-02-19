package lwm2m

import (
  "fmt"
  "log"
  "net"
  "strconv"
  "../coap"
)


func Listen(t taskArg){
  msg,addr,_ := coap.Receive(t.Conn,buf,ResponseTimeout)
  if msg.Option(11) != nil {
    if msg.Option(11) == "rd" {
      task.Add(Task{toDo:Register,Conn: t.Conn,From: addr,Msg:msg})
    } else if  _, isPresent := db[msg.Option(11).(string)]; isPresent{
      task.Add(Task{toDo:Update,Conn: t.Conn,From: addr, Msg:msg})
    }
  }
}

func Register(t taskArg) {
  log.Println("New client try to register from",t.From)
  newclientid := "dev"+strconv.Itoa(len(db))

  msg := coap.Message {
    Type:      coap.Acknowledgement,
    Code:      coap.Created,
    MessageID: t.Msg.MessageID,
    Token:     t.Msg.Token,
  }
  msg.SetOption(coap.ContentFormat, coap.TextPlain)
  msg.SetOption(coap.LocationPath, newclientid)
  err := coap.Transmit(t.Conn, t.From, msg)
  if err != nil {
    log.Printf("Error")
  } else {
    log.Println("Endpoint",newclientid,"added into the database")
  }
  db[newclientid] = Client{}
  task.Add(Task{toDo:Read,Conn: t.Conn,From: t.From, Msg:msg})
}

func Update(t taskArg) {
  log.Println("Update request from",t.Msg.Option(11))
  msg := coap.Message {
    Type:      coap.Acknowledgement,
    Code:      coap.Changed,
    MessageID: t.Msg.MessageID,
    Token:     t.Msg.Token,
  }
  msg.SetOption(coap.ContentFormat, coap.TextPlain)
  err := coap.Transmit(t.Conn, t.From, msg)
  if err != nil {
    log.Printf("Error")
  } else {
    log.Println("Endpoint",t.Msg.Option(11),"updated")
  }
  task.Add(Task{toDo:Read,Conn: t.Conn,From: t.From, Msg:msg})
}

func getObjectTree(t taskArg) {
  log.Println("Test")
  msg := coap.Message {
    Type:      coap.Confirmable,
    Code:      coap.GET,
    MessageID: 1,
  }
  msg.SetOption(coap.Accept,coap.AppLinkFormat)
  msg.SetPath([]string{"3","0"})
  err := coap.Transmit(t.Conn, t.From, msg)
  if err != nil {
    log.Printf("Error")
  } else {
    log.Println("Transmit done")
  }
}



func Read(t taskArg) {
  log.Println("Test")
  msg := coap.Message {
    Type:      coap.Confirmable,
    Code:      coap.GET,
    MessageID: 11,
  }
  msg.SetOption(coap.Accept,coap.AppLinkFormat)
  msg.SetPath([]string{"3","0"})
  err := coap.Transmit(t.Conn, t.From, msg)
  if err != nil {
    log.Printf("Error")
  } else {
    log.Println("Transmit done")
  }
}

func Write(l *net.UDPConn) {
  log.Println("Write Task")
}
