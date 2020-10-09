package netHttp

import (
	"API1008Project/src/DataStruct"
	"API1008Project/src/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//手机归属地查询
func GetPhoneNumberAddress(){
	Phone := "18379139021"
	resp,err := http.Get("http://apis.juhe.cn/mobile/get?"+"phone="+Phone+"&dtype=&key=a9ba64bd6a71a873e6d0688f93a0ba04")
	errors.CheckError(err)

	bytes, err := ioutil.ReadAll(resp.Body)
	errors.CheckError(err)


	var result  DataStruct.PhoneAddress
	err = json.Unmarshal(bytes,&result)
	errors.CheckError(err)

	fmt.Println(result.Result)

}

//星座 匹配查询
func GetConstellation() {

	var men string
	var women string
	men,women = "狮子","白羊"
	resp,err := http.Get("http://apis.juhe.cn/xzpd/query?" +
		"men="+men+""+"&women="+women+
		"&key=53dcbada00f8259b9673bfffb08610a5")
	errors.CheckError(err)

	bytes, err := ioutil.ReadAll(resp.Body)
	errors.CheckError(err)

	var result  DataStruct.Constellation
	err = json.Unmarshal(bytes,&result)
	errors.CheckError(err)

	fmt.Println(result.Result)

}
// ip地址查询：
func GetIPAddress() {
	//  62148ce9ff587ed8ed9bc4addf1ccd71
//   http://apis.juhe.cn/ip/ipNew?ip=10.15.1.99&key=62148ce9ff587ed8ed9bc4addf1ccd71

	var Ip string
	Ip = "192.168.1.20"
	resp,err := http.Get("http://apis.juhe.cn/ip/ipNew?ip="+Ip+"&key=62148ce9ff587ed8ed9bc4addf1ccd71")
	errors.CheckError(err)

	bytes, err := ioutil.ReadAll(resp.Body)
	errors.CheckError(err)

	var result  DataStruct.IpAddress
	err = json.Unmarshal(bytes,&result)
	errors.CheckError(err)

	fmt.Println(result.Result)
}
//新闻查询
func GetNewsinfo()  {
	NewsChannel := []string{"top","shehui","guoji","yule","tiyu","junshi","keji","caijing","shishang"}
	resp,err := http.Get("http://v.juhe.cn/toutiao/index?type="+NewsChannel[0]+"&key=355a66d069c8e0e7b4017ba91094e05c")
	errors.CheckError(err)
	bytes, err := ioutil.ReadAll(resp.Body)
	errors.CheckError(err)

	var result  DataStruct.NewsInfo
	err = json.Unmarshal(bytes,&result)
	errors.CheckError(err)

	fmt.Println(result.Result.NewsDataArray)
}
//历史上的今天
func GetTodayOnHistory()  {
	var Months string
	var Days string
	Months = "10"
	Days = "9"
	resp,err := http.Get("http://api.juheapi.com/japi/toh?v=1.0&month="+Months+
		"&day="+Days+"&key=008f8600fe5dfd5481606aa2793a53f2")
	errors.CheckError(err)

	bytes, err := ioutil.ReadAll(resp.Body)
	errors.CheckError(err)

	var result  DataStruct.TodayOnHistory
	err = json.Unmarshal(bytes,&result)
	errors.CheckError(err)

	fmt.Println(result.Result)
}
func GetQQnumber(){
	var QQnumber string
	QQnumber = "1843607154"
	// http://japi.juhe.cn/qqevaluate/qq?qq=2967378768&key=706fcfeac774f4a048a967413792aba7
	resp,err := http.Get("http://japi.juhe.cn/qqevaluate/qq?qq="+QQnumber+
		"&key=706fcfeac774f4a048a967413792aba7")
	errors.CheckError(err)

	bytes, err := ioutil.ReadAll(resp.Body)
	errors.CheckError(err)

	var result  DataStruct.QQnumber
	err = json.Unmarshal(bytes,&result)
	errors.CheckError(err)

	fmt.Println(result.Result)
}
//国内城市天气查询
func GetWeather(){
	City := "南昌"
	resp,err := http.Get("http://apis.juhe.cn/simpleWeather/query?city="+City+"&key=a838e832b80e8fb10009960bd6237a52")
	errors.CheckError(err)

	bytes,err := ioutil.ReadAll(resp.Body)
	errors.CheckError(err)

	var result DataStruct.ResponseWeather
	err = json.Unmarshal(bytes,&result)
	errors.CheckError(err)

	fmt.Println(result.Result)
}