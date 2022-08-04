package grumble

func interfaceSliceConvertToGenericSlice[T any](s []any) (res []T) {
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

func interfaceConvertTpGenericSlice[T any](i any) ([]T, bool) {
	s, ok := i.([]T)
	if ok {
		return s, true
	}

	s2, ok := i.([]any)
	if !ok {
		return nil, false
	}

	return interfaceSliceConvertToGenericSlice[T](s2), true
}