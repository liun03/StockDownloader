package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	inputDate := "20200911"
	url := fmt.Sprintf("https://www.twse.com.tw/exchangeReport/MI_INDEX?response=csv&date=%s&type=ALLBUT0999", inputDate)
	err := Downloader("/gotmp/", "stock.csv", url)
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
