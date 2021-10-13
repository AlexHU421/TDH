package querymap

import (
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/service/crawldetilspage"
	"time"
)

func GetQueryMap(cquerymap chan map[string]entity.Query,
				cservermap chan map[int]entity.JsonServer,
				queriesurl string,token string,
				batchsize string,
				looptime time.Duration) {
	for {
		//fmt.Println(time.Now())
		for _, serverinfo := range <-cservermap {
			go crawldetilspage.CrawQueryPage(cquerymap,queriesurl,token,serverinfo.DataKey,batchsize)
		}
		time.Sleep(looptime)
	}}
