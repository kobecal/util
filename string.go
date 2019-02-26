package util

import "bytes"

func ConcatString(inString ... string)string{
	bufferStr := bytes.Buffer{}
	for _,value := range inString{
		bufferStr.WriteString(value)
	}
	return bufferStr.String()
}
