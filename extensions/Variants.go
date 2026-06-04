package extensions

import "sort"

func Variants() []string {

	result := make([]string, 0)

	for name, _ := range UserAgents {
		result = append(result, name)
	}

	sort.Strings(result)

	return result

}
