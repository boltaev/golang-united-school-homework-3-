package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	//saving keys of map into slice
	keys := make([]int, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}

	// sorting keys
	sort.Ints(keys)

	//saving values of map into slice
	for _, k := range keys {
		result = append(result, input[k])
	}

	return result
}
