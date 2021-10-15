package entity

import (
	"strconv"
	"sync"
	"tdhdbamonithr/src/util"
)

type JsonStage struct {
	Data struct {
		AttemptID int64 `json:"attemptId"`
		Attempts  []struct {
			AttemptID      int64  `json:"attemptId"`
			CompletionTime int64  `json:"completionTime"`
			Message        string `json:"message"`
			StageID        int64  `json:"stageId"`
			SubmissionTime int64  `json:"submissionTime"`
		} `json:"attempts"`
		CompletionTime int64  `json:"completionTime"`
		Description    string `json:"description"`
		Executors      []struct {
			ActiveTasks        int64       `json:"activeTasks"`
			CompletedTasks     int64       `json:"completedTasks"`
			DirStatus          interface{} `json:"dirStatus"`
			ExecutorID         string      `json:"executorId"`
			FailedTasks        int64       `json:"failedTasks"`
			FirstSeen          int64       `json:"firstSeen"`
			FirstSeenLocal     int64       `json:"firstSeenLocal"`
			GcInfo             interface{} `json:"gcInfo"`
			Host               string      `json:"host"`
			Input              int64       `json:"input"`
			Jobs               interface{} `json:"jobs"`
			LastSeen           int64       `json:"lastSeen"`
			LastSeenLocal      int64       `json:"lastSeenLocal"`
			LocalTasks         interface{} `json:"localTasks"`
			LogLevel           interface{} `json:"logLevel"`
			MaxMemory          int64       `json:"maxMemory"`
			Port               int64       `json:"port"`
			ShuffleReadBytes   int64       `json:"shuffleReadBytes"`
			ShuffleSpillDisk   int64       `json:"shuffleSpillDisk"`
			ShuffleSpillMemory int64       `json:"shuffleSpillMemory"`
			ShuffleWriteBytes  int64       `json:"shuffleWriteBytes"`
			Sqls               interface{} `json:"sqls"`
			Stages             interface{} `json:"stages"`
			Statuses           interface{} `json:"statuses"`
			Tasks              interface{} `json:"tasks"`
			TasksActive        int64       `json:"tasksActive"`
			TasksCompleted     int64       `json:"tasksCompleted"`
			TasksFailed        int64       `json:"tasksFailed"`
			TotalCores         int64       `json:"totalCores"`
			TotalTask          int64       `json:"totalTask"`
			TotalTaskTime      int64       `json:"totalTaskTime"`
			Version            int64       `json:"version"`
		} `json:"executors"`
		Input             int64       `json:"input"`
		InputRowsString   interface{} `json:"inputRowsString"`
		JobID             int64       `json:"jobId"`
		Message           string      `json:"message"`
		OutputRowsString  interface{} `json:"outputRowsString"`
		PartitionSize     int64       `json:"partitionSize"`
		ShuffleReadBytes  int64       `json:"shuffleReadBytes"`
		ShuffleWriteBytes int64       `json:"shuffleWriteBytes"`
		Sources           interface{} `json:"sources"`
		SqlGroup          string      `json:"sqlGroup"`
		SqlID             int64       `json:"sqlId"`
		StageID           int64       `json:"stageId"`
		StageMode         interface{} `json:"stageMode"`
		StagePoolName     interface{} `json:"stagePoolName"`
		Status            string      `json:"status"`
		SubmissionTime    int64       `json:"submissionTime"`
		TaskTimes         interface{} `json:"taskTimes"`
		Tasks             []struct {
			Accumulators                 struct{}    `json:"accumulators"`
			Aggregation                  interface{} `json:"aggregation"`
			Attempt                      int64       `json:"attempt"`
			CompletionTime               int64       `json:"completionTime"`
			Durations                    struct{}    `json:"durations"`
			ExecutorID                   string      `json:"executorId"`
			GarbageCollectionTime        int64       `json:"garbageCollectionTime"`
			Host                         string      `json:"host"`
			Input                        int64       `json:"input"`
			Message                      string      `json:"message"`
			Metrics                      struct{}    `json:"metrics"`
			Port                         int64       `json:"port"`
			ResultSerializationTime      int64       `json:"resultSerializationTime"`
			SchedulerDelay               int64       `json:"schedulerDelay"`
			Sequence                     interface{} `json:"sequence"`
			ShuffleReadBytes             int64       `json:"shuffleReadBytes"`
			ShuffleSpillDisk             int64       `json:"shuffleSpillDisk"`
			ShuffleSpillMemory           int64       `json:"shuffleSpillMemory"`
			ShuffleWriteBytes            int64       `json:"shuffleWriteBytes"`
			StageID                      int64       `json:"stageId"`
			SubmissionTime               int64       `json:"submissionTime"`
			TaskID                       int64       `json:"taskId"`
			TaskIndexInStage             int64       `json:"taskIndexInStage"`
			TaskLocality                 string      `json:"taskLocality"`
			TimeSpentFetchingTaskResults int64       `json:"timeSpentFetchingTaskResults"`
			WriteTime                    int64       `json:"writeTime"`
		} `json:"tasks"`
		TasksActive    int64       `json:"tasksActive"`
		TasksCompleted int64       `json:"tasksCompleted"`
		TasksExpected  int64       `json:"tasksExpected"`
		TasksFailed    int64       `json:"tasksFailed"`
		TasksSkipped   int64       `json:"tasksSkipped"`
		TotalTaskTime  int64       `json:"totalTaskTime"`
		UserName       interface{} `json:"userName"`
	} `json:"data"`
	Error []interface{} `json:"error"`
	Info  []interface{} `json:"info"`
	Query struct {
		DataKey  string      `json:"dataKey"`
		DataSize int64       `json:"dataSize"`
		Host     interface{} `json:"host"`
		ID       int64       `json:"id"`
		Port     interface{} `json:"port"`
		StringID interface{} `json:"stringId"`
	} `json:"query"`
	Warning           []interface{} `json:"warning"`
	WatchmanTimestamp int64         `json:"watchmanTimestamp"`
}



type StageTaskInfo struct {
	TaskID              int64
	StageID 			int64
	SqlID				int64
	Status				string
	Message 			string
	Host 				string
	SubmissionTime		int64
	CompletionTime		int64
}

func TaskInfoListSplitToString (ServerKey string,tl []StageTaskInfo) string{
	if tl == nil {
		return "[]"
	}
	str := "["
	for i := 0; i < len(tl); i++ {
		str +=ServerKey	+	"||"	+
			util.Int64ToString(tl[i].SqlID)		+	"||"	+
			tl[i].Host +	"||"	+
			util.Int64ToString(tl[i].TaskID)	+	"||"	+
			util.Int64ToString(tl[i].SubmissionTime)	+	","
	}

	return  str[:len(str)-1] +"]"
}

func StagesListToString (il []int64) string {
	if il == nil {
		return "[]"
	}
	b := "["
	if len(il) == 0 {
		return b+"]"
	} else {
		for i := 0; i < len(il); i++ {
			b += strconv.FormatInt(il[i], 10) + ","
		}
		return b[:len(b)-1] +"]"
	}
}

func GeTaskList (jsstage JsonStage) map[int]StageTaskInfo {
	var n sync.WaitGroup
	taskmap := make(map[int]StageTaskInfo)

	for i := 0; i < len(jsstage.Data.Tasks); i++ {
		n.Add(1)
		go func (i int, n *sync.WaitGroup){
			defer n.Done()
			var task StageTaskInfo
			task.TaskID=jsstage.Data.Tasks[i].TaskID
			task.StageID=jsstage.Data.StageID
			task.SqlID=jsstage.Data.SqlID
			task.Status=jsstage.Data.Status
			task.Message=jsstage.Data.Message
			task.Host=jsstage.Data.Tasks[i].Host
			task.SubmissionTime=jsstage.Data.Tasks[i].SubmissionTime
			task.CompletionTime=jsstage.Data.Tasks[i].CompletionTime
			taskmap[i]=task
		}(i, &n)
		n.Wait()
	}
	return taskmap
}


