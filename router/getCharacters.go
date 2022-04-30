package router

import (
	"GoAssignment-1/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var List []model.FinalList
var flag = 0
//Function to get all the characters info from the urls
func GetCharacters() []model.FinalList{
	if flag == 0 {
		urls := []string{
			"http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b",
			"http://www.mocky.io/v2/5ecfd630320000f1aee3d64d",
			"http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e",
		}
		for i := 0; i < len(urls); i++ {
			response, err := http.Get(urls[i])
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			var responseObject model.Marvel
			json.Unmarshal(responseData, &responseObject)

			for i := 0; i < len(responseObject.Character); i++ {
				results := []model.FinalList{{Name: strings.ToLower(responseObject.Character[i].Name), MaxPower: responseObject.Character[i].MaxPower}}
				List = append(List, results...)
			}
		}
		flag = 1
	}
	return List
}
