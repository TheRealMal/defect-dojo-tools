package main

import (
	"bufio"
	"defect-dojo-tools/pkg/ddclient"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
		var name, description, commaSeparatedTags, productType, productSla string
		multipleInput(&name, &description, &commaSeparatedTags, &productType, &productSla)
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
		fmt.Printf("%d", result)
	case "find_product":
		var productToFindName string
		multipleInput(&productToFindName)
		result, err := client.FindProduct(productToFindName)
		if err != nil {
			log.Fatalf("failed to find product: %v", err)
		}
		fmt.Printf("%d", result)
	case "find_or_create_product":
		var name, description, commaSeparatedTags, productType, productSla string
		multipleInput(&name, &description, &commaSeparatedTags, &productType, &productSla)
		productData := ddclient.Product{
			Name:             name,
			Description:      description,
			Tags:             strings.Split(commaSeparatedTags, ","),
			ProdType:         productType,
			SlaConfiguration: productSla,
		}
		result, err := client.FindProduct(name)
		if err == nil {
			fmt.Printf("%d", result)
		}
		result, err = client.CreateProduct(productData)
		if err != nil {
			log.Fatalf("failed to create product: %v", err)
		}
		fmt.Printf("%d", result)
	case "create_engagement":
		var productID, name, description, commitHash, branchTag, status string
		multipleInput(&productID, &name, &description, &commitHash, &branchTag, &status)
		today := time.Now()
		engagementData := ddclient.Engagement{
			Name:           name,
			Description:    description,
			Product:        productID,
			CommitHash:     commitHash,
			BranchTag:      branchTag,
			TargetStart:    today.Format("2006-01-02"),
			TargetEnd:      today.AddDate(0, 1, 0).Format("2006-01-02"),
			Status:         status,
			EngagementType: "CI/CD",
		}
		result, err := client.CreateEngagement(engagementData)
		if err != nil {
			log.Fatalf("failed to create engagement: %v", err)
		}
		fmt.Printf("%d", result)
	case "find_engagement":
		var productID, engagementToFindName string
		multipleInput(&productID, &engagementToFindName)
		result, err := client.FindEngagement(productID, engagementToFindName)
		if err != nil {
			log.Fatalf("failed to find engagement: %v", err)
		}
		fmt.Printf("%d", result)
	case "upload_report":
		var engagementID, format, filename, closeOldFindings string
		multipleInput(&engagementID, &format, &filename, &closeOldFindings)
		result, err := client.UploadScanReport(engagementID, format, filename, closeOldFindings)
		if err != nil {
			log.Fatalf("failed to upload report: %v", err)
		}
		fmt.Printf("%d", result)
	}
}

func multipleInput(inputVariables ...*string) {
	input := bufio.NewScanner(os.Stdin)
	for _, variable := range inputVariables {
		input.Scan()
		*variable = input.Text()
	}
}
