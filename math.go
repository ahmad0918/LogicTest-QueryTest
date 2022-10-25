package main

import (
	"fmt"
)


func main()  {
	var array = []string{"cook", "save", "taste", "aves", "vase", "state", "map"}	
	addSlice(array)
}

func addSlice(array []string) {
	arr :=[][]string{}
	for i:= 0; i < len(array); i++ {
		ar := []string{}
		for _ , n := range array {
			valid := isAnagram(array[i],n)
			if valid {
				ar = append(ar, n)
			}
		}
		arr = append(arr, ar)
	}
	arr = append(arr[:3], arr[3+1:]...)
	arr = append(arr[:3], arr[3+1:]...)
	arr = append(arr[:3], arr[3+1:]...)
	fmt.Println(arr)
}

func isAnagram(s string, t string) bool {
	if(len(s) != len(t)) { return false }
    count := make([]int, 26)
    
    for index, _ := range count {
        count[index] = 0
    }
    
    for i:= 0; i < len(s); i++ {
        count[int(s[i]) - int('a')]++
        count[int(t[i]) - int('a')]--
    }
    
    for _, val := range count {
        if val != 0 {
            return false
        }
    }
    return true
}
