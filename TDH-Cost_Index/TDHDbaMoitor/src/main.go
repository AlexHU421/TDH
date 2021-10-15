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
	querymapsGuard sync.Mutex
	cquerymapsGuard sync.RWMutex
	delmapsGuard sync.RWMutex

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
	cquerymapsGuard.Lock()
	go querymap.GetQueryMap(cquerymap,cservermap,queriesurl,token,"1000",time.Second*5)
	cquerymapsGuard.Unlock()

	//解析querymap 爬取stage关键页获取task信息
	go func() {
		for {
			cquerymapsGuard.RLock()
			for k,v :=range <- cquerymap {
				_,ok := querymaps[k]
				if !ok {
					querymapsGuard.Lock()
					querymaps[k]=v
					querymapsGuard.Unlock()
				}
			}
			cquerymapsGuard.RUnlock()
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
				 if len(v.Stages)>0 && len(v.TaskInfo)==0 {
				 	if (v.CrawlMessage == "")   {
						 querymaps[k] = crawldetilspage.CrawStagePage(v,stagetsurl,token,producer,TopicInformation,Separator)
					}
				}else if len(v.Stages)==0 && (v.CrawlMessage == "") {
					 aa := querymaps[k]
					 aa.CrawlMessage = "CrawlWARN_NoStagesInfo"
					 querymaps[k] = aa
				 }


				if !(v.CrawlMessage == "")  && util.FilterByUnixtime(v.SubmissionTime,8,"minute"){
					delmaps[k]=util.UnixMillTime(time.Now().UnixNano())
				}
			}

			fmt.Println("lenquerymap",len(querymaps))
			delmapsGuard.RLock()
			fmt.Println("lendelmap",len(delmaps))
			delmapsGuard.RUnlock()
			querymapsGuard.Unlock()
			time.Sleep(time.Second * 5)
		}
	}()







	//清理符合规定的sql清单
	go querymap.CleanQueryMap(querymaps,delmaps,querymapsGuard,delmapsGuard,producer,TopicInformation,Separator,time.Second * 5)



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
