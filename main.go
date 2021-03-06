package main

//Run the below command when ever the static files changed from public folder
//go-bindata -fs -prefix "public/" public/...

import(	
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"kafkabrowser/app"
	"kafkabrowser/server"

	"gopkg.in/yaml.v2"
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
		err = appconfig.Rep.Initialize()
		if err != nil {
			fmt.Printf("Error occurred when initialize Database.\nError: %v\n", err)
			os.Exit(1)
		}
		err = appconfig.Rep.AddDefaultConfig("Default Kafka Config",
				fmt.Sprintf("---\nBrokers: localhost:9092\nTopic: test\nPartition: 0"))
		if err != nil {
			fmt.Printf("Failed to create default config.\nError occurred: %v",err)
			os.Exit(1)
		}
		fmt.Printf("Initialised database successfully.")
		os.Exit(1)
	}
	appconfig.SpalshScreen()
	server.ServeApp(appconfig.Port)
}

