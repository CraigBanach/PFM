package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	res, err := http.Post("http://posttestserver.com/post.php", "application/x-www-form-urlencoded", url.QueryEscape("Authorization: oauth_callback=http://www.bbc.co.uk/news"))
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	fmt.Printf("%s", url.QueryEscape("Authorization: oauth_callback=http://www.bbc.co.uk/news"))
}
