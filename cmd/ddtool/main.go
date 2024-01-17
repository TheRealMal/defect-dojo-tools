package main

import (
	"bufio"
	"defect-dojo-tools/pkg/ddclient"
	"log"
	"os"
	"strings"
)

// 1st cmd argument - command
// 2nd cmd argument - API URL
// 3rd cmd argument - API Token
func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) != 3 {
		log.Fatalf("incorrect arguments quantity: %d", len(cmdArgs))
	}
	command, url, token := cmdArgs[0], cmdArgs[1], cmdArgs[2]
	client := ddclient.Client{
		ApiURL:   url,
		ApiToken: "Token " + token,
	}
	switch command {
	case "create_product":
		input := bufio.NewReader(os.Stdin)
		name, _ := input.ReadString('\n')
		description, _ := input.ReadString('\n')
		commaSeparatedTags, _ := input.ReadString('\n')
		productType, _ := input.ReadString('\n')
		productSla, _ := input.ReadString('\n')
		productData := ddclient.Product{
			Name:             name,
			Description:      description,
			Tags:             strings.Split(commaSeparatedTags, ","),
			ProdType:         productType,
			SlaConfiguration: productSla,
		}
		result, err := client.CreateProduct(productData)
		if err != nil {
			log.Fatalf("failed to create product: %v", err)
		}
		log.Printf("successfully created product: %s", result)
	case "find_product":
		input := bufio.NewReader(os.Stdin)
		productToFindName, _ := input.ReadString('\n')
		result, err := client.FindProduct(productToFindName)
		if err != nil {
			log.Fatalf("failed to find product: %v", err)
		}
		log.Printf("successfully found product: %d", result)
	}
}
