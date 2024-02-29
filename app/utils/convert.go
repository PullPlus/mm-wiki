package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

var Convert = NewConvert()

type convert struct{}

func NewConvert() *convert {
	return &convert{}
}

// bool <LABEL_846>
func (c *convert) BoolToString(boolValue bool) string {
	if boolValue == true {
		return "true"
	} else {
		return "false"
	}
}

//bool <LABEL_1512> int
func (c *convert) BoolToInt(boolValue bool) int {
	if boolValue == true {
		return 1
	} else {
		return 0
	}
}

//int <LABEL_1512> bool
func (c *convert) IntToBool(number int) bool {
	if number == 0 {
		return false
	} else {
		return true
	}
}

//int <LABEL_846>
//base <LABEL_1645> 2-32 <LABEL_1646>
func (c *convert) IntToString(number int64, base int) string {
	return strconv.FormatInt(number, base)
}

//string to int(10<LABEL_1646>)
func (c *convert) StringToInt(str string) int {
	intValue, _ := strconv.Atoi(str)
	return intValue
}

// string to int64(10<LABEL_1646>)
func (c *convert) StringToInt64(str string) int64 {
	intValue, _ := strconv.ParseInt(str, 10, 64)
	return intValue
}

// int <LABEL_1512>10<LABEL_990> IntToString(number, 10)
func (c *convert) IntToTenString(number int) string {
	return strconv.Itoa(number)
}

// float <LABEL_846>
func (c *convert) FloatToString(f float64, fmt byte, prec, bitSize int) string {
	return strconv.FormatFloat(f, fmt, prec, bitSize)
}

// <LABEL_595> int64
func (c *convert) ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
