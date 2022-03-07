package util

import (
	"bytes"
	"encoding/binary"
)

func ConcatString(inString ... string)string{
	bufferStr := bytes.Buffer{}
	for _,value := range inString{
		bufferStr.WriteString(value)
	}
	return bufferStr.String()
}

// encode string array 
func array2string(array []string) string {
	var buffer bytes.Buffer
	for _, str := range array {
		length := make([]byte, 4)
		binary.BigEndian.PutUint32(length, uint32(len(str)))
		buffer.Write(length)
		buffer.Write([]byte(str))
	}
	return buffer.String()
}

// decode string array 
func string2Array(str string) ([]string, error) {
	if len(str) <= 4 {
		return []string{str}, nil
	}
	strBytes := []byte(str)
	var array []string
	pivot := 0
	for {
		length := binary.BigEndian.Uint32(strBytes[pivot : pivot+4])
		pivot = pivot + 4
		if length > uint32(len(strBytes[pivot:])) {
			return []string{}, errors.New("could not decode the string " + str + " to string array")
		}
		array = append(array, string(strBytes[pivot:pivot+int(length)]))
		pivot = pivot + int(length)
		if len(strBytes[pivot:]) == 0 {
			return array, nil
		}
		if len(strBytes[pivot:]) < 4 {
			return []string{}, errors.New("could not decode the string " + str + " to string array.wrong encode format")
		}
	}
}
