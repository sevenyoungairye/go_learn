package httpclient

import (
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"regexp"
	"time"
	"top.lel.dn/main/pkg/logger"
)

const (
	WIN10   = "win10"
	MAC     = "mac"
	Android = "android"
	IOS     = "IOS"
	IPAD    = "Ipad"
)

var AgentArr []string
var Headers map[string]string

func init() {
	AgentArr = []string{WIN10, MAC, Android, IOS, IPAD}
	Headers = map[string]string{}
	Headers[WIN10] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.3"
	Headers[MAC] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4280.88 Safari/537.36"
	Headers[IPAD] = "Mozilla/5.0 (iPad; CPU OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.77 Mobile/15E148 Safari/604.1"
	Headers[Android] = "Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Mobile Safari/537.36"
	Headers[IOS] = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"
}

// GetRandomUA get random user-agent.
func GetRandomUA() string {
	arrLen := len(AgentArr)
	rand.Seed(time.Now().UnixNano())
	// [0 - arrLen)
	i := rand.Intn(arrLen)
	return Headers[AgentArr[i]]
}

func GetDateByAttrSelector(url, htmlSelector string) string {
	res := ""

	request := initRequest(nil)

	// E[foo="bar"]
	// span[property="v:initialReleaseDate"]
	// 2022-06-10(韩国)
	request.OnHTML(htmlSelector, func(e *colly.HTMLElement) {
		re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
		res = re.FindString(e.Text)
	})
	// get json... then use regexp match
	//request.OnHTML("script[type=\"application/ld+json\"]", func(e *colly.HTMLElement) {
	//})

	_ = request.Visit(url)

	return res
}

// HttpWithGet get way
func HttpWithGet(URL string, headers map[string]string) string {
	logger.Info(fmt.Sprintf("request url %s", URL))

	var retStr = ""

	c := initRequest(headers)

	c.OnResponse(func(resp *colly.Response) {

		data := string(resp.Body)
		retStr = data
		// this can call to resp.Save(data)
		logger.Info(fmt.Sprintf("response code: %d, response body: %s", resp.StatusCode, data))
	})

	err := c.Visit(URL)
	if err != nil {
		logger.Warn(fmt.Sprintf("http get invoke error, req url: %s, err: %s", URL, err))
	}
	return retStr
}

// HttpWithPost post way
func HttpWithPost(URL string, reqMap map[string]string, headers map[string]string) string {
	logger.Info(fmt.Sprintf("request url: %s, request body: %s", URL, reqMap))
	c := initRequest(headers)
	c.OnRequest(func(r *colly.Request) {
		// 请求头json发送
		r.Headers.Add("content-type", "application/json")
	})

	c.OnResponse(func(resp *colly.Response) {
		data := string(resp.Body)
		logger.Info(fmt.Sprintf("response code: %d, response body: %s", resp.StatusCode, data))
	})

	err := c.Post(URL, reqMap)
	if err != nil {
		logger.Warn(fmt.Sprintf("http post invoke error, req url: %s, err: %s", URL, err))
	}

	return ""
}

func initRequest(headers map[string]string) *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		if headers != nil && len(headers) > 0 {
			for k, v := range headers {
				r.Headers.Add(k, v)
			}
		}
		r.Headers.Add("user-agent", GetRandomUA())
		r.Headers.Add("accept-language", "zh-CN,zh;q=0.9,ja;q=0.8")
		r.Headers.Add("accept", "application/json, text/javascript, */*; q=0.01")
	})
	return c
}
