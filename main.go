package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Anakros/imgur-uploader/imgur"
	"github.com/spf13/viper"
)

var (
	inputFile  = flag.String("i", "", "Input File path")
	deleteHash = flag.String("d", "", "Delete Hash")
)

func main() {
	flag.Parse()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/imgur-uploader")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Can't find config file: %s", err)
	}

	clientID := viper.GetString("ClientID")

	imgur.ClientID = clientID

	switch {
	case len(*inputFile) > 0:
		answer, err := imgur.Upload(*inputFile)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Image was successfully uploaded.\n"+
			"Link: %s\n"+
			"Delete Hash: %s\n", answer.Link, answer.DeleteHash)
	case len(*deleteHash) > 0:
		_, err := imgur.Delete(*deleteHash)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Image was successfully deleted.")
	}
}
