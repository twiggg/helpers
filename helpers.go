package helpers

/*func (f *UserFields) remove(i int) {
	if !f.removable(i) {
		return
	}
	s := *f
	s = append(s[:i], s[i+1:]...) //
	*f = s
}
func (f UserFields) removable(i int) bool {
	if i < 0 || i > (len(f)-1) {
		return false
	}
	return true
}*/

func RemovedFromStringSlice(slice *[]string, i int) bool {
	n := len(*slice)
	if !removableFromStringSlice(*slice, i) {
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
func RemovableFromStringSlice(slice []string, i int) bool {
	if i < 0 || i > (len(slice)-1) {
		return false
	}
	return true
}
