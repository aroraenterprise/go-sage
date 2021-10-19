package main

import (
	"bytes"
	"math"
)

func ToBinary(x int, bytes int) string {
	if bytes == 0 {
		return ""
	}
	return ToBinary(int(x/256), bytes-1) + string(rune(x%256))
}

// NumToVarInt returns string
func NumToVarInt(x int) string {
	if x < 253 {
		return string(rune(x))
	}
	if x < int(math.Pow(2, 16)) {
		return string(rune(253)) + ToBinary(x, 2)
	}
	if x < int(math.Pow(2, 32)) {
		return string(rune(253)) + ToBinary(x, 4)
	}
	return string(rune(253)) + ToBinary(x, 8)
}

// RlpEncode returns string
func RlpEncode(object interface{}) string {
	if str, ok := object.(string); ok {
		return "\x00" + NumToVarInt(len(str)) + str
	} else if slice, ok := object.([]interface{}); ok {
		var buffer bytes.Buffer
		for _, val := range slice {
			if v, ok := val.(string); ok {
				buffer.WriteString(RlpEncode(v))
			} else {
				buffer.WriteString(RlpEncode(val))
			}
		}
		return "\x01" + RlpEncode(len(buffer.String())) + buffer.String()
	} else if slice, ok := object.([]string); ok {
		var buffer bytes.Buffer
		for _, val := range slice {
			buffer.WriteString(RlpEncode(val))
		}
		return "\x01" + RlpEncode(len(buffer.String())) + buffer.String()
	}

	return ""
}
