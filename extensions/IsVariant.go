package extensions

func IsVariant(name string) bool {

	result := false

	_, ok := UserAgents[name]

	if ok == true {
		result = true
	}

	return result

}
