package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// 用于接受用户参数
type Arg struct {
	account     string // 身份证
	address     string // 签到地址
	temperature string // 体温
}

func main() {

	var arg Arg

	flag.StringVar(&(arg.account), "id", "", "打卡必须值")
	flag.StringVar(&(arg.address), "a", "山西省太原市尖草坪区长征街1号", "打卡必须值")
	flag.StringVar(&(arg.temperature), "t", "36.5", "打卡必须值")

	flag.Parse()

	fmt.Printf("用户的参数%+v\n", arg)

	province := arg.address[:strings.Index(arg.address, "省")+3]
	city := arg.address[strings.Index(arg.address, "省")+3 : strings.Index(arg.address, "市")+3]
	district := arg.address[strings.Index(arg.address, "市")+3 : strings.Index(arg.address, "区")+3]

	urlValues := url.Values{}
	urlValues.Add("wc_type", "否")
	urlValues.Add("title", arg.temperature)
	urlValues.Add("province", province)
	urlValues.Add("jk_type", "健康")
	urlValues.Add("jc_type", "否")
	urlValues.Add("is_verify", "0")
	urlValues.Add("district", district)
	urlValues.Add("city", city)
	urlValues.Add("address", arg.address)
	urlValues.Add("mobile", arg.account)

	resp, _ := http.PostForm("http://yx.ty-ke.com/Home/Monitor/monitor_add", urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
