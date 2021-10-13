package entity




type Task struct {
	ServerKey 			string
	SqlID 				int64
	StageID              int64
	TaskHost 				string
	TaskStatus				string
	TaskMessage 			string
	TaskSubmissionTime		int64
	TaskCompletionTime		int64
	User				string
}




