package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Weather interface {
	GetForecast(city string) (ForecastResponse, error)
}

type ForecastResponse struct {
	City struct {
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		ID         int64  `json:"id"`
		Name       string `json:"name"`
		Population int64  `json:"population"`
		Sunrise    int64  `json:"sunrise"`
		Sunset     int64  `json:"sunset"`
		Timezone   int64  `json:"timezone"`
	} `json:"city"`
	Cnt  int64  `json:"cnt"`
	Cod  string `json:"cod"`
	List []struct {
		Clouds struct {
			All int64 `json:"all"`
		} `json:"clouds"`
		Dt    int64  `json:"dt"`
		DtTxt string `json:"dt_txt"`
		Main  struct {
			FeelsLike float64 `json:"feels_like"`
			GrndLevel int64   `json:"grnd_level"`
			Humidity  int64   `json:"humidity"`
			Pressure  int64   `json:"pressure"`
			SeaLevel  int64   `json:"sea_level"`
			Temp      float64 `json:"temp"`
			TempKf    float64 `json:"temp_kf"`
			TempMax   float64 `json:"temp_max"`
			TempMin   float64 `json:"temp_min"`
		} `json:"main"`
		Pop  float64 `json:"pop"`
		Rain struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		Visibility int64 `json:"visibility"`
		Weather    []struct {
			Description string `json:"description"`
			Icon        string `json:"icon"`
			ID          int64  `json:"id"`
			Main        string `json:"main"`
		} `json:"weather"`
		Wind struct {
			Deg   int64   `json:"deg"`
			Gust  float64 `json:"gust"`
			Speed float64 `json:"speed"`
		} `json:"wind"`
	} `json:"list"`
	Message int64 `json:"message"`
}

type Client struct {
	ApiKey string
	Host   string
	Path   string
}

func NewClient(apiKey string) *Client {
	return &Client{ApiKey: apiKey}
}

func (c *Client) GetForecast(city string) (ForecastResponse, error) {

	url := "http://api.openweathermap.org/data/2.5/forecast?q=" + city + "&units=metric&lang=pl&mode=json&appid=" + c.ApiKey
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := getHttpClient().Do(req)

	if err != nil {
		fmt.Println(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		fmt.Println(err2)
	}

	f := ForecastResponse{}

	jsonErr := json.Unmarshal(body, &f)

	if jsonErr != nil {
		return ForecastResponse{}, err
	}

	return f, nil
}

func getHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 1,
	}
}
