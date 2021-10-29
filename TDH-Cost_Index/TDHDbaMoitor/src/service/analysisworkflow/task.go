package analysisworkflow

import (
	"sort"
	"tdhdbamonithr/src/entity"
	"tdhdbamonithr/src/util"
)

func GetSimilaryListInFo (str string, wftasklist map[string]entity.Wftaskinfo) string{

	SimilartyList := make(entity.Similarlist,len(wftasklist))
	i:=0
	for _,v := range wftasklist {

		similar := util.SimilarText(v.Configuration,str)
		if  similar>80 {
			SimilartyList[i] = entity.Similar{v.Taskid,similar}
			i++
		}
	}
	sort.Stable(SimilartyList)

	s:="|+|"
	s+=SimilartyList[len(SimilartyList)-1].Taskid+"|+|"+util.StringfFormatFloat(SimilartyList[len(SimilartyList)-1].Similarty)+"|+|"
	s+=SimilartyList[len(SimilartyList)-2].Taskid+"|+|"+util.StringfFormatFloat(SimilartyList[len(SimilartyList)-2].Similarty)+"|+|"
	s+=SimilartyList[len(SimilartyList)-3].Taskid+"|+|"+util.StringfFormatFloat(SimilartyList[len(SimilartyList)-3].Similarty)+"|+|"
	return s
}
