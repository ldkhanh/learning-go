package main 

import ("fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml")

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (L Location) String() string {
	return fmt.Sprintf(L.Loc)
}


func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)
	fmt.Printf("Here %s some %s", "are", "variables")
	for _, Location := range s.Locations {
		fmt.Printf("\n%s",Location)
	}
}