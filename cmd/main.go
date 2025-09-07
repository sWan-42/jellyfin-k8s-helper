package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"
)

type NetworkConfiguration struct {
	XMLName      xml.Name `xml:"NetworkConfiguration"`
	KnownProxies []string `xml:"KnownProxies>string"`
}

func main() {
	for {
		fmt.Println("Trying Load the XML file")
		data, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
		if err != nil {
			panic(err)
		}

		fmt.Println("Trying Unmarshal only the KnownProxies section")
		var config NetworkConfiguration
		if err := xml.Unmarshal(data, &config); err != nil {
			panic(err)
		}

		fmt.Println("Trying Replace KnownProxies with new values from updateIPs()")
		config.KnownProxies = getIngressIPs()

		fmt.Println("Trying Marshal the updated KnownProxies")
		proxyXML, err := xml.MarshalIndent(config.KnownProxies, "  ", "    ")
		if err != nil {
			panic(err)
		}

		fmt.Println("Trying Replace the old KnownProxies block in the original XML")
		startTag := "<KnownProxies>"
		endTag := "  </KnownProxies>"
		startIdx := strings.Index(string(data), startTag)
		endIdx := strings.Index(string(data), endTag) + len(endTag)

		if startIdx == -1 || endIdx == -1 {
			panic("KnownProxies section not found")
		}

		updated := string(data[:startIdx]) +
			startTag + "\n" +
			string(proxyXML) + "\n" +
			endTag +
			string(data[endIdx:])

		fmt.Println("Trying Save the updated XML")
		if err := os.WriteFile(os.Getenv("CONFIG_PATH"), []byte(updated), 0644); err != nil {
			panic(err)
		}

		fmt.Println("KnownProxies updated successfully.")

		fmt.Println("Sleeping for 5 minutes...")
		time.Sleep(5 * time.Minute)
	}
}
