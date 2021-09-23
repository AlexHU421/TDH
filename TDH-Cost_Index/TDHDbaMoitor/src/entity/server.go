package entity

import (
	"encoding/json"
	"tdhdbamonithr/src/util"
)

type JsonServer struct {
	ServerId int `json:"serverId"`
	DataKey string `json:"key"`
	Hosts string `json:"host"`
	Port  int `json:"port"`
	Timestamp int64 `json:"timestamp"`
	FirstSeen int64 `json:"firstSeen"`
	LocalTimestamp int64 `json:"localTimestamp"`
}


func MapByJson (dataString string) map[int]JsonServer{

	bMap :=util.JsonUnmarshalByString(dataString)
	servermapallinfo := make(map[string]JsonServer)
	servermap := make(map[int]JsonServer)
	activeservermap := make(map[string]int64)
	for _,v := range bMap {
		dataType1 , _ :=json.Marshal(v)
		ds := string(dataType1)
		var server JsonServer
		err := json.Unmarshal([]byte(ds), &server)
		if err != nil{
			panic(err)
		}
		activetimestamp,nilkey := activeservermap[server.Hosts]
		if nilkey {
			//fmt.Println("nilkey",server.ServerId,server.Hosts,server.Timestamp,"|+|",activeservermap)
			if activetimestamp<server.Timestamp {
				activetimestamp = server.Timestamp
				activeservermap[server.Hosts]=server.Timestamp
				servermapallinfo[server.Hosts]=server
			}
		}else {
			//fmt.Println("existkey",server.ServerId,server.Hosts,server.Timestamp,"|+|",activeservermap)
			if activetimestamp<server.Timestamp {
				activetimestamp = server.Timestamp
				activeservermap[server.Hosts]=server.Timestamp
				servermapallinfo[server.Hosts]=server
			}
		}
	}
	for _,v := range servermapallinfo{
		//servermap =entity.MapByJson(v)
		servermap[v.ServerId]=v
		//delete(servermapallinfo,k)
	}
	return servermap
}
