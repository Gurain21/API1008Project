package http

import (
	"API1008Project/src/DataStruct"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPhoneNumberAddress(){
//手机归属地查询
	Phone := "18379139021"
	resp,err := http.Get("http://apis.juhe.cn/mobile/get?"+"phone="+Phone+"&dtype=&key=a9ba64bd6a71a873e6d0688f93a0ba04")
	if err != nil {
		fmt.Println("get请求错误，请重试！错误代码为：",err.Error())
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body读取错误，请重试！错误代码为：",err.Error())
	}
	var result  DataStruct.PhoneAddress
	err = json.Unmarshal(bytes,&result)
	if err != nil {
		fmt.Println("bytes解析错误，请重试！错误代码为：",err.Error())
	}
	fmt.Println(result.Result)

}
func GetConstellation()  {
	//星座 匹配查询
	var men string
	var women string
	men,women = "狮子","白羊"
	resp,err := http.Get("http://apis.juhe.cn/xzpd/query?" +
		"men="+men+""+"&women="+women+
		"&key=53dcbada00f8259b9673bfffb08610a5")
	if err != nil {
		fmt.Println("get请求错误，请重试！错误代码为：",err.Error())
	}
	 bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Body读取错误，请重试！错误代码为：",err.Error())
	}
	var result  DataStruct.Constellation
	err = json.Unmarshal(bytes,&result)
	if err != nil {
		fmt.Println("bytes解析错误，请重试！错误代码为：",err.Error())
	}
	fmt.Println(result.Result)
}
