package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
)

//go:embed fw/fw.go
var fwGo []byte

func main() {
	dir := flag.String("d", "./watch/fw", "Output Directory")
	filename := flag.String("f", "fw.go", "File name without directory, example 'fw.go'")
	flag.Parse()

	err := os.MkdirAll(*dir, 0777)
	if err != nil {
		log.Fatal(err)
	}
	outputFile := *dir + "/" + *filename
	f, err := os.Create(outputFile)
	if err != nil {
		log.Panicf("os.Create[%T]: %v", err, err)
	}
	defer f.Close()
	_, err = f.Write(fwGo)
	if err != nil {
		log.Panicf("f.Write[%T]: %v", err, err)
	}
	fmt.Println("File Created : ", outputFile)
	fmt.Println("Start by running:")
	fmt.Println("go run ", outputFile)
}
