package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var urls []string
var ipBlocksStruct []IpBlock

func GetAllCountryUrl() []string {
	// Make HTTP request
	response, err := http.Get("https://github.com/ipverse/rir-ip/tree/master/country")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)

	return urls
}

// GetIpUrls := TAKE COUNTRY URL && RETURN IPV4URL
func GetIpUrls(url string) string {
	// Make HTTP request
	response, err := http.Get("https://github.com/" + url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier

	var _urls []string
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		// See if the href attribute exists on the element
		href, exists := element.Attr("href")
		if exists {
			//fmt.Println(href)
			_urls = append(_urls, href)
		}
	})

	var finalUrl string
	for i := 0; i < len(_urls); i++ {
		if strings.Contains(_urls[i], "ipv4-aggregated.txt") == true {
			finalUrl = _urls[i]
		}
	}

	return finalUrl
}

func GetIpv4Urls() []string {
	var ipV4urls []string

	urls := GetAllCountryUrl()

	for i := 0; i < len(urls); i++ {
		if strings.Contains(urls[i], "ipverse/rir-ip/tree/master/country") == true {
			ipV4urls = append(ipV4urls, urls[i]+"/ipv4-aggregated.txt")
		}
	}

	return ipV4urls
}

func GetIpBlocks() []IpBlock {
	urls := GetIpv4Urls()
	for i := 0; i < len(urls); i++ {
		response, err := http.Get("https://github.com" + urls[i])
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		// Read response data in to memory
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal("Error reading HTTP body. ", err)
		}

		//reCountry := regexp.MustCompile("")
		//country := reCountry.FindAllString(urls[i], -1)
		//fmt.Println(country)

		// Create a regular expression to find ipBlocks
		re := regexp.MustCompile("([0-9]{1,3}).([0-9]{1,3}).([0-9]{1,3}).([0-9]{1,3})[/]([0-9]{2})")
		ipBlocks := re.FindAllString(string(body), -1)
		fmt.Println("url: ", urls[i])
		var temp IpBlock
		if ipBlocks == nil {
			fmt.Println("No matches.")
		} else {
			for _, ipBlock := range ipBlocks {
				fmt.Println(ipBlock)
				temp.Cidr = ipBlock
				countryCode := GetCountryCode(urls[i])
				temp.Country = countryCode
				ipBlocksStruct = append(ipBlocksStruct, temp)
			}
		}
	}
	//fmt.Println(ipBlocksStruct)
	fmt.Println("count: ", len(ipBlocksStruct))
	return ipBlocksStruct
}

func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
		urls = append(urls, href)
	}
}

func GetCountryCode(val string) string {
	return strings.ToUpper(strings.Split(strings.Split(val, "/")[6], ".")[0])
}
