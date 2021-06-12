package clu

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net"
	cypher "recuperator/clu/cypher"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// wyślij dane z http_api do clu
// omkthttps://gobyexample.com/tickers

type Clu interface {
	CallFunction(fn string, p ...string) (*Response, error)
	CallRaw(raw string) (*Response, error)
	SetVar(name string, value string) (*Response, error)
}

type Client struct {
	cluIp   string
	cluPort string
	key     []byte
	iv      []byte
	fromIp  string
}

func NewClient(cluIp string, cluPort string, cluKey string, cluIv string, fromIp string) *Client {
	key, _ := base64.StdEncoding.DecodeString(cluKey)
	iv, _ := base64.StdEncoding.DecodeString(cluIv)

	return &Client{
		cluIp:   cluIp,
		cluPort: cluPort,
		key:     key,
		iv:      iv,
		fromIp:  fromIp,
	}
}

func (c *Client) CallRaw(raw string) (*Response, error) {
	rsp, err := c.send([]byte(raw))

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *Client) CallFunction(fn string, p ...string) (*Response, error) {
	cmd := prepareFunctionCommand(fn, p, c.fromIp)
	rsp, err := c.send(cmd)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *Client) SetVar(name string, value string) (*Response, error) {
	return c.CallFunction("setVar", name, value)
}

func (c *Client) send(data []byte) (*Response, error) {
	msg, _ := cypher.Encrypt(c.key, c.iv, data)
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", c.cluIp+":"+c.cluPort)
	conn.Write(msg)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	_, err = bufio.NewReader(conn).Read(p)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
	}
	p = bytes.Trim(p, "\x00")
	resp, _ := cypher.Decrypt(c.key, c.iv, p)

	return NewResponse(string(resp[:])), nil
}

type Response struct {
	Msg       string
	Type      string
	Ip        string
	SessionId string
}

func NewResponse(msg string) *Response {
	re := regexp.MustCompile(`(req|resp):([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}):([a-z0-9]{6,8}):(.*)`)

	if re.Match([]byte(msg)) == false {
		return &Response{Type: "err", Msg: msg}
	}

	res := re.FindAllSubmatch([]byte(msg), -1)

	return &Response{
		Msg:       string(res[0][4][:]),
		Type:      string(res[0][1][:]),
		Ip:        string(res[0][2][:]),
		SessionId: string(res[0][3][:]),
	}
}

func prepareFunctionCommand(fn string, p []string, ip string) []byte {
	cmd := "req:" + ip + ":" + generateRandomSessionId() + ":" + fn + "(" + strings.Join(p, ",") + ");"
	return []byte(cmd)
}

func generateRandomSessionId() string {
	rand.Seed(time.Now().UnixNano())
	min := 10000000
	max := 90000000
	return strconv.Itoa(rand.Intn(max-min) + min)
}
