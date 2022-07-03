package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here

	var keys []int

	for i := range input {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, v := range keys {
		result = append(result, input[v])
	}

	return
}
