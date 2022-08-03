package grumble

func interfaceSliceConvertToGenericSlice[T any](s []interface{}) (res []T) {
	res = make([]T, 0, len(s))
	for i := range s {
		t, ok := s[i].(T)
		if !ok {
			continue
		}
		res = append(res, t)
	}
	return
}
