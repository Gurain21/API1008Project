package main

import (
	"API1008Project/src/netHttp"
	"API1008Project/src/requHandle"
	"fmt"
	"net/http"
)

func main() {
	netHttp.GetPhoneNumberAddress()
	//netHttp.GetWeather()
	//netHttp.GetConstellation()
	//netHttp.GetIPAddress()
	//netHttp.GetNewsinfo()
	//netHttp.GetTodayOnHistory()
	//netHttp.GetQQnumber()
	//netHttp.GetWeather()
	http.HandleFunc("/index", requHandle.Index)
	//http.HandleFunc("/error", requHandle.Error)
	http.Handle("/web/css/", http.StripPrefix("/web/css", http.FileServer(http.Dir("./web/css/"))))
	http.Handle("/web/img/", http.StripPrefix("/web/img", http.FileServer(http.Dir("./web/img/"))))
	http.Handle("/web/js/", http.StripPrefix("/web/js", http.FileServer(http.Dir("./web/js/"))))

	fmt.Println("正在监听端口: 127.0.0.1:6699 ")
	http.ListenAndServe("127.0.0.1:6699", nil)

}
