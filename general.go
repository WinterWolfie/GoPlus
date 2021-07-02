package GoPlus

import (
	"log"
	"fmt"
)

func HandleErr(err error)  {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func HandleLightErr(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveFromSliceOrdered(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func RemoveFromSliceUnordered(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

