package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	err := Downloader("/gotmp/", "stock.csv", "https://www.twse.com.tw/exchangeReport/MI_INDEX?response=csv&date=20200914&type=ALLBUT0999")
	if err != nil {
		fmt.Println(err)
	}
}

func Downloader(savedPath string, fileName string, url string) error {
	if _, err := os.Stat(savedPath); os.IsNotExist(err) {
		os.Mkdir(savedPath, 0755)
	}
	outputFile, err := os.Create(savedPath + fileName)
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	defer outputFile.Close()
	_, err = io.Copy(outputFile, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
