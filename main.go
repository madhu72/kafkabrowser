package main

import(	
	"flag"
	"fmt"
	"kafkabrowser/app"
	"kafkabrowser/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	version = "v0.0.1"
)

type Configuration struct {
	Port string
	Rep *app.DbConnector
}

func (config *Configuration) SpalshScreen() {
	fmt.Printf("Starting Kafka Browser\n")
	fmt.Printf("Version: %s\n",version)
	fmt.Printf("---------------------------------------------------------------------\n")
	fmt.Printf("Listening on port: %v\n",config.Port)
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
	var initdb bool
	flag.StringVar(&conf, "conf", "", "configuration file name")
	flag.BoolVar(&initdb, "initdb", false, "initialise database")
	flag.Parse()
	appconfig := &Configuration{}
	err := yaml.Unmarshal([]byte(ReadFile(conf)), &appconfig)
	if err!= nil {
		fmt.Printf("Failed to read configuration\n")
		fmt.Printf("Error occurred:%v", err)
		os.Exit(1)
	}
	appconfig.Rep = &app.DbConnector{Dbname:"kafkabrowser.db"}
	err =  appconfig.Rep.Connect()
	if initdb {
		err = appconfig.Rep.Initialise()
		if err != nil {
			fmt.Printf("Error occurred when establish connection with Database.\nError:%v\n", err)
			os.Exit(1)
		}
		err = appconfig.Rep.AddDefaultConfig("Default Kafka Config",fmt.Sprintf("Broker: localhost:9092\nTopic: test\nPartition: 0\n"))
		if err != nil {
			fmt.Printf("Failed to create default config.\nError occurred:%v",err)
			os.Exit(1)
		}
		os.Exit(1)
	}
	appconfig.SpalshScreen()
	server.ServeApp(appconfig.Port)
}

