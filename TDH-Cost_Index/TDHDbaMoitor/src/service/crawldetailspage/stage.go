package crawldetilspage

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"strings"
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/util"
)

func CrawStagePage (query entity.Query,
					stagetsurl string,
					token string,
					producer sarama.SyncProducer,
					TopicInformation string,
					Separator	string) entity.Query{
	stagequeryurl := stagetsurl  + query.ServerKey + "&id="
	var tasksmap []map[int]entity.StageTaskInfo
	var taskarr  []entity.StageTaskInfo

		for i := 0; i < len(query.Stages); i++  {
		var stage entity.JsonStage
		err := json.Unmarshal(
			[]byte(
				util.CrawlPage(
					stagequeryurl	+	util.Int64ToString(
						query.Stages[i]), token)), &stage)
		if err != nil {
			panic(err)
		}
		tasksmap =  append(
			tasksmap,entity.GeTaskList(
				stage))
		}

		for i := 0; i < len(tasksmap); i++ {
			for j := 0; j < len(tasksmap[i]); j++ {
			taskarr = append(taskarr,tasksmap[i][j])

			var task entity.Task
			task.ServerKey=query.ServerKey
			task.TaskID=tasksmap[i][j].TaskID
			task.SqlID=query.SqlID
			task.StageID=tasksmap[i][j].StageID
			if len(tasksmap[i][j].Host) == 0|| tasksmap[i][j].Host  == "" {
				task.TaskHost=query.ServerKey[strings.Index(query.ServerKey, "::")	+	2:
				strings.Index(query.ServerKey, "::")	+	2	+
					strings.Index(query.ServerKey[strings.Index(query.ServerKey, "::")+2:],
				"::")]
			}else{
				task.TaskHost=tasksmap[i][j].Host
			}
			task.TaskStatus=tasksmap[i][j].Status
			task.TaskMessage=tasksmap[i][j].Message
			task.TaskSubmissionTime=tasksmap[i][j].SubmissionTime
			task.TaskCompletionTime=tasksmap[i][j].CompletionTime
			task.User=query.User
			util.ProduceSendMsg(util.CleanNewlineChart(
						entity.TaskToStringBySeparator(task,Separator)),
						producer,
						TopicInformation)
			}
	}
	query.CrawlMessage="CrawlSuccess"
	query.TaskInfo=taskarr
	return query
}
