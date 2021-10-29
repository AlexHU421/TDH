package querymap

import (
	"github.com/Shopify/sarama"
	"sync"
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/util"
	"time"
)

func CleanQueryMap (querymaps map[string]entity.Query,
					delmaps map[string]int64,
					wftasklist map[string]entity.Wftaskinfo,
					querymapsGuard sync.Mutex,
					delmapsGuard sync.RWMutex,
					producer sarama.SyncProducer,
					TopicInformation string,
					Separator string,
					looptime time.Duration) {
	for {
		for k, _ := range delmaps {
			querymapsGuard.Lock()
			delmapsGuard.Lock()
			util.ProduceSendMsg(
				util.CleanNewlineChart(
					entity.QueryToStringBySeparator(querymaps[k],Separator,wftasklist)),producer,TopicInformation,)
			delete(querymaps, k)
			delete(delmaps,k)
			querymapsGuard.Unlock()
			delmapsGuard.Unlock()
		}
		time.Sleep(looptime)
	}
}
