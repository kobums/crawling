package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	c := make(chan string)
	startNum := 1
	endNum := 2
	for i := startNum; i < endNum; i++ {
		go check(i, c)
	}
	for j := startNum; j < endNum; j++ {
		fmt.Println(<-c)
	}
}


func checkErr(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}

func checkCode(res *http.Response) {
    if res.StatusCode != 200 {
        log.Fatalf("Status code err: %d %s", res.StatusCode, res.Status)
    }
}

func check(i int, c chan <- string) {
	url := "https://booktoki283.com/novel/5199024?stx=%EA%B2%80%EC%88%A0%EB%AA%85%EA%B0%80&book=%EC%9D%BC%EB%B0%98%EC%86%8C%EC%84%A4&spage="
	index := strconv.Itoa(i)

	res, err := http.Get(url + index)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(res)
	checkErr(err)
	
	// aaa, _ := doc.Find("#novel_content").Html()
	aaa, _ := doc.Find("div").Html()
	fmt.Println(aaa)
	c <- aaa
}
