package lwm2m

import (
  "log"
)
type Client struct {
  ID string
  Data map[string]map[string][]Ressource
}

type Ressource struct {
  ID string
  Value string
  PollingDelay int
}

func (c Client)getAvailalableObjects()[]string{
  var s = []string{}
  for key,_ := range (c.Data){
    s = append(s,key)
  }
  return s
}

func (c Client)getInstancesOf(ObjID string)[]string{
  var s = []string{}
  for key,_ := range (c.Data[ObjID]){
    s = append(s,key)
  }
  return s
}

func (c Client)getRessourceOf(ObjID string,InID string)[]Ressource{
  return c.Data[ObjID][InID]
  }

func (c Client)setPollingDelayOf(o string,i string,r string ,d int){
  for ID,value := range c.Data[o][i] {
    if value.ID == r {
      c.Data[o][i][ID].PollingDelay = d
      log.Println("Polling delay of ressource <"+o+"/"+i+"/"+r+"> is set to",d)
    }
  }
}
