package arpc

func ParsePageSize(pageRaw, sizeRaw any, args ...any) (int, int) {
	page := toInt(pageRaw)
	size := toInt(sizeRaw)

	if page < 1 {
		page = 1
	}

	pageSize := 20 // 每页大小
	if len(args) > 0 {
		// 每页大小
		customMaxSize := toInt(args[0])
		if customMaxSize > 0 {
			pageSize = customMaxSize
		}
	}

	if size < 1 || size > pageSize {
		size = pageSize
	}

	offset := (page - 1) * size
	return offset, size
}

func toInt(val any) int {
	switch v := val.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}
