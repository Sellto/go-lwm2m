package lwm2m

import(
  "net"
  "../coap"
)

type Task struct {
	toDo  func(taskArg)
  From *net.UDPAddr
  Conn *net.UDPConn
  Msg  coap.Message
}
func (t Task)run() {
  t.toDo(taskArg{From:t.From,Conn:t.Conn,Msg:t.Msg})
}

type taskArg struct {
  From *net.UDPAddr
  Conn *net.UDPConn
  Msg  coap.Message
}
