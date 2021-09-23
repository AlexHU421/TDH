package util

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func CrawlPage(url string,token string)string{
	tr := &http.Transport{ TLSClientConfig: &tls.Config{InsecureSkipVerify:true}}
	client :=&http.Client{Transport:tr}
	request,err := http.NewRequest("GET",url,nil)
	request.Header.Add("Guardian-Access-Token",token)

	//获取prometheus   每分钟cpu占比信息
	//resp, err := http.Get(url)
	resp,_ := client.Do(request)
	//判空报错
	if err != nil {
		// handle error
		panic(err)
	}

	//关闭http session
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}
