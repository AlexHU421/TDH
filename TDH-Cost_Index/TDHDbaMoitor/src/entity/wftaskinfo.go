package entity

import (
	"sort"
	"strings"
	"tdhdbamonithr/src/util"
)

type Oldtaskinfo struct {
	Taskid	string
	Wfid	string
	Name	string
	Configuration []string
}
type Wftaskinfo struct {
	Taskid	string
	Wfid	string
	Name	string
	Configuration string
}

type Similar struct {
	Taskid string
	Similarty float64
}
type Similarlist []Similar
func (s Similarlist) Swap(i,j int){
	s[i],s[j]=s[j],s[i]
}
func (s Similarlist) Len()int{
	return len(s)
}
func (s Similarlist) Less(i,j int)bool{
	return s[i].Similarty < s[j].Similarty
}



func GetWfTaskList(query string,mysqlconn string, args ...interface{}) map[string]Wftaskinfo {
	var oldtaskinfomap = make(map[string]Oldtaskinfo)
	var wftaskinfomap = make(map[string]Wftaskinfo)
	rows,err := util.GetdbConn(mysqlconn).Query(query,args...)
	if err != nil {
		// handle error
		panic(err)
	}
	for rows.Next(){
		var oldtaskinfo Oldtaskinfo
		var taskid,wfid,name,configuration string
		if err:= rows.Scan(&taskid, &wfid,&name,&configuration); err==nil {}
		//fmt.Println(taskid,wfid,name,strings.Replace(configuration,"\\n","\n",-1))
		oldtaskinfo.Taskid=taskid
		oldtaskinfo.Wfid=wfid
		oldtaskinfo.Name=name
		oldtaskinfo.Configuration=strings.Split(configuration,";\\n")
		oldtaskinfomap[taskid]=oldtaskinfo
	}

	for k,v := range(oldtaskinfomap){
		for i := 0; i < len(v.Configuration); i++ {
			if i+1 < len(v.Configuration){
				v.Configuration[i]+=";"
			}
			if !util.FilterBySQL(v.Configuration[i]){
			strings.Replace(v.Configuration[i],"\\n","|^---^|",-1)
				var wftaskinfo Wftaskinfo
				wftaskinfo.Taskid=k
				wftaskinfo.Wfid=v.Wfid
				wftaskinfo.Name=v.Name
				wftaskinfo.Configuration=v.Configuration[i]
				wftaskinfomap[k+"SQL"+util.IntToString(i)]=wftaskinfo
			}
		}
	}
	return wftaskinfomap
}



func GetSimilaryListInFo (str string, wftasklist map[string]Wftaskinfo,separator string) string{

	SimilartyList := make(Similarlist,len(wftasklist))
	i:=0
	for _,v := range wftasklist {

		similar := util.SimilarText(v.Configuration,str)
		if  similar>80 {
			SimilartyList[i] = Similar{v.Taskid,similar}
			i++
		}
	}
	sort.Stable(SimilartyList)

	return 	SimilartyList[len(SimilartyList)-1].Taskid	+	separator	+
		util.StringfFormatFloat(SimilartyList[len(SimilartyList)-1].Similarty)	+	separator	+
		SimilartyList[len(SimilartyList)-2].Taskid	+	separator	+
		util.StringfFormatFloat(SimilartyList[len(SimilartyList)-2].Similarty)	+	separator	+
		SimilartyList[len(SimilartyList)-3].Taskid	+	separator	+
		util.StringfFormatFloat(SimilartyList[len(SimilartyList)-3].Similarty)
}
