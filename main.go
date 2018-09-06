package main

import (
	"io/ioutil"
	"time"
	"os"
	"net/http"
	"os/exec"
	"log"
	"fmt"
	"encoding/json"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	WgetUrl []string `json:"wgetUrl"`
	CheckIpUrl string `json:"checkIpUrl"`
}

var fileLocation = "./lastIp"

// Panic if error occurs.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
// Write ip to the file.
func writeToFile(ip string) {
	ipBytes := []byte(ip)
	err := ioutil.WriteFile(fileLocation, ipBytes, 0644)
	check(err)
}

// Reding saved ip from file.
func readFromFile() string {
	bytes, err := ioutil.ReadFile(fileLocation)
	check(err)
	lastIp := string(bytes)
	return lastIp
}
// Function will perform update to the dynamic DNS service.
func updateIpToServer(username string, password string, urls []string) {
	for _, url := range urls {
		cmd := exec.Command("wget", "--delete-after", "--no-check-certificate", "--no-proxy", "--user", username, "--password", password, url)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
// Load configuration from file.
func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

// Main function that contains loop where the IP address is checked.
func main() {
	config := loadConfiguration("./config.json")
	lastIp := readFromFile()
	for {
		resp, _ := http.Get(config.CheckIpUrl)
		bytes, _ := ioutil.ReadAll(resp.Body)
		ip := string(bytes)
		if(resp.StatusCode == 200) {
			if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
				// File doesn't exist so creating one
				writeToFile(ip)
			} else {
				if (lastIp != ip) {
					// Ip changed
					writeToFile(ip)
					lastIp = ip
					updateIpToServer(config.Username, config.Password, config.WgetUrl)
				}
			}
		}
		resp.Body.Close()
		time.Sleep(10000 * time.Millisecond)
	}
}

