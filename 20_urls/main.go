package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.andacademy.com:8080/lp2/uiuxdesign-courses/?utm_source=google+ads&utm_medium=src+ui+ux+imp+exact+desktop+geo+pune&utm_campaign=pune-keywords&utm_campaignid=21299500784&utm_adgroupid=160438344697&utm_creativeid=699807720857&utm_matchtype=e&utm_device=c&utm_network=g&utm_keyword=ui%20ux%20classes%20in%20pune&utm_placement=&gad_source=1&gclid=CjwKCAjwhvi0BhA4EiwAX25uj6TC7keJR-5evMmaTet8zY58xcI6_aNREy3MvyL-DzrOdSmrBdKuOhoCclUQAvD_BwE"

func main() {

	// url paring
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Hostname())
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	fmt.Println("--------------------------")

	qparams := result.Query() // Query() parse rawQuery and retrun key value

	fmt.Printf("Type of qparams is %T\n", qparams)
	fmt.Println(qparams["utm_creativeid"])

	for index, value := range qparams {
		fmt.Printf("key = %v, value = %v\n", index, value)
	}

	// Format URL

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "SohamK.com",
		Path:    "/buyCoffee",
		RawPath: "user=demo",
	}

	anotherUrl := partsOfURL.String()
	fmt.Printf("url is %v\n", anotherUrl) //  https://SohamK.com/buyCoffee

}
