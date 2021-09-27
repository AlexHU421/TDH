package entity

import "strconv"

type JsonQuery1 struct {
	Data []struct {
		AccumulableDurations          interface{} `json:"accumulableDurations"`
		ActiveTasks                   int64       `json:"activeTasks"`
		CompletionTime                int64       `json:"completionTime"`
		Description                   string      `json:"description"`
		DetailedDescription           string      `json:"detailedDescription"`
		DurationBlockMap              struct{}    `json:"durationBlockMap"`
		Durations                     interface{} `json:"durations"`
		ExecutionFinishTime           int64       `json:"executionFinishTime"`
		ExecutionStartTime            int64       `json:"executionStartTime"`
		ExtraJobID                    int64       `json:"extraJobId"`
		FetchComputeTotal             int64       `json:"fetchComputeTotal"`
		FetchCount                    int64       `json:"fetchCount"`
		FetchResultStartTime          int64       `json:"fetchResultStartTime"`
		FetchRowCount                 int64       `json:"fetchRowCount"`
		FileCommitFinishTime          int64       `json:"fileCommitFinishTime"`
		Group                         string      `json:"group"`
		Jobs                          string      `json:"jobs"`
		JobsActive                    int64       `json:"jobsActive"`
		JobsCompleted                 int64       `json:"jobsCompleted"`
		JobsFailed                    int64       `json:"jobsFailed"`
		LogicalPlanString             interface{} `json:"logicalPlanString"`
		LogicalPlanStrings            interface{} `json:"logicalPlanStrings"`
		MemoryPerTask                 int64       `json:"memoryPerTask"`
		Message                       string `json:"message"`
		MessageNameStr                interface{} `json:"messageNameStr"`
		Meta                          struct{}    `json:"meta"`
		Metrics                       interface{} `json:"metrics"`
		Mode                          string      `json:"mode"`
		NonLinearAccumulableDurations interface{} `json:"nonLinearAccumulableDurations"`
		NonLinearDurations            interface{} `json:"nonLinearDurations"`
		PhysicalPlanString            interface{} `json:"physicalPlanString"`
		PhysicalPlanStrings           interface{} `json:"physicalPlanStrings"`
		PlanInputs                    interface{} `json:"planInputs"`
		PlanOutputs                   interface{} `json:"planOutputs"`
		RealFetchResultStartTime      int64       `json:"realFetchResultStartTime"`
		SessionID                     int64       `json:"sessionId"`
		SessionName                   string      `json:"sessionName"`
		SqlBlockJSON                  interface{} `json:"sqlBlockJson"`
		SqlID                         int64       `json:"sqlId"`
		SqlName                       interface{} `json:"sqlName"`
		SqlType                       interface{} `json:"sqlType"`
		Stages                        interface{} `json:"stages"`
		StagesActive                  int64       `json:"stagesActive"`
		StagesCompleted               int64       `json:"stagesCompleted"`
		StagesFailed                  int64       `json:"stagesFailed"`
		StagesSkipped                 int64       `json:"stagesSkipped"`
		State                         string      `json:"state"`
		SubmissionTime                int64       `json:"submissionTime"`
		TableScans                    interface{} `json:"tableScans"`
		Tags                          string      `json:"tags"`
		TasksActive                   int64       `json:"tasksActive"`
		TasksCompleted                int64       `json:"tasksCompleted"`
		TasksExpected                 int64       `json:"tasksExpected"`
		TasksFailed                   int64       `json:"tasksFailed"`
		TasksSkipped                  int64       `json:"tasksSkipped"`
		TotalTaskTime                 int64       `json:"totalTaskTime"`
		User                          string      `json:"user"`
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


type Query struct {
	SqlID 				int64
	User				string
	Description			string
	ExecutionFinishTime int64
	ExecutionStartTime  int64
	CompletionTime      int64
	Message				string
}

func GetQueriesList (queies JsonQuery1) map[string]Query{

	querymap  := make(map[string]Query)
		//fmt.Printf("%T",queies.Data)s
	for _,v :=range queies.Data{
		var query Query
		query.SqlID=v.SqlID
		query.CompletionTime=v.CompletionTime
		query.Description=v.Description
		query.ExecutionFinishTime=v.ExecutionFinishTime
		query.ExecutionStartTime=v.ExecutionStartTime
		query.User=v.User
		query.Message=v.Message
		querymap [queies.Query.DataKey+"||"+strconv.FormatInt(v.SqlID,10)]=query
	}
	return querymap
}
