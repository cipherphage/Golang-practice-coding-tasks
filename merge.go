package main

import "fmt"

// func removeIndex(s []string, index int) []string {
// 	return append(s[:index], s[index+1:]...)
// }

// merges two string slices and dedupes
func sliceMerge(a, b []string) []string {
	var r []string
	m := make(map[string]string)

	for i := range a {
		m[a[i]] = a[i]
	}

	for i := range b {
		m[b[i]] = b[i]
	}

	for key := range m {
		r = append(r, m[key])
	}

	return r
}

func main() {
	// should print [Allan, Ernie, Oliver, Samuel]
	fmt.Println(sliceMerge(
		[]string{"Allan", "Ernie", "Oliver", "Allan"},
		[]string{"Oliver", "Allan", "Samuel", "Ernie", "Samuel"}))
}
