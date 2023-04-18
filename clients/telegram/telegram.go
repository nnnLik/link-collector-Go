package telegram

import (
	"encoding/json"
	"io"
	"link-collector-bot/lib/e"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host    string
	baseURL string
	client  http.Client
}

const(
	getUpdateMethod = "getUpdate"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) Client {
	return Client{
		host: host,
		baseURL: newbaseUrl(token),
		client: http.Client{},
	}
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}


func (c *Client) Updates(offset int, limit int) (updates []Update, err error) {
	defer func() { err = e.WrapIsErr("error creating request", err) }()
	
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))
	
	data, err := c.doRequest(getUpdateMethod, q)
	if err != nil {
		return nil, err
	}
	
	var res UpdateResponse
	
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	
	return res.Result, nil
}

func newbaseUrl(token string) string {
	return "bot" + token
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIsErr("error creating request", err) }()
	const errMsg = "error creating request"
	
	u := url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.baseURL, method),
	}
	
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {_ = response.Body.Close()}()
	
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}