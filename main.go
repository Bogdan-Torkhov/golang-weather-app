package main

import (
	"log"
	"fmt"
	// "os"
    owm "github.com/briandowns/openweathermap"
    "net/http"
    "net/url"
)

// var apiKey = os.Getenv("OWM_API_KEY")

const proxy = "http://38.55.182.255:80"

const apiKey = "6ab7fc8050c5668c33576f3a93d1deea"

func main() {
<<<<<<< HEAD
	i, err := url.Parse(proxy)
	if err != nil { panic(err) }
	cli := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(i)}}
	w, err := owm.NewCurrent("C", "ru", apiKey, owm.WithHttpClient(cli)) // celseus (imperial) with Russian output
=======
	w, err := owm.NewCurrent("C", "RU", apiKey) // fahrenheit (imperial) with Russian output
>>>>>>> refs/remotes/origin/main
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(apiKey)
	err = w.CurrentByName("Phoenix")
	if err != nil { panic(err) }
	fmt.Println(w.Weather)
}