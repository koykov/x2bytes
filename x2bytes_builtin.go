package x2bytes

// Collection of conversion functions from builtin types, like int, float, ...

import (
	"strconv"
	"time"
)

// BytesToBytes converts from byte array.
func BytesToBytes(dst []byte, val any) ([]byte, error) {
	switch val.(type) {
	case *[]byte:
		dst = append(dst, *val.(*[]byte)...)
	case []byte:
		dst = append(dst, val.([]byte)...)
	default:
		return dst, ErrUnknownType
	}

	return dst, nil
}

// StrToBytes converts from string.
func StrToBytes(dst []byte, val any) ([]byte, error) {
	switch val.(type) {
	case *string:
		dst = append(dst, *val.(*string)...)
	case string:
		dst = append(dst, val.(string)...)
	default:
		return dst, ErrUnknownType
	}

	return dst, nil
}

// BoolToBytes converts from boolean.
func BoolToBytes(dst []byte, val any) ([]byte, error) {
	var b bool
	switch val.(type) {
	case *bool:
		b = *val.(*bool)
	case bool:
		b = val.(bool)
	default:
		return dst, ErrUnknownType
	}

	if b {
		dst = append(dst, "true"...)
	} else {
		dst = append(dst, "false"...)
	}

	return dst, nil
}

// IntToBytes converts from int (including int8, int16, ...).
func IntToBytes(dst []byte, val any) ([]byte, error) {
	var i int64
	switch val.(type) {
	case int:
		i = int64(val.(int))
	case *int:
		i = int64(*val.(*int))
	case int8:
		i = int64(val.(int8))
	case *int8:
		i = int64(*val.(*int8))
	case int16:
		i = int64(val.(int16))
	case *int16:
		i = int64(*val.(*int16))
	case int32:
		i = int64(val.(int32))
	case *int32:
		i = int64(*val.(*int32))
	case int64:
		i = val.(int64)
	case *int64:
		i = *val.(*int64)
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendInt(dst, i, 10)
	return dst, nil
}

// UintToBytes converts from uint (including uint8, uint16, ...).
func UintToBytes(dst []byte, val any) ([]byte, error) {
	var i uint64
	switch val.(type) {
	case uint:
		i = uint64(val.(uint))
	case *uint:
		i = uint64(*val.(*uint))
	case uint8:
		i = uint64(val.(uint8))
	case *uint8:
		i = uint64(*val.(*uint8))
	case uint16:
		i = uint64(val.(uint16))
	case *uint16:
		i = uint64(*val.(*uint16))
	case uint32:
		i = uint64(val.(uint32))
	case *uint32:
		i = uint64(*val.(*uint32))
	case uint64:
		i = val.(uint64)
	case *uint64:
		i = *val.(*uint64)
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendUint(dst, i, 10)
	return dst, nil
}

// FloatToBytes converts from float (32 and 64 bits).
func FloatToBytes(dst []byte, val any) ([]byte, error) {
	var f float64
	switch val.(type) {
	case float32:
		f = float64(val.(float32))
	case *float32:
		f = float64(*val.(*float32))
	case float64:
		f = val.(float64)
	case *float64:
		f = *val.(*float64)
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendFloat(dst, f, 'f', -1, 64)
	return dst, nil
}

// TimeToBytes converts from time.Time.
func TimeToBytes(dst []byte, val any) ([]byte, error) {
	var t time.Time
	switch val.(type) {
	case time.Time:
		t = val.(time.Time)
	case *time.Time:
		t = *val.(*time.Time)
	default:
		return dst, ErrUnknownType
	}
	dst = t.AppendFormat(dst, time.RFC3339) // todo find a way to pass variadic layouts
	return dst, nil
}
