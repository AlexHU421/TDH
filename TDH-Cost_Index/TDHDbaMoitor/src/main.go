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
	serverurl = "https://tdh2:4040/api/inceptor/servers"
	queriesurl = dbaurl + "sqls" + dataKey
	querysurl= dbaurl + "sql" + dataKey
	stagetsurl = dbaurl + "stage" + dataKey
	token = "XxTBdghnvfdPu2zr1NRe-TDH"
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



func main() {
	now := time.Now().UnixNano()

	//tiker := time.NewTicker(time.Second*5)
	//for i:=1; i>0 ;i++{
	//	i=1

	fmt.Println(util.UnixMillTime(now))
	//var xm MyJsonName
	//err := json.Unmarshal([]byte(util.CrawlPage(url,token)), &xm)
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(xm.Data)
	//fmt.Println(util.CrawlPage(url,token))

	dataType, _ := json.Marshal(util.JsonUnmarshalByString(util.CrawlPage(serverurl, token))["data"])
	//Get lastserver map info    后续切片并发需存入redis做并发访问池
	servermap := entity.MapByJson(string(dataType))

	//Test look
	for serverid, serverinfo := range servermap {
		fmt.Println(serverid, serverinfo.DataKey)
	}

	//Test  query数据量太大，需定时器定时获取
	serverqueryurl := queriesurl + "Inceptor::tdh2::7fd73b47-8aec-4668-951e-8c8a72255832" + "&dataSize=10"
	//fmt.Println(serverqueryurl)
	//fmt.Println(util.CrawlPage(serverqueryurl,token))

	//Get querymap info    后续切片并发需存入redis做并发访问池
	querymap := make(map[string]entity.Query)
	var queies entity.JsonQuery1
	err := json.Unmarshal([]byte(util.CrawlPage(serverqueryurl, token)), &queies)
	if err != nil {
		panic(err)
	}
	querymap = entity.GetQueriesList(queies)
	fmt.Println(querymap)








	//获取stage相关详情，得到executor信息

	//stagequeryurl := stagetsurl + "Inceptor::tdh2::7fd73b47-8aec-4668-951e-8c8a72255832" + "&id=10119"
	stagequeryurl := stagetsurl + "Inceptor::tdh2::7fd73b47-8aec-4668-951e-8c8a72255832" + "&id=12447"
	//fmt.Println(util.CrawlPage(stagequeryurl, token))
	//fmt.Println(stagequeryurl)


	//Get querymap info    后续切片并发需存入redis做并发访问池
	taskmap := make(map[int]entity.Task)
	var stage entity.JsonStage
	err = json.Unmarshal([]byte(util.CrawlPage(stagequeryurl, token)), &stage)
	if err != nil {
		panic(err)
	}
	taskmap = entity.GeTaskList(stage)
	fmt.Println(taskmap)

}
