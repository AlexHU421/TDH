package entity

import (
	"tdhdbamonithr/src/util"
	"time"
)

type Task struct {
	ServerKey 			string
	TaskID				int64
	SqlID 				int64
	StageID              int64
	TaskHost 				string
	TaskStatus				string
	TaskMessage 			string
	TaskSubmissionTime		int64
	TaskCompletionTime		int64
	User				string
}

//taskmaps[task.ServerKey+"||"+
//	util.Int64ToString(task.SqlID)	+	"||"	+
//	task.TaskHost+"||"+
//	util.Int64ToString(task.TaskID)]=task



func TaskToStringBySeparator (task Task,separator string)string {
	return "task"+	separator	+
		task.ServerKey+"||"	+util.Int64ToString(task.SqlID)+"||"+task.TaskHost+"||"+util.Int64ToString(task.TaskID)+"||"+
		util.Int64ToString(task.TaskSubmissionTime)	+	separator	+task.ServerKey	+	separator	+
		util.Int64ToString(task.TaskID)	+	separator	+
		util.Int64ToString(task.SqlID)	+	separator	+
		util.Int64ToString(task.StageID)	+	separator	+
		task.TaskHost	+	separator	+
		task.TaskStatus +	separator	+
		util.CleanNewlineChart(task.TaskMessage) +	separator	+
		time.Unix(task.TaskSubmissionTime/1000,0).Format("2006-01-02") +	separator	+
		util.Int64ToString(task.TaskSubmissionTime) +	separator	+
		util.Int64ToString(task.TaskCompletionTime) +	separator	+
		task.User


}

