package DataStruct

//星座匹配
type Constellation struct {
	Reason string `json:"reason"`
	Result Result_constellation `json:"result"`
	Error_code int `json:"error_code"`
}

type Result_constellation struct {
	Men string `json:"men"`
	Women string `json:"women"`
	Zhishu string `json:"zhishu"`
	Bizhong string `json:"bizhong"`
	Xiangyue string `json:"xiangyue"`
	Tcdj string `json:"tcdj"`
	Jieguo string `json:"jieguo"`
	Lianai string `json:"lianai"`
}


//手机号码归属地
type PhoneAddress struct {
	Reason string `json:"reason"`
	Result Result_phoneaddress `json:"result"`
	Error_code int `json:"error_code"`
}
type Result_phoneaddress struct {
	Province string `json:"province"`
	City string `json:"city"`
	Areacode string `json:"areacode"`
	Zip string `json:"zip"`
	Company string `json:"company"`
	Card string `json:"card"`
}


//IP查询
type IpAddress struct {
	Reason string `json:"reason"`
	Result Result_Ipaddress `json:"result"`
	Error_code int `json:"error_code"`
}
type Result_Ipaddress struct {
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	Isp string `json:"isp"`
}


//新闻查询
type NewsInfo struct{
	Reason string `json:"reason"`
	Result Result_NewsInfo `json:"result"`
	Error_code int `json:"error_code"`
}
type Result_NewsInfo struct {
	Stat string `json:"stat"`
	NewsDataArray []NewsData `json:"data"`
}
type NewsData struct {
	Uniquekey string `json:"uniquekey"`
	Uitle string `json:"uitle"`
	Date string `json:"date"`
	Category string `json:"category"`
	Author_name string `json:"author_name"`
	Url string `json:"url"`
	Thumbnail_pic_s string `json:"thumbnail_pic_s"`
	Thumbnail_pic_s02 string `json:"thumbnail_pic_s_02"`
	Thumbnail_pic_s03 string `json:"thumbnail_pic_s_03"`

}

//历史上的今天
type TodayOnHistory struct{
	Reason string `json:"reason"`
	Result []Result_History `json:"result"`
	Error_code int `json:"error_code"`
}
type  Result_History struct {
	Id string `json:"_id"`
	Title string `json:"title"`
	Pic string `json:"pic"`
	Year int `json:"year"`
	Month int `json:"month"`
	Day int `json:"day"`
	Des string `json:"des"`
	Lunar string `json:"lunar"`
}

//QQ号测吉凶
type QQnumber struct {
	Reason string `json:"reason"`
	Result Result_QQ `json:"result"`
	Error_code int `json:"error_code"`
}
type Result_QQ struct {
	Data QQData `json:"data"`
}
type  QQData struct {
	Conclusion string `json:"conclusion"`
	Analysis string `json:"analysis"`
}

//天气查询
//实时天气
type RealTime struct {
	Info        string `json:"info"`
	Wid         string `json:"wid"`
	Direct      string `json:"direct"`
	Power       string `json:"power"`
	Aqi         string `json:"aqi"`
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
}

/**
 * 某个城市的未来天气的结构体
 */
type Future struct {
	Date        string `json:"date"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
	Wid         Wid    `json:"wid"`
	Direct      string `json:"direct"`
}

/**
 * 早晚温差
 */
type Wid struct {
	Day   string `json:"day"`
	Night string `json:"night"`
}

/**
 * 查询城市的天气的结果
 */
type Result struct {
	City     string   `json:"city"`     //南昌
	RealTime RealTime `json:"realtime"` //实时天气
	Future   []Future `json:"future"`
}

//返回json整体对象结构体
//entity:实体
type ResponseWeather struct {
	Reason     string `json:"reason"`
	Result     Result `json:"result"`
	Error_code int    `json:"error_code"`
}