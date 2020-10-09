package requHandle

import (
	"html/template"
	"net/http"
)

func Index(writer http.ResponseWriter,request *http.Request)  {
	temp, err := template.ParseFiles("web/html/index.html")
	if err != nil {
		//writer.Write([]byte(err.Error()))
		errorTmp, _ := template.ParseFiles("web/html/404.html")
		errorTmp.Execute(writer, err.Error())
		return
	}

	//解析模板正常
	temp.Execute(writer, nil)
}
