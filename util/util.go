package util

func GetMapKeys(m map[string]string) (keys []string) {
	keys = make([]string, len(m))

	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	return
}