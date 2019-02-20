package lwm2m

import (
  "strings"
    "../coap"
)

func uriPathParser(msg []byte) map[string]map[string][]Ressource {
  data := make(map[string]map[string][]Ressource)
  items := strings.Split(string(msg),",")
  var splitpointer = 0
  for _,item := range items {
      if string(item[1]) == "/" {
        splitpointer = 2
      } else {
        splitpointer = 1
      }
      object := strings.Split(item[splitpointer:len(item)-1],"/")
      objID := object[0]
      sub := object[1]
      if data[objID] == nil {
        data[objID]= make(map[string][]Ressource)
      }
      if len(object) > 2 {
        res := object[2]
        data[objID][sub] = append(data[objID][sub],Ressource{ID:res,PollingDelay:30})
      } else {
          data[objID][sub] = []Ressource{}
      }
  }
  return data
}


func checkResponse(m coap.Message)bool{

return true
}
