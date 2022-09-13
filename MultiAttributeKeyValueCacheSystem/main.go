package main

import (
	"credMC/storage"
	"fmt"
)

func main() {
	stg := storage.NewStorage()
	if err := stg.Add("delhi", map[string]interface{}{
		"pollution-level": "very high",
		"population":      "10 million",
	}); err != nil {
		fmt.Println(err)
	}
	if err := stg.Add("jakarta", map[string]interface{}{
		"pollution-level": "high",
		"latitude":        -6.0,
		"longitude":       106.0,
	}); err != nil {
		fmt.Println(err)
	}
	if err := stg.Add("bangalore", map[string]interface{}{
		"pollution-level": "moderate",
		"latitude":        12.94,
		"longitude":       77.64,
		"free_food":       "true",
	}); err != nil {
		fmt.Println(err)
	}
	if err := stg.Add("india", map[string]interface{}{
		"capital":    "delhi",
		"population": "1.2 billion",
	}); err != nil {
		fmt.Println(err)
	}
	if err := stg.Add("crocin", map[string]interface{}{
		"manufacturer": "GSK",
	}); err != nil {
		fmt.Println(err)
	}

	res, err := stg.GetValueByKey("delhi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	keyRes, err := stg.GetKeyByAttributeValue("pollution-level", "high")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(keyRes)
	}
	keyRes, err = stg.GetKeyByAttributeValue("manufacturer", "GSK")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(keyRes)
	}
	stg.DeleteKey("delhi")
	res, err = stg.GetValueByKey("delhi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
