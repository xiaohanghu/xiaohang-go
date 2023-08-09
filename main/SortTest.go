package main

import (
	"fmt"
	"sort"
	"encoding/json"
)

// type Person struct 

func main()  {
	strs := []string{"United States", "India", "France", "United Kingdom", "Spain"}
	
	sort.Slice(strs, func(i, j int) bool {
		// return len(strs[i]) < len(strs[j])
		return strs[i] < strs[j]
	})

	jsonText, _ := json.Marshal(strs)
	fmt.Printf("jsonText type: %T \n Sorted: %v \n", jsonText, string(jsonText))

}