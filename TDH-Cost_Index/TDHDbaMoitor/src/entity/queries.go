package entity

import (
	"strconv"
)

type JsonQuery struct {
	Data []struct {
		CompletionTime      int64  `json:"completionTime"`
		Description         string `json:"description"`
		DetailedDescription string `json:"detailedDescription"`
		Durations           struct {
			Data_Move int64 `json:"Data Move"`
			Compile   int64 `json:"compile"`
		} `json:"durations"`
		ExecutionFinishTime int64       `json:"executionFinishTime"`
		ExecutionStartTime  int64       `json:"executionStartTime"`
		ExtraJobID          int64       `json:"extraJobId"`
		Group               string      `json:"group"`
		Jobs                interface{} `json:"jobs"`
		JobsActive          int64       `json:"jobsActive"`
		JobsCompleted       int64       `json:"jobsCompleted"`
		JobsFailed          int64       `json:"jobsFailed"`
		LogicalPlanString   interface{} `json:"logicalPlanString"`
		LogicalPlanStrings  interface{} `json:"logicalPlanStrings"`
		Message             string `json:"message"`
		Metrics             struct{}    `json:"metrics"`
		Mode                interface{} `json:"mode"`
		PhysicalPlanString  interface{} `json:"physicalPlanString"`
		PhysicalPlanStrings interface{} `json:"physicalPlanStrings"`
		PlanInputs          interface{} `json:"planInputs"`
		PlanOutputs         interface{} `json:"planOutputs"`
		SessionID           int64       `json:"sessionId"`
		SqlID               int64       `json:"sqlId"`
		Stages              interface{} `json:"stages"`
		StagesActive        int64       `json:"stagesActive"`
		StagesCompleted     int64       `json:"stagesCompleted"`
		StagesFailed        int64       `json:"stagesFailed"`
		StagesSkipped       int64       `json:"stagesSkipped"`
		State               string      `json:"state"`
		SubmissionTime      int64       `json:"submissionTime"`
		TotalTaskTime       int64       `json:"totalTaskTime"`
		User                string      `json:"user"`
	} `json:"data"`
	Error []interface{} `json:"error"`
	Info  []interface{} `json:"info"`
	Query struct {
		DataKey  string      `json:"dataKey"`
		DataSize int64       `json:"dataSize"`
		ID       int64       `json:"id"`
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

func GetQueriesList (queies JsonQuery) map[string]Query{

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
