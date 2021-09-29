package main

import (
	"encoding/json"
	"fmt"
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/util"
	"time"
)

const(
	dbaurl=	"https://tdh3:4040/api/inceptor/"
	dataKey=	"?dataKey="
	serverurl = "https://tdh3:4040/api/inceptor/servers"
	queriesurl = dbaurl + "sqls" + dataKey
	querysurl= dbaurl + "sql" + dataKey
	stagetsurl = dbaurl + "stage" + dataKey
	token = "XxTBdghnvfdPu2zr1NRe-TDH"
)
var (
	querymaps = make(map[string]entity.Query)
	delmaps =  make(map[string]string)
)

type MyJsonName struct {
	Data []interface{}  `json:"data"`
	Error []interface{} `json:"error"`
	Info  []interface{} `json:"info"`
	Query struct {
		DataKey  interface{} `json:"dataKey"`
		DataSize int64       `json:"dataSize"`
		ID       int64       `json:"id"`
		StringID interface{} `json:"stringId"`
	} `json:"query"`
	Warning           []interface{} `json:"warning"`
	WatchmanTimestamp int64         `json:"watchmanTimestamp"`
}

func crawserverpage (c chan map[int]entity.JsonServer){
	for {
	dataType, _ := json.Marshal(util.JsonUnmarshalByString(util.CrawlPage(serverurl, token))["data"])
	c <- entity.MapByJson(string(dataType))
		time.Sleep(time.Minute)
	}
}
func crawquerypage (cquerymap chan map[string]entity.Query,serverKey string){
	for {
		serverqueryurl := queriesurl + serverKey + "&dataSize=100"
		var queies entity.JsonQuery1
		err := json.Unmarshal([]byte(util.CrawlPage(serverqueryurl, token)), &queies)
		if err != nil {
			panic(err)
		}
		cquerymap <- entity.GetQueriesList(queies)
		time.Sleep(time.Second*5)
	}
}


func main() {

	now := util.UnixMillTime(time.Now().UnixNano())
	fmt.Println(now)
	//Get lastserver map info    后续切片并发需存入redis做并发访问池
	cservermap :=make(chan map[int]entity.JsonServer)
	//爬server列表（只保留最近一次启动）
	go 	crawserverpage(cservermap)
	//Test  query数据量太大，需定时器定时获取
	cquerymap :=make(chan map[string]entity.Query)
	//爬query列表（增量map）
	go func() {
		for {
			fmt.Println(time.Now())
			for _, serverinfo := range <-cservermap {
			go crawquerypage(cquerymap,serverinfo.DataKey)
			}
			time.Sleep(time.Second*5)
		}}()
	go func() {
		for {
			for k,v :=range <- cquerymap {
				querymaps[k]=v
			}
			for k,v :=range querymaps {
				if 	util.FilterByUnixtime(v.SubmissionTime,1,"minute") {
					delmaps[k]=k
				}
			}

			fmt.Println(querymaps)
			fmt.Println(len(querymaps))
			time.Sleep(time.Second * 5)
		}
	}()


	go func() {
		for {
			for _, v := range delmaps {
				delete(querymaps, v)
			}
			time.Sleep(time.Second * 5)
		}
	}()

	//fmt.Println(querymap)












	for {
		fmt.Println("delmap:",len(delmaps),delmaps,len(delmaps))
		time.Sleep(time.Second*5)
	}



	//
	//
	//
	////获取stage相关详情，得到executor信息
	//
	////stagequeryurl := stagetsurl + "Inceptor::tdh2::7fd73b47-8aec-4668-951e-8c8a72255832" + "&id=10119"
	//stagequeryurl := stagetsurl + "Inceptor::tdh2::7fd73b47-8aec-4668-951e-8c8a72255832" + "&id=12447"
	////fmt.Println(util.CrawlPage(stagequeryurl, token))
	////fmt.Println(stagequeryurl)
	//
	//
	////Get querymap info    后续切片并发需存入redis做并发访问池
	//taskmap := make(map[int]entity.Task)
	//var stage entity.JsonStage
	//err = json.Unmarshal([]byte(util.CrawlPage(stagequeryurl, token)), &stage)
	//if err != nil {
	//	panic(err)
	//}
	//taskmap = entity.GeTaskList(stage)
	//fmt.Println(taskmap)

}
