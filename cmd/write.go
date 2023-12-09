package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func Generate() {
	dir := viper.GetString("dir")
	filename := viper.GetString("filename")
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Fatal(err)
	}
	outputFile := dir + "/" + filename
	os.WriteFile(outputFile, []byte(FwTemplate()), 0777)
	fmt.Println("File Created : ", outputFile)
	fmt.Println("Start Watching by running:")
	fmt.Println("go run ", outputFile)
	fmt.Println()
}
