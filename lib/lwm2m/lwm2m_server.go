package lwm2m

import (
  //"log"
  //"fmt"
  "time"
  "net"
  "github.com/foize/go.fifo"
  "../coap"
)

const(
  ResponseTimeout = time.Millisecond*2000
  maxPktLen = 1500
)

type Task struct {
	toDo  func(taskArg)
  Arg   taskArg
}

type taskArg struct {
  From *net.UDPAddr
  Conn *net.UDPConn
  Msg  coap.Message
}


var task = fifo.NewQueue()
var buf = make([]byte, maxPktLen)

func Listen(t taskArg){
  msg,addr,_ := coap.Receive(t.Conn,buf,ResponseTimeout)
  if msg.Option(11) == "rd" {
    if len(msg.Options(11)) == 1 {
      task.Add(Task{toDo:Register, Arg:taskArg{Conn: t.Conn,From: addr,Msg:msg}})
    } else {
      task.Add(Task{toDo:Update, Arg:taskArg{Conn: t.Conn,From: addr,Msg:msg}})
    }
  }
}

func Start(l *net.UDPConn) {
  for true {
    if next := task.Next(); next != nil {
      next.(Task).toDo(next.(Task).Arg)
    } else {
      task.Add(Task{toDo:Listen, Arg:taskArg{Conn: l}})
    }
  }
}
