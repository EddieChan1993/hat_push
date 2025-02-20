package util

//MapMergeSum map合并，值累加
func MapMergeSum(base, add map[int32]int64) map[int32]int64 {
	if len(base) == 0 {
		return add
	}
	for k, v := range add {
		if old, had := base[k]; had {
			base[k] = old + v
		} else {
			base[k] = v
		}
	}
	return base
}

//Map32MergeSum map合并，值累加
func Map32MergeSum(base, add map[int32]int32) map[int32]int32 {
	if len(base) == 0 {
		return add
	}
	for k, v := range add {
		if old, had := base[k]; had {
			base[k] = old + v
		} else {
			base[k] = v
		}
	}
	return base
}

//CopyMapInt64 map复制
func CopyMapInt64(from map[int32]int64) map[int32]int64 {
	res := make(map[int32]int64, len(from))
	for k, v := range from {
		res[k] = v
	}
	return res
}

//CopyMapInt32 map复制
func CopyMapInt32(from map[int32]int32) map[int32]int32 {
	res := make(map[int32]int32, len(from))
	for k, v := range from {
		res[k] = v
	}
	return res
}
