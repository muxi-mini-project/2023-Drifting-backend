package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type accountReqeustParams struct {
	lt         string
	execution  string
	_eventId   string
	submit     string
	JSESSIONID string
}

type SuInfo struct {
	Errcode string `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	User    struct {
		DeptID       string `json:"deptId"`
		DeptName     string `json:"deptName"`
		ID           string `json:"id"`
		Mobile       string `json:"mobile"`
		Name         string `json:"name"`
		SchoolEmail  string `json:"schoolEmail"`
		Status       int    `json:"status"`
		UserFace     string `json:"userFace"`
		Username     string `json:"username"`
		Usernumber   string `json:"usernumber"`
		Usertype     string `json:"usertype"`
		UsertypeName string `json:"usertypeName"`
		Xb           string `json:"xb"`
	} `json:"user"`
}

func GetUserInfoFormOne(sid string, pwd string) (SuInfo, error) {
	var suInfo SuInfo
	params, err := makeAccountPreflightRequest()
	if err != nil {
		log.Println(err)
		return suInfo, err
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Println(err)
		return suInfo, err
	}
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
		Jar:     jar,
	}
	err = makeAccountRequest(sid, pwd, params, &client)
	// err := MakeAccountRequest( "", "", params, &client)
	if err != nil {
		log.Println(err)
		return suInfo, err
	}
	// MakeXKRequest(&client)
	pt, err := MakeONERequest(&client)
	if err != nil {
		log.Println(err)
		return suInfo, err
	}
	pt = "Bearer " + pt
	suInfo = getInfo(pt)
	return suInfo, nil
}

func MakeONERequest(client *http.Client) (portal_token string, err error) {
	request, err := http.NewRequest("GET", "http://one.ccnu.edu.cn", nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return "", err
	}

	u, err := url.Parse("http://one.ccnu.edu.cn")
	if err != nil {
		log.Println(err)
		return "", err
	}

	var pt string

	for _, cookie := range client.Jar.Cookies(u) {
		if cookie.Name == "PORTAL_TOKEN" {
			pt = cookie.Value
		}
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
	return pt, nil
}

// 预处理，打开 account.ccnu.edu.cn 获取模拟登陆需要的表单字段
func makeAccountPreflightRequest() (*accountReqeustParams, error) {
	var JSESSIONID string
	var lt string
	var execution string
	var _eventId string

	params := &accountReqeustParams{}

	// 初始化 http client
	client := http.Client{
		//Timeout: TIMEOUT,
	}

	// 初始化 http request
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login", nil)
	if err != nil {
		log.Println(err)
		return params, err
	}
	//request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	// 发起请求
	resp, err := client.Do(request)
	if err != nil {

		log.Println(err)
		return params, err
	}

	// 读取 Body
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		return params, err
	}

	// 获取 Cookie 中的 JSESSIONID
	for _, cookie := range resp.Cookies() {
		//fmt.Println(cookie.Value)
		if cookie.Name == "JSESSIONID" {
			JSESSIONID = cookie.Value
		}
	}

	if JSESSIONID == "" {
		log.Println("Can not get JSESSIONID")
		return params, errors.New("can not get JSESSIONID")
	}

	// 正则匹配 HTML 返回的表单字段
	ltReg := regexp.MustCompile("name=\"lt\".+value=\"(.+)\"")
	executionReg := regexp.MustCompile("name=\"execution\".+value=\"(.+)\"")
	_eventIdReg := regexp.MustCompile("name=\"_eventId\".+value=\"(.+)\"")

	bodyStr := string(body)

	ltArr := ltReg.FindStringSubmatch(bodyStr)
	if len(ltArr) != 2 {
		log.Println("Can not get form paramater: lt")
		return params, errors.New("can not get form paramater: lt")
	}
	lt = ltArr[1]

	execArr := executionReg.FindStringSubmatch(bodyStr)
	if len(execArr) != 2 {
		log.Println("Can not get form paramater: execution")
		return params, errors.New("can not get form paramater: execution")
	}
	execution = execArr[1]

	_eventIdArr := _eventIdReg.FindStringSubmatch(bodyStr)
	if len(_eventIdArr) != 2 {
		log.Println("Can not get form paramater: _eventId")
		return params, errors.New("can not get form paramater: _eventId")
	}
	_eventId = _eventIdArr[1]

	log.Println("Get params successfully", lt, execution, _eventId)

	params.lt = lt
	params.execution = execution
	params._eventId = _eventId
	params.submit = "LOGIN"
	params.JSESSIONID = JSESSIONID

	return params, nil
}

func makeAccountRequest(sid, password string, params *accountReqeustParams, client *http.Client) error {
	v := url.Values{}
	v.Set("username", sid)
	v.Set("password", password)
	v.Set("lt", params.lt)
	v.Set("execution", params.execution)
	v.Set("_eventId", params._eventId)
	v.Set("submit", params.submit)
	request, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login;jsessionid="+params.JSESSIONID, strings.NewReader(v.Encode()))
	if err != nil {
		log.Print(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	// check
	reg := regexp.MustCompile("class=\"success\"")
	matched := reg.MatchString(string(body))
	if !matched {
		log.Println("Wrong sid or pwd")
		return errors.New("wrong sid or pwd")
	}

	log.Println("Login successfully")
	return nil
}

func getInfo(pt string) SuInfo {
	client1 := http.Client{}
	v := url.Values{}

	request, _ := http.NewRequest("POST", "http://one.ccnu.edu.cn/user_portal/index", strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Authorization", pt)
	resp, err := client1.Do(request)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(body))
	var tmpInfo SuInfo
	err = json.Unmarshal(body, &tmpInfo)
	if err != nil {
		log.Println(err)
	}
	return tmpInfo
}
