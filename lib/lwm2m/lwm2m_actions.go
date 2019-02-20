package lwm2m

import (
  //"fmt"
  "log"
  "net"
  "strconv"
  "../coap"
  "strings"
  "encoding/json"
//"fmt"
"io/ioutil"
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
  db[newclientid] = Client{ID:"dev"+strconv.Itoa(len(db)),Data:uriPathParser(t.Msg.Payload)}
  task.Add(Task{toDo:Discover,Conn: t.Conn,From: t.From, Msg:msg, EndPoint:  db[newclientid]})
  //task.Add(Task{toDo:getAllValues,Conn: t.Conn,From: t.From, Msg:msg, EndPoint:  db[t.Msg.Option(11).(string)]})

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
    task.Add(Task{toDo:Discover,Conn: t.Conn,From: t.From, Msg:msg, EndPoint:  db[t.Msg.Option(11).(string)]})
    //task.Add(Task{toDo:getAllValues,Conn: t.Conn,From: t.From, Msg:msg, EndPoint:  db[t.Msg.Option(11).(string)]})
  }
}

func Discover(t taskArg){
  for _,obj := range t.EndPoint.getAvailalableObjects() {
    for _,inst := range t.EndPoint.getInstancesOf(obj) {
      //Create Message to get the All ressource of Client instance.
      msg := coap.Message {
        Type:      coap.Confirmable,
        Code:      coap.GET,
        MessageID: 101,
      }
      msg.SetOption(coap.Accept,coap.AppLinkFormat)
      msg.SetPath([]string{obj,inst})
      err := coap.Transmit(t.Conn, t.From, msg)
      if err != nil {
        log.Printf("Error")
      } else {
        log.Println("Send message to get tree of ",t.EndPoint.ID)
      }
      //Listen to get response.
      response,addr,_ := coap.Receive(t.Conn,buf,ResponseTimeout)
      //Check if the response is the waited response
      if response.MessageID == 101 && response.Type == coap.Acknowledgement && t.From.String() == addr.String(){
        t.EndPoint.Data = uriPathParser(response.Payload)
        //upload EP in the database
        db[t.EndPoint.ID] = t.EndPoint
        log.Println("Get a valid response from",t.EndPoint.ID)
      } else {
        log.Println("Failed get a message from other device")
      }

      //Create Message to get the All data of Client instance.
      msg = coap.Message {
        Type:      coap.Confirmable,
        Code:      coap.GET,
        MessageID: 102,
      }
      msg.SetOption(coap.ContentFormat, coap.TextPlain)
      msg.SetPath([]string{obj,inst})
      err = coap.Transmit(t.Conn, t.From, msg)
      if err != nil {
        log.Printf("Error")
      } else {
        log.Println("Send message to get all values of",t.EndPoint.ID)
      }
      //Listen to get response.
      response,addr,_ = coap.Receive(t.Conn,buf,ResponseTimeout)
      //Check if the response is the waited response
      if response.MessageID == 102 && response.Type == coap.Acknowledgement && t.From.String() == addr.String(){
        values := strings.Split(string(response.Payload),",")
        for id,_ := range t.EndPoint.getRessourceOf(obj,inst){
          db[t.EndPoint.ID].Data[obj][inst][id].Value = values[id]
        }
        log.Println("Get a valid response from",t.EndPoint.ID)
      } else {
        log.Println("Failed get a message from other device")
      }

      }
    }
    rankingsJson, _ := json.Marshal(db)
    ioutil.WriteFile("output.json", rankingsJson, 0644)
    log.Println("Database saved into the output.json file")
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
