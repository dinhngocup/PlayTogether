package redis

func GenPrefixKey(object string, id string, field string) string {
	if field == "" {
		return object + ":" + id
	}
	return object + ":" + id + ":" + field
}
