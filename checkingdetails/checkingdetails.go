package checkingdetails

import (
	"encoding/json"
	"fmt"
	"os"
)

type WebsiteUrl struct {
	targeting_website_details string `json"website_url"`
	website_timeout           int    `json"website_timeout"`
}

var websitedetails = WebsiteUrl{
	targeting_website_details: "sandbox.swift.com",
	website_timeout:           30,
}

func CheckWebsiteDetails() {

	website_data_information, website_error := json.MarshalIndent(
		websitedetails, "", " ")
	if website_error != nil {
		fmt.Println("An error occured: ", website_error)
		return
	}
	website_error = os.WriteFile("websitedtails.json", website_data_information, 0644)
	if website_error != nil {
		fmt.Println("Checking website error: ", website_error)
		return
	}
	fmt.Println()
}
