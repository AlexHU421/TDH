package crawldetilspage

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/util"
)

func CrawStagePage (query entity.Query,
					stagetsurl string,
					token string,
					taskGuard sync.RWMutex,
					taskmaps map[string]entity.Task) entity.Query{
	stagequeryurl := stagetsurl  + query.ServerKey + "&id="
	var tasksmap []map[int]entity.StageTaskInfo
	var taskarr  []entity.StageTaskInfo

		for i := 0; i < len(query.Stages); i++  {
		var stage entity.JsonStage
		err := json.Unmarshal(
			[]byte(
				util.CrawlPage(
					stagequeryurl+strconv.FormatInt(
						query.Stages[i],10), token)), &stage)
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
			taskGuard.Lock()
			var task entity.Task
			task.ServerKey=query.ServerKey
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
			taskmaps[task.ServerKey+"||"+
				strconv.FormatInt(task.SqlID,10)+"||"+
				strconv.FormatInt(task.StageID,10)+"||"+
				task.TaskHost+"||"+
				strconv.FormatInt(task.TaskCompletionTime,10)]=task
			taskGuard.Unlock()
			}
	}
	query.TaskInfo=taskarr
	return query
}
