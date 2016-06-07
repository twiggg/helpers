package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	d,err:=ExtractFromString("16/30/03", "YY/DD/MM")
	fmt.Println(*d,err)
	fmt.Println(d.GenTime())
}

//custom types declaration
type date struct {
	day   int
	month int
	year  int
}
type dateString map[string]string
type day string
type month string
type year string


var months =map[int]string{1:"Jan",2:"Feb",3:"Mar",4:"Apr",5:"May",6:"Jun",7:"Jul",8:"Aug",9:"Sep",10:"Oct",11:"Nov",12:"Dec"}


//methods
func (d date) GenTime() time.Time{
	const shortForm = "02-Jan-2006"
	phrase:=fmt.Sprintf("%d-%s-%d",d.day,months[d.month],d.year)
	fmt.Println(phrase)
    t, _ := time.Parse(shortForm, phrase)
    //t, _ := time.Parse(shortForm, "13-Mar-2016")
    //fmt.Println(t)
    return t
}

func (ds dateString) isValid() (*date, error) {
	d, dok := isDayOk(ds["day"])
	m, mok := isMonthOk(ds["month"])
	y, yok := isYearOk(ds["year"])
	if dok == nil && mok == nil && yok == nil {
		return &date{day: d, month: m, year: y}, nil
	}
	msg := ""
	errs:=[]error{dok, mok, yok}
	for i, m := range errs {
		if m != nil {
			if i > 0 {
				msg = fmt.Sprintf("%s%s", msg, ",")
			}
			msg = fmt.Sprintf("%s%s", msg, m)
		}
	}
	return nil, errors.New(msg)
}

//dispatch components of the date in the date map
func (dm dateString) dispatch(dpos, mpos, ypos int, list []string) error {
	if dpos >= 0 && dpos < len(list) && mpos >= 0 && mpos < len(list) && ypos >= 0 && ypos < len(list) && dpos != mpos && dpos != ypos && mpos != ypos {
		dm["day"] = list[dpos]
		dm["month"] = list[mpos]
		dm["year"] = list[ypos]
		return nil
	}
	return errors.New("please make sure to provide positions with 0<indices<len(list) that are not overlapping.")
}

//extract the components of a date from a string
func ExtractFromString(date string, format string) (*date, error) {
	//dmap:=make(map[string]int,3)
	parts := strings.Split(date, "/")
	dm := dateString{}
	if len(parts) != 3 {
		return nil, errors.New("a valid date is made of 3 components:day,month,year")
	}
	switch strings.ToUpper(format) {
	case "DD/MM/YY":
		err := dm.dispatch(0, 1, 2, parts)
		if err != nil {
			return nil, err
		}
		return dm.isValid()
	case "JJ/MM/AA":
			err := dm.dispatch(0, 1, 2, parts)
			if err != nil {
				return nil, err
			}
			return dm.isValid()
	case "DD/MM/YYYY" :
		err := dm.dispatch(0, 1, 2, parts)
		if err != nil {
			return nil, err
		}
		return dm.isValid()
		case  "JJ/MM/AAAA":
			err := dm.dispatch(0, 1, 2, parts)
			if err != nil {
				return nil, err
			}
			return dm.isValid()
	case "MM/DD/YY" :
		err := dm.dispatch(1, 0, 2, parts)
		if err != nil {
			return nil, err
		}
		return dm.isValid()
		case  "MM/JJ/AA":
			err := dm.dispatch(1, 0, 2, parts)
			if err != nil {
				return nil, err
			}
			return dm.isValid()

	case "MM/DD/YYYY" :
		err := dm.dispatch(1, 0, 2, parts)
		if err != nil {
			return nil, err
		}
		return dm.isValid()
	case  "MM/JJ/AAAA":
			err := dm.dispatch(1, 0, 2, parts)
			if err != nil {
				return nil, err
			}
			return dm.isValid()
	case "YY/DD/MM" :
		err := dm.dispatch(1, 2, 0, parts)
		if err != nil {
			return nil, err
		}
		return dm.isValid()
		case  "AA/JJ/MM":
			err := dm.dispatch(1, 2, 0, parts)
			if err != nil {
				return nil, err
			}
			return dm.isValid()
	case "YYYY/DD/MM" :
		err := dm.dispatch(1, 2, 0, parts)
		if err != nil {
			return nil, err
		}
		return dm.isValid()
	case  "AAAA/JJ/MM":
			err := dm.dispatch(1, 2, 0, parts)
			if err != nil {
				return nil, err
			}
			return dm.isValid()
	case "YY/MM/DD" :
		err := dm.dispatch(2, 1, 0, parts)
		if err != nil {
			return nil, err
		}
		return dm.isValid()
	case  "AA/MM/JJ":
			err := dm.dispatch(2, 1, 0, parts)
			if err != nil {
				return nil, err
			}
			return dm.isValid()
	case "YYYY/MM/DD" :
		err := dm.dispatch(2, 1, 0, parts)
		if err != nil {
			return nil, err
		}
		return dm.isValid()
	case  "AAAA/MM/JJ":
			err := dm.dispatch(2, 1, 0, parts)
			if err != nil {
				return nil, err
			}
			return dm.isValid()
	default:
		return nil, errors.New("a valid date is made of 3 components:day,month,year")
	}
}

//components validators
func isDayOk(d string) (int, error) {
	val, err := strconv.Atoi(d)
	if err != nil {
		return -1, errors.New("could not parse day string in a day integer")
	}
	if val < 1 || val > 31 {
		return -1, errors.New("please provide a valid 1<day<31")
	}
	return val, nil
}
func isMonthOk(m string) (int, error) {
	val, err := strconv.Atoi(m)
	if err != nil {
		return -1, errors.New("could not parse month string in a month integer")
	}
	if val < 1 || val > 12 {
		return -1, errors.New("please provide a valid 1<month<12")
	}
	return val, nil
}
func isYearOk(y string) (int, error) {
	val, err := strconv.Atoi(y)
	if err != nil {
		return -1, errors.New("could not parse year string in a year integer")
	}
	switch len(y) {
	case 2:
		if val < 1 || val > (time.Now().Year()%100)+100 {
			return -1, errors.New("please provide a valid 1<year<now+100 in YY format")
		}
		//fmt.Println("now:",time.Now().Year()%2000,"val:",val)
		return val+2000, nil
	case 4:
		if val < 1 || val > time.Now().Year()+100 {
			return -1, errors.New("please provide a valid 1<year<now+100 in YYYY format")
		}
		return val, nil
	default:
		return -1, errors.New("please provide a year in YY or YYYY format")
	}
}
