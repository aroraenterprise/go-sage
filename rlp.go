package main

import "bytes"

func EncodeSlice(slice []interface{}) []byte {
	var buff bytes.Buffer

	for _, val := range slice {
		switch t := val.(type) {
		case []interface{}, []string:
			buff.Write(Encode(t))
		}
	}
	return buff.Bytes()
}

func Encode(object interface{}) []byte {
	var buff bytes.Buffer

	switch t := object.(type) {
	case string:
		if len(t) < 56 {

		}
	case uint32, uint64:
		var num uint64
		if _num, ok := t.(uint64); ok {
			num = _num
		} else if _num, ok := t.(uint32); ok {
			num = uint64(_num)
		}

		if num >= 0 && num < 24 {
			buff.WriteString(string(rune(num)))
		}
	case []byte:
		buff.Write(Encode(string(t)))
	case []interface{}:
		buff.Write(EncodeSlice(t))
	}
	return buff.Bytes()
}
