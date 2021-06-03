package recuperator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Data struct {
	SupplyAir       string `json:"supplyAir"`
	ExhaustAir      string `json:"exhaustAir"`
	ExtractAir      string `json:"extractAir"`
	OutsideAir      string `json:"outsideAir"`
	Humidity        string `json:"humidity"`
	SupplyFanSpeed  string `json:"supplyFanSpeed"`
	ExtractFanSpeed string `json:"extractFanSpeed"`
	FanSpeed        string `json:"fanSpeed"`
	Temperature     string `json:"temperature"`
}

type Temperature struct {
	Temperature string `json:"temperature"`
}

type FanSpeed struct {
	FanSpeed string `json:"fanSpeed"`
}

type Recuperator interface {
	GetData() (Data, error)
	GetTemperature() (Temperature, error)
	SetTemperature(t int) (Temperature, error)
	SetFanSpeed(s int) (FanSpeed, error)
}

type Client struct {
	IpAddress string
	Login     string
	Password  string
}

func NewClient(ipAddress string, login string, password string) *Client {
	return &Client{IpAddress: ipAddress, Login: login, Password: password}
}

func (c *Client) GetData() (Data, error) {
	f := "FUNC(4,1,4,0,24)"
	req, err := newRequest(f, c)
	resp, err := getHttpClient().Do(req)

	if err != nil {
		return Data{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	a := strings.Split(string(body), ";")
	t, err := c.GetTemperature()

	if err != nil {
		return Data{}, err
	}

	return Data{
		SupplyAir:       parseTemp(a[0]),
		ExhaustAir:      parseTemp(a[6]),
		ExtractAir:      parseTemp(a[6]),
		OutsideAir:      parseTemp(a[9]),
		Humidity:        a[13],
		SupplyFanSpeed:  a[15],
		ExtractFanSpeed: a[16],
		Temperature:     t.Temperature,
		FanSpeed:        fanSpeed(a[15]),
	}, nil
}

func (c *Client) GetTemperature() (Temperature, error) {
	f := "FUNC(4,1,3,0,111)"
	req, err := newRequest(f, c)
	resp, err := getHttpClient().Do(req)

	if err != nil {
		return Temperature{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	a := strings.Split(string(body), ";")

	return Temperature{
		Temperature: a[1],
	}, nil
}

func (c *Client) SetTemperature(t int) (Temperature, error) {
	f := "FUNC(4,1,6,1," + strconv.Itoa(t) + ")"
	req, err := newRequest(f, c)
	resp, err := getHttpClient().Do(req)

	if err != nil {
		return Temperature{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	a := strings.Split(string(body), ";")

	return Temperature{
		Temperature: a[0],
	}, nil
}

func (c *Client) SetFanSpeed(s int) (FanSpeed, error) {
	f := "FUNC(4,1,6,0," + strconv.Itoa(s) + ")"
	req, err := newRequest(f, c)
	resp, err := getHttpClient().Do(req)

	if err != nil {
		return FanSpeed{}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	a := strings.Split(string(body), ";")

	return FanSpeed{
		FanSpeed: a[0],
	}, nil
}

func newRequest(f string, c *Client) (*http.Request, error) {
	req, err := http.NewRequest("GET", getUrl(c, f), nil)
	req.SetBasicAuth(c.Login, c.Password)

	return req, err
}

func getUrl(c *Client, f string) string {
	return "http://" + c.IpAddress + "/" + f
}

func getHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 1,
	}
}

func parseTemp(t string) string {
	s, err := strconv.ParseFloat(t, 32)
	if err != nil {
		return "0.0"
	}
	return fmt.Sprintf("%.2f", s/10)
}

func fanSpeed(s string) string {
	if s == "0" {
		return "0"
	}

	if s == "30" {
		return "1"
	}

	if s == "60" {
		return "2"
	}

	if s == "80" {
		return "3"
	}

	if s == "100" {
		return "3"
	}

	return "4"
}
