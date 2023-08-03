package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ExternalContext struct {
	Source string `json:"source"`
}

type Data struct {
	ExternalContext ExternalContext `json:"external_context"`
}

type Attachment struct {
	Data []Data `json:"data"`
}

type SaveV2 struct {
	Attachments []Attachment `json:"attachments"`
}

type JSONFile struct {
	SavesV2 []SaveV2 `json:"saves_v2"`
}

func main() {
	jsonFile, err := os.Open("your_saved_items.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result JSONFile
	json.Unmarshal(byteValue, &result)

	for i := range result.SavesV2 {
		for j := range result.SavesV2[i].Attachments {
			for k := range result.SavesV2[i].Attachments[j].Data {
				fmt.Println(result.SavesV2[i].Attachments[j].Data[k].ExternalContext.Source)
			}
		}
	}
}