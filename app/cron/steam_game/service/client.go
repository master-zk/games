package service

type Client struct {
	WithCookie     bool
	Cookie         string
	AcceptLanguage string
	CC             string
	Host           string
}

func NewClient(cookie string) *Client {
	return &Client{
		WithCookie:     false,
		Cookie:         "",
		AcceptLanguage: "zh-CN,zh;q=0.9",
		Host:           "store.steampowered.com",
	}
}

// SetCookie 用于获取有色情/暴力内容的游戏
func (c *Client) SetCookie(cookie string) *Client {
	c.WithCookie = true
	c.Cookie = cookie

	return c
}

// ClearCookie 清除cookie
func (c *Client) ClearCookie() *Client {
	c.WithCookie = false
	c.Cookie = ""

	return c
}

func (c *Client) GetCookie() string {
	// TODO 读取数据库或者缓存
	cookie := "browserid=3713865080483596555; sessionid=415ef369603ab037a7507853; timezoneOffset=28800,0; steamCountry=US%7C9f5097f739594bb6202d83fe65cb993e; recentapps=%7B%221085660%22%3A1721792507%7D; app_impressions=730@1_4_4__129_1|1085660@1_4_4__129_1|1675200@1_4_4__147|359550@1_4_4__137_1|892970@1_4_4__137_1|2881650@1_4_4__137_1|2591280@1_4_4__145_2|1286990@1_4_4__145_1|1771980@1_4_4__145_5|2224640@1_4_4__145_4|1619230@1_4_4__145_3|1540210@1_4_4__141_1|587430@1_4_4__141_1|1059530:1059550:1059570@1_4_4__141_1|3111170@1_4_4__42|1857090@1_4_4__42|2448970@1_4_4__42|1172470@1_7_7_7000_150_1|230410@1_7_7_7000_150_1|3015760@1_7_7_7000_150_1|1144200@1_7_7_7000_150_1|1245620@1_7_7_7000_150_1|2074920@1_7_7_7000_150_1|730@1_7_7_7000_150_1|1675200@1_7_7_7000_150_1"

	return cookie
}

// SetCC 设置国家地区
func (c *Client) SetCC(cc string) *Client {
	c.CC = cc

	return c
}
