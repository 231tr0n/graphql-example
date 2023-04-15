package database

var keyStore map[string]string

func init() {
	keyStore = make(map[string]string)
}

func Store(key string, value string) [2]string {
	keyStore[key] = value
	return [2]string{key, value}
}

func Delete(key string) [2]string {
	var value, ok = keyStore[key]
	delete(keyStore, key)
	if ok {
		return [2]string{key, value}
	}
	return [2]string{"", ""}
}

func Clear() [][2]string {
	var entries [][2]string
	for key, value := range keyStore {
		entries = append(entries, [2]string{key, value})
		delete(keyStore, key)
	}
	return entries
}

func Get(key string) (string, bool) {
	var value, ok = keyStore[key]
	return value, ok
}

func List() [][2]string {
	var entries [][2]string
	for key, value := range keyStore {
		entries = append(entries, [2]string{key, value})
	}
	return entries
}
