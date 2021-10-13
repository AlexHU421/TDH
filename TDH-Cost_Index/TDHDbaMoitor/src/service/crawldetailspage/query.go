package crawldetilspage

import (
	"encoding/json"
	"strconv"
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/util"
	"time"
)

func CrawQueryPage (cquerymap chan map[string]entity.Query,
					queriesurl string,
					token string,
					serverKey string,
					batchSize string){
	for {
		serverqueryurl := queriesurl + serverKey + "&dataSize=" + batchSize
		var queies entity.JsonQuery1
		//fmt.Println(util.CrawlPage(serverqueryurl, token))
		err := json.Unmarshal([]byte(util.CrawlPage(serverqueryurl, token)), &queies)
		if err != nil {
			panic(err)
		}
		cquerymap <- entity.GetQueriesList(queies,serverKey)
		time.Sleep(time.Second*5)
	}
}



func ReCrawQueryPage (query entity.Query,querysurl string ,token string) entity.Query{
	querystages := querysurl  + query.ServerKey + "&id=" +strconv.FormatInt(query.SqlID,10)
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
