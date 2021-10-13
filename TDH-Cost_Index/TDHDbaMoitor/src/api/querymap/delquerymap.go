package querymap

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
	"tdhdbamonithr/src/entity"
	"time"
)

func CleanQueryMap (querymaps map[string]entity.Query,
					delmaps map[string]int64,
					querymapsGuard sync.RWMutex,
					delmapsGuard sync.RWMutex,
					producer sarama.SyncProducer,
					TopicInformation string,
					looptime time.Duration) {
	for {
		for k, _ := range delmaps {
			querymapsGuard.Lock()
			delmapsGuard.Lock()
			//fmt.Println("要发的")
			//fmt.Println(querymaps[k])
			query,_:=json.Marshal(querymaps[k])
			//fmt.Println(
			//	strings.Replace(
			//		strings.Replace(
			//			string(query),
			//			"\n","|+---+|",-1),
			//		"\r","|+---+|",-1))
			//fmt.Println("发完了")
			producer.SendMessage(&sarama.ProducerMessage{Topic:TopicInformation,Key:nil,Value: sarama.StringEncoder(
				strings.Replace(
				strings.Replace(
					string(query),
					"\n","|+---+|",-1),
				"\r","|+---+|",-1))})
			delete(querymaps, k)
			delete(delmaps,k)
			querymapsGuard.Unlock()
			delmapsGuard.Unlock()
		}
		time.Sleep(looptime)
	}
}
