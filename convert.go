package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type ResourceRecord struct {
	Value string `json:"Value"`
}

type AliasTarget struct {
	HostedZoneId         string `json:"HostedZoneId"`
	DNSName              string `json:"DNSName"`
	EvaluateTargetHealth bool   `json:"EvaluateTargetHealth"`
}

type ResourceRecordSet struct {
	Name            string           `json:"Name"`
	Type            string           `json:"Type"`
	TTL             int              `json:"TTL,omitempty"`
	ResourceRecords []ResourceRecord `json:"ResourceRecords,omitempty"`
	AliasTarget     *AliasTarget     `json:"AliasTarget,omitempty"`
}

type BeforeFormat struct {
	ResourceRecordSets []ResourceRecordSet `json:"ResourceRecordSets"`
}

type Change struct {
	Action            string            `json:"Action"`
	ResourceRecordSet ResourceRecordSet `json:"ResourceRecordSet"`
}

type AfterFormat struct {
	Comment string   `json:"Comment"`
	Changes []Change `json:"Changes"`
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <input-json-name> <domain-name> \n", os.Args[0])
	}

	inputFileName := os.Args[1]
	domainName := os.Args[2]
	dotDomainName := domainName + "."

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("Failed to open input file: %v\n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Failed to close input file: %v\n", err)
		}
	}(file)

	// read file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read input file: %v\n", err)
	}

	var before BeforeFormat
	err = json.Unmarshal(fileContent, &before)
	if err != nil {
		log.Fatalf("Failed to unmarshal input JSON: %v\n", err)
	}

	var after AfterFormat
	after.Comment = domainName

	for _, recordSet := range before.ResourceRecordSets {
		// if record name does not end with domain name, exit
		if !strings.HasSuffix(recordSet.Name, dotDomainName) {
			log.Fatalf("Invalid record name: %s\n", recordSet.Name)
		}

		if recordSet.Name == dotDomainName && (recordSet.Type == "NS" || recordSet.Type == "SOA") {
			continue
		}

		change := Change{
			Action:            "CREATE",
			ResourceRecordSet: recordSet,
		}

		after.Changes = append(after.Changes, change)
	}

	outputJSON, err := json.MarshalIndent(after, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal output JSON: %v\n", err)
	}

	fmt.Println(string(outputJSON))
}
