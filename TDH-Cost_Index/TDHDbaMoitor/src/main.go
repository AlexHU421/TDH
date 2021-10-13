package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"tdhdbamonithr/src/api/querymap"
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/service/crawldetilspage"
	"tdhdbamonithr/src/util"
	"time"
)

const(

	//KafkaInfomation="****.kafka.******.***:****"
	//TopicInformation="tdh-dbaInfo"
	//dbaurl=	"https://***:4040/api/inceptor/"


	KafkaInfomation="***.***.***.***:****"
	TopicInformation="testgo1"
	Separator string ="|+|"
	dbaurl=	"https://***:4040/api/inceptor/"
	token = "="****************"-TDH"


	dataKey=	"?dataKey="
	serverurl = "https://tdh3:4040/api/inceptor/servers"
	queriesurl = dbaurl + "sqls" + dataKey
	querysurl= dbaurl + "sql" + dataKey
	stagetsurl = dbaurl + "stage" + dataKey

)

var (
	querymaps = make(map[string]entity.Query)
	taskmaps = make(map[string]entity.Task)
	delmaps =  make(map[string]int64)
	querymapsGuard sync.RWMutex
	delmapsGuard sync.RWMutex
	taskGuard sync.RWMutex
)





func main() {

	now := util.UnixMillTime(time.Now().UnixNano())
	fmt.Println(now)

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	producer,err := sarama.NewSyncProducer([]string{KafkaInfomation},config)
	if err != nil {
		panic(err)
	}


	//Get lastserver map info    后续切片并发需存入redis做并发访问池
	cservermap :=make(chan map[int]entity.JsonServer)
	//爬server列表（只保留最近一次启动）
	go 	crawldetilspage.CrawServerPage(cservermap,serverurl,token)

	cquerymap :=make(chan map[string]entity.Query)
	//爬query列表（增量map）
	go querymap.GetQueryMap(cquerymap,cservermap,queriesurl,token,"100",time.Second*5)


	//解析querymap 爬取stage关键页获取task信息
	go func() {
		for {
			for k,v :=range <- cquerymap {
				_,ok := querymaps[k]
				if !ok {
					querymapsGuard.Lock()
					querymaps[k]=v
					querymapsGuard.Unlock()
				}
			}

			for k,v :=range querymaps {
				//fmt.Println("berore:",k,querymaps[k])
				querymapsGuard.Lock()
				querymaps[k] =crawldetilspage.CrawServerPageFindStages(v,querysurl,token)
				querymapsGuard.Unlock()
				//fmt.Println("after:",k,querymaps[k])
			}

			time.Sleep(time.Second * 5)
		}
	}()




	// 判断querymap如何清理 确定清理逻辑 推送key至delmap中
	go func(){
		for {
			querymapsGuard.Lock()

			for k, v := range querymaps {
				 if !(v.Stages == nil) || (v.CrawlMessage != "") {
					if (v.TaskInfo == nil) && (v.CrawlMessage == ""){
					querymaps[k] = crawldetilspage.CrawStagePage(v,stagetsurl,token,taskGuard,taskmaps)
					}else {
						if 	util.FilterByUnixtime(v.SubmissionTime,3,"minute") {
							aa := querymaps[k]
							aa.CrawlMessage="CrawlSuccess"
							delmaps[k]=util.UnixMillTime(time.Now().UnixNano())
						}
					}
				}
				if v.Stages == nil || v.TaskInfo == nil {
					 aa := querymaps[k]
						aa.CrawlMessage="CrawlERR_NotFoundStages"
					 querymaps[k]=aa
				}
				if 	util.FilterByUnixtime(v.SubmissionTime,3,"minute") {
					aa := querymaps[k]
					if  v.Stages == nil || v.TaskInfo == nil {
						aa.CrawlMessage = "CrawlWARN_NoDetailInfo"
					}else {
						aa.CrawlMessage = "CrawlWARN_Timeout"
					}
					delmaps[k]=util.UnixMillTime(time.Now().UnixNano())
				}
			}
			querymapsGuard.Unlock()
			fmt.Println("lentaskmap:",len(taskmaps),"lenquerymap",len(querymaps),"lendelmap",len(delmaps))
			//fmt.Println("taskmap:",len(taskmaps),taskmaps)
			fmt.Println("querymap:",len(querymaps),querymaps)
			time.Sleep(time.Second * 5)
		}
	}()



	//清理符合规定的sql清单
	go querymap.CleanQueryMap(querymaps,delmaps,querymapsGuard,delmapsGuard,producer,TopicInformation,time.Second * 5)



	//fmt.Println(querymap)












	for {
		//delmapsGuard.RLock()
		//fmt.Println("delmap:",len(delmaps),delmaps,len(delmaps))
		//delmapsGuard.RUnlock()
		time.Sleep(time.Hour*999)
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
