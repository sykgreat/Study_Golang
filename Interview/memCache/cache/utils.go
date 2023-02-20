package cache

import (
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatalln(err)
	}
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, err := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	unit = strings.ToUpper(unit)

	var byteNum int64 = 0

	switch unit {
	case "B":
		byteNum = num * B
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		log.Fatalln("ParseSize error")
		num = 100
		byteNum = num * MB
		unit = "MB"
	}

	sizeStr := strconv.FormatInt(num, 10) + unit

	return byteNum, sizeStr
}

func GetValueSize(value interface{}) int64 {
	var size int64 = 0
	vo := reflect.ValueOf(value)
	switch vo.Kind() {
	case reflect.Map:
		for _, key := range vo.MapKeys() {
			size += GetValueSize(key.Interface())
			size += GetValueSize(vo.MapIndex(key).Interface())
		}
	case reflect.Slice:
		for i := 0; i < vo.Len(); i++ {
			size += GetValueSize(vo.Index(i).Interface())
		}
	case reflect.Struct:
		for i := 0; i < vo.NumField(); i++ {
			size += GetValueSize(vo.Field(i).Interface())
		}
	}
	size += int64(vo.Type().Size())
	return size
}
