package main

import(
	"flag"
	"fmt"
	"kafkabrowser/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	version = "v0.0.1"
)

func SpalshScreen() {
	fmt.Printf("Starting Kafka Browser\n")
	fmt.Printf("Version: %s\n",version)
	fmt.Printf("---------------------------------------------------------------------\n")
}

type Configuration struct {
	Port string
}

func ReadFile(filename string) []byte {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	return dat
} 

func main() {
	var conf string
	flag.StringVar(&conf, "conf", "", "configuration file name")
	flag.Parse()
	appconfig := &Configuration{}
	err := yaml.Unmarshal([]byte(ReadFile(conf)), &appconfig)
	if err!= nil {
		fmt.Printf("Failed to read configuration\n")
		fmt.Printf("Error occurred:%v", err)
		os.Exit(1)
	}
	SpalshScreen()
	fmt.Printf("Listening on port: %v\n",appconfig.Port)
	server.ServeApp(appconfig.Port)
}