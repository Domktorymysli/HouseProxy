package main

import (
	"fmt"
	"net/http"
	"recuperator/clu"
	"recuperator/config"
	"recuperator/http_api"
	"recuperator/recuperator"
	"recuperator/weather"
	"time"
)

func getListenAddress(port string) string {
	return ":" + port
}

func main() {
	cnf := config.LoadConfiguration("config.json")
	fromIp := "192.168.0.1"

	fmt.Println("--------------------------------------------------")
	fmt.Println("--              HOME PROXY 1.0                  --")
	fmt.Println("--------------------------------------------------")

	r := recuperator.NewClient(cnf.Recuperator.Ip, cnf.Recuperator.Login, cnf.Recuperator.Password)
	w := weather.NewClient(cnf.Weather.ApiKey)
	c := clu.NewClient(cnf.Clus[0].Ip, cnf.Clus[0].Port, cnf.Clus[0].Key, cnf.Clus[0].Iv, fromIp)

	startRecuperatorTicker(r, c)
	startWeatherTicker(w, c, cnf.Weather.City)

	http_api.RegisterHandlers(r)

	if err := http.ListenAndServe(getListenAddress(cnf.Server.Port), nil); err != nil {
		panic(err)
	}
}

func startWeatherTicker(w weather.Weather, c clu.Clu, city string) {
	go func() {
		for now := range time.Tick(time.Minute * 1) {
			fmt.Println("weather: ", now)

			f, err := w.GetForecast(city)

			if err != nil {
				fmt.Println(err)
			}

			if len(f.List) > 1 {

				sunrise := time.Unix(f.City.Sunrise, f.City.Timezone)
				sunset := time.Unix(f.City.Sunset, f.City.Timezone)

				c.SetVar("\"weatherSunrise\"", "\""+sunrise.Format(time.Kitchen)+"\"")
				c.SetVar("\"weatherSunset\"", "\""+sunset.Format(time.Kitchen)+"\"")
				c.SetVar("\"weatherDescription\"", "\""+f.List[0].Weather[0].Description+"\"")
				c.SetVar("\"weatherFeelsLike\"", fmt.Sprintf("%f", f.List[0].Main.FeelsLike))
				c.SetVar("\"weatherPressure\"", "10") //fmt.Sprintf("%i", f.List[0].Main.Pressure))
				c.SetVar("\"weatherTemperature\"", fmt.Sprintf("%f", f.List[0].Main.Temp))
			}
		}
	}()
}

func startRecuperatorTicker(r recuperator.Recuperator, c clu.Clu) {
	go func() {
		for now := range time.Tick(time.Second * 5) {
			fmt.Println("recuperator", now)

			d, err := r.GetData()
			if err != nil {
				fmt.Println(err)
			}

			c.SetVar("\"recuperatorTemperature\"", d.Temperature)
			c.SetVar("\"recuperatorFanSpeed\"", d.FanSpeed)
			c.SetVar("\"recuperatorSupplyAir\"", d.SupplyAir)
			c.SetVar("\"recuperatorExhaustAir\"", d.ExhaustAir)
			c.SetVar("\"recuperatorExtractAir\"", d.ExtractAir)
			c.SetVar("\"recuperatorOutsideAir\"", d.OutsideAir)
			c.SetVar("\"recuperatorHumidity\"", d.Humidity)
			fmt.Println("Recuperator updated:", d)
		}
	}()
}
