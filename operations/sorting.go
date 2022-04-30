package operations

import (
	"GoAssignment-1/model"
	"sort"
)
//Function to sort the list
func Sorting(list []model.FinalList) {
	sort.Slice(list, func(i, j int) bool {
		if list[i].Count > list[j].Count {
			return true
		}
		if list[i].Count < list[j].Count {
			return false
		}
		return list[i].MaxPower > list[j].MaxPower
	})
}