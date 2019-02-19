package lwm2m

import (
  "strings"
)

func registerParser(msg []byte) Client {
  data := make(map[string]map[string][]Ressource)
  items := strings.Split(string(msg),",")
  for id,item := range items {
    if id > 0 {
      object := strings.Split(item[2:len(item)-1],"/")
      objID := object[0]
      sub := object[1]
      res := object[2]
      if data[objID] == nil {
        data[objID]= make(map[string][]Ressource)
      }
      data[objID][sub] = append(data[objID][sub],Ressource{ID:res,PollingDelay:30})
    }
  }
  return Client{Data:data}
}
