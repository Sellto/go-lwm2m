package lwm2m

import (
  //"log"
  //"fmt"
  "time"
  "net"
  //"../coap"
)

const(
  ResponseTimeout = time.Millisecond*2000
  maxPktLen = 1500
)

var task = NewQueue()
var buf = make([]byte, maxPktLen)
var	db = make(map[string]Client)

func Start(l *net.UDPConn) {
  for true {
    if next := task.Next(); next != nil {
      next.run()
    } else {
      task.Add(Task{toDo:Listen, Conn: l})
    }
  }
}
