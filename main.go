package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/atotto/clipboard"
	"github.com/joho/godotenv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func httpGet(url string) []byte {
	res, err := http.Get(url)
	check(err)
	bytes, err := ioutil.ReadAll(res.Body)
	check(err)
	res.Body.Close()
	return bytes
}

func checkIPStatus(IP string) {
	var hasChanged = false
	if _, err := os.Stat("/tmp/myip"); err != nil {
		currentIPFile, err := os.Create("/tmp/myip")
		check(err)
		currentIPFile.WriteString(IP)
		currentIPFile.Sync()
		hasChanged = true
	} else {
		currentIPFile, err := ioutil.ReadFile("/tmp/myip")
		check(err)
		var currentIP = string(currentIPFile)
		if IP != currentIP {
			hasChanged = true
			err := ioutil.WriteFile("/tmp/myip", []byte(IP), 0644)
			check(err)
		}
	}
	clipboard.WriteAll(IP + "/32")
	if hasChanged {
		fmt.Println("IP has changed and copied Successfully to Clipboard")
	} else {
		fmt.Println("IP hasn't changed since last time!!!")
	}
}

func getGeoLocationInfo(IP string) {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	apiKey := os.Getenv("IPIFY_API_KEY")
	if apiKey == "" {
		return
	}
	geoLocation := httpGet("https://geo.ipify.org/api/v1?apiKey=" + apiKey + "&ipAddress=" + IP)
	var data map[string]interface{}
	json.Unmarshal(geoLocation, &data)
	location := data["location"].(map[string]interface{})
	fmt.Println()
	fmt.Println("current IP>> ", IP)
	fmt.Println("City>>       ", location["city"])
	fmt.Println("Country>>    ", location["country"])
	fmt.Println("Region>>     ", location["region"])
}

func getIP() string {
	IPBytes := httpGet("http://ifconfig.me/ip")
	return string(IPBytes)
}

func main() {
	IP := getIP()
	checkIPStatus(IP)
	getGeoLocationInfo(IP)
}
