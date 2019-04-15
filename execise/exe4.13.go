package exe

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		DownloadImage(url)
	}
}

func DownloadImage(url string) {
	//"http://img.omdbapi.com/?i=tt3896198&apikey=53ff5b13"
	var client = http.Client{}
	reqImg, err := client.Get(url)
	if err != nil {
		log.Fatalf("http.Get -> %v", err)
		return
	}
	data, err := ioutil.ReadAll(reqImg.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}
	reqImg.Body.Close()
	ioutil.WriteFile("movie.png", data, 0666)

	log.Println("I saved your image buddy!")
}
