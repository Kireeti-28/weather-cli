package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kireeti-28/weather-cli/api"
)

func main() {
	for {
		fmt.Print("-> ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		cityName := strings.Trim(text, "\n")

		client := api.GetClient()
		resp, err := client.GetCityWeather(cityName)

		if err != nil {
			log.Fatal(err)
		}

		prettyPrintWeather(resp)
	}
}

func prettyPrintWeather(resp api.CityResp) {
	fmt.Println(resp.Name + "'s Weather Details:")
	fmt.Println("------------------------------")
	fmt.Printf(" - Description: %v\n", resp.Weather[0].Description)
	if resp.Clouds.All != 0 {
		fmt.Printf(" - Cloud: %v%v\n", resp.Clouds.All, "%")
	}
	fmt.Printf(" - Humidity %v%v\n", resp.Main.Humidity, "%")
	fmt.Printf(" - Temperature: %vK\n", resp.Main.Temp)
	fmt.Printf(" - Min Temperature: %vK\n", resp.Main.TempMin)
	fmt.Printf(" - Max Temperature: %vK\n", resp.Main.TempMax)
	fmt.Printf(" - Pressure: %vhPa\n", resp.Main.Pressure)
	if resp.Rain.OneH != 0 {
		fmt.Printf(" - Rain: %vmm\n", resp.Rain.OneH)
	}
	fmt.Printf(" - Wind Speed: %vm/s\n", resp.Wind.Speed)
	fmt.Println("------------------------------")
}
