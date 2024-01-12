package main

import (
	"bufio"
	"defect-dojo-tools/pkg/ddclient"
	"log"
	"os"
	"strings"
)

// 1st cmd argument - command
// 2nd cmd argument - API URI
// 3rd cmd argument - API Token
func main() {
	if len(os.Args) != 4 {
		log.Fatalf("incorrect arguments quantity: %d", len(os.Args)-1)
	}
	command, uri, token := os.Args[1], os.Args[2], os.Args[3]
	client := ddclient.Client{
		ApiURL:   uri,
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
	}
}
