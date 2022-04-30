package operations

import (
	"GoAssignment-1/model"
)

//Function to remove extra characters from the list
func RemoveCharacters(list []model.FinalList) []model.FinalList {
	if len(list) > 15 {
		for len(list) != 15 {
			list = list[:len(list)-1]
		}
	}
	return list
}
