package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func main() {
	url := "https://api.openweathermap.org/data/2.5/weather?id=1850144&appid=" + os.Getenv("OPEN_WEATHER_API_KEY")

	res, _ := http.Get(url)

	// レスポンスを取得
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var data WeatherData

	// 返却されたJSONをパース
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	// 取得したものを出力（一部）
	fmt.Println("Main：" + data.Weather[0].Main)
	fmt.Println("Description：" + data.Weather[0].Description)
	fmt.Println("Temp：" + strconv.FormatFloat(data.Main.Temp, 'f', 4, 64))
	fmt.Println("TempMax：" + strconv.FormatFloat(data.Main.TempMax, 'f', 4, 64))
	fmt.Println("TempMin：" + strconv.FormatFloat(data.Main.TempMin, 'f', 4, 64))
	fmt.Println("Humidity：" + strconv.Itoa(data.Main.Humidity))
	fmt.Println("Humidity：" + strconv.Itoa(data.Main.Pressure))
}
