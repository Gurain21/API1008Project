package DataStruct

type Constellation struct {
	Reason string `json:"reason"`
	Result Result_constellation `json:"result"`
	Error_code int `json:"error_code"`
}
type PhoneAddress struct {
	Reason string `json:"reason"`
	Result Result_phoneaddress `json:"result"`
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
type Result_phoneaddress struct {
	Province string `json:"province"`
	City string `json:"city"`
	Areacode string `json:"areacode"`
	Zip string `json:"zip"`
	Company string `json:"company"`
	Card string `json:"card"`
}