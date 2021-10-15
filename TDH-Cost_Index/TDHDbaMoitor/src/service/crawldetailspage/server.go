package crawldetilspage

import (
	"encoding/json"
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/util"
	"time"
)

func CrawServerPage (c chan map[int]entity.JsonServer,serverurl string,token string){
	for {
		dataType, _ := json.Marshal(util.JsonUnmarshalByString(util.CrawlPage(serverurl, token))["data"])
		c <- entity.MapByJson(string(dataType))
		time.Sleep(time.Minute)
	}
}

func CrawServerPageFindStages (query entity.Query,querysurl string ,token string) entity.Query{
	querystages := querysurl  + query.ServerKey + "&id=" + util.Int64ToString(query.SqlID)
	var relquery entity.JsonQueryStageInfo
	err :=json.Unmarshal([]byte(util.CrawlPage(querystages, token)), &relquery)
	if err != nil {
		panic(err)
	}
	stagearr := make([]int64,len(relquery.Data.Stages))
	for i :=0 ;i< len(stagearr) ;i++  {
		stagearr[i]=relquery.Data.Stages[i].StageID
	}
	query.Stages=stagearr
	return query
}
