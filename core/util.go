package core

func ContainsStr(coll []string, key string) bool {
	for _, s := range coll {
		if s == key {
			return true
		}
	}
	return false
}
