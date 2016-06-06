package forstrings

import (
	"errors"
	"fmt"
	"html"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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

func RemoveFromSlice(slice *[]string, i int) bool {
	n := len(*slice)
	if !RemovableFromSlice(*slice, i) {
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
func RemovableFromSlice(slice []string, i int) bool {
	if i < 0 || i > (len(slice)-1) {
		return false
	}
	return true
}

func FindInSlice(slice []string, toFind string) (int, bool) {
	for k, v := range slice {
		if strings.ToLower(html.EscapeString(v)) == strings.ToLower(html.EscapeString(toFind)) {
			return k, true
		}
	}
	return -1, false
}

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func RandomIntString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func StringsToInt32s(list []string) ([]int32, []error) {
	res := []int32{}
	errs := []error{}
	for k, v := range list {
		v = strings.ToLower(html.EscapeString(strings.TrimSpace(v)))
		r, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			if len(errs) == 0 {
				msg0 := fmt.Sprintf("emitter: forstrings.StringsToInt32s")
				errs = append(errs, errors.New(msg0))
			}
			msg := fmt.Sprintf("index: %d, error: %s", k, err.Error())
			errs = append(errs, errors.New(msg))
		} else {
			res = append(res, int32(r))
		}
	}
	return res, errs
}

func Concatenate(list []string) string {
	res := ""
	//l := len(list)
	for k, v := range list {
		if k != 0 {
			res = fmt.Sprintf("%s, %s", res, v)
		} else {
			res = fmt.Sprintf("%s", v)
		}
	}
	return res
}

func ConcatenateErrs(list []error) string {
	res := ""
	//l := len(list)
	for k, v := range list {
		if k != 0 {
			res = fmt.Sprintf("%s, %s", res, v.Error())
		} else {
			res = fmt.Sprintf("%s", v.Error())
		}
	}
	return res
}
