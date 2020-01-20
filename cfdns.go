package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
)

type Config struct {
	IPEndpoint string `json:"IPEndpoint"`
	Sleep      int    `json:"Sleep"`
	Records    []struct {
		Username string `json:"Username"`
		APIKey   string `json:"API-Key"`
		Zone     string `json:"Zone"`
		Entry    string `json:"Entry"`
	} `json:"Records"`
}

func main() {
	ConfigFile := flag.String("config", "config.json", "Config File")
	flag.Parse()

	log.Println("Using: ", *ConfigFile)

	jsonFile, err := os.Open(*ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteVal, _ := ioutil.ReadAll(jsonFile)

	var config Config
	json.Unmarshal(byteVal, &config)

	for {
		for _, record := range config.Records {

			currentIP, err := GetCurrentIP(config.IPEndpoint)
			if err != nil {
				log.Println(err)
				continue
			}

			api, err := cloudflare.New(record.APIKey, record.Username)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("Updating record \"%v\" from zone \"%v\"", record.Entry, record.Zone)

			zoneID, err := api.ZoneIDByName(record.Zone)
			if err != nil {
				log.Println(err)
				continue
			}

			currRecords, err := api.DNSRecords(zoneID, cloudflare.DNSRecord{Type: "A", Name: record.Entry})
			if err != nil {
				log.Println(err)
				continue
			}

			if len(currRecords) == 0 {
				log.Println("Found DNS record but it is not type A")
				continue
			}

			if currentIP == currRecords[0].Content {
				log.Println("Current IP matches Cloudflare's record, no need to update")
				continue
			}

			DNSRecord := cloudflare.DNSRecord{TTL: 1, Type: "A", Name: record.Entry, Content: currentIP}
			err = api.UpdateDNSRecord(zoneID, currRecords[0].ID, DNSRecord)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("Successfully updated record")
		}
		log.Printf("Next update in %v seconds", config.Sleep)
		time.Sleep(time.Duration(config.Sleep) * time.Second)
	}
}

func GetCurrentIP(Endpoint string) (IP string, err error) {
	res, err := http.Get(Endpoint)
	if err != nil {
		return
	}

	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	IP = string(content)
	IP = strings.TrimRight(IP, "\r\n")
	return
}
