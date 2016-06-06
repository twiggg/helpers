package fornums

func RemoveFromSlice(slice *[]int32, i int32) bool {
	n := len(*slice)
	if !removableFromSlice(*slice, i) {
		return false
	}
	s := *slice
	s = append(s[:i], s[i+1:]...) //
	*slice = s
	if len(*slice) == n-1 {
		return true
	} else {
		return false
	}
}
func removableFromSlice(slice []int32, i int32) bool {
	if i < 0 || i > int32(len(slice)-1) {
		return false
	}
	return true
}

func FindInInt32Slice(slice []int32, toFind int32) (int, bool) {
	for k, v := range slice {
		if v == toFind {
			return k, true
		}
	}
	return -1, false
}
