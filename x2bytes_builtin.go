package x2bytes

// Collection of conversion functions from builtin types, like int, float, ...

import (
	"strconv"
)

// BytesToBytes converts from byte array.
func BytesToBytes(dst []byte, val any, _ ...any) ([]byte, error) {
	switch x := val.(type) {
	case *[]byte:
		dst = append(dst, *x...)
	case []byte:
		dst = append(dst, x...)
	default:
		return dst, ErrUnknownType
	}

	return dst, nil
}

// StrToBytes converts from string.
func StrToBytes(dst []byte, val any, _ ...any) ([]byte, error) {
	switch x := val.(type) {
	case *string:
		dst = append(dst, *x...)
	case string:
		dst = append(dst, x...)
	default:
		return dst, ErrUnknownType
	}

	return dst, nil
}

// BoolToBytes converts from boolean.
func BoolToBytes(dst []byte, val any, _ ...any) ([]byte, error) {
	var b bool
	switch x := val.(type) {
	case *bool:
		b = *x
	case bool:
		b = x
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
func IntToBytes(dst []byte, val any, _ ...any) ([]byte, error) {
	var i int64
	switch x := val.(type) {
	case int:
		i = int64(x)
	case *int:
		i = int64(*x)
	case int8:
		i = int64(x)
	case *int8:
		i = int64(*x)
	case int16:
		i = int64(x)
	case *int16:
		i = int64(*x)
	case int32:
		i = int64(x)
	case *int32:
		i = int64(*x)
	case int64:
		i = x
	case *int64:
		i = *x
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendInt(dst, i, 10)
	return dst, nil
}

// UintToBytes converts from uint (including uint8, uint16, ...).
func UintToBytes(dst []byte, val any, _ ...any) ([]byte, error) {
	var i uint64
	switch x := val.(type) {
	case uint:
		i = uint64(x)
	case *uint:
		i = uint64(*x)
	case uint8:
		i = uint64(x)
	case *uint8:
		i = uint64(*x)
	case uint16:
		i = uint64(x)
	case *uint16:
		i = uint64(*x)
	case uint32:
		i = uint64(x)
	case *uint32:
		i = uint64(*x)
	case uint64:
		i = x
	case *uint64:
		i = *x
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendUint(dst, i, 10)
	return dst, nil
}

// FloatToBytes converts from float (32 and 64 bits).
func FloatToBytes(dst []byte, val any, _ ...any) ([]byte, error) {
	var f float64
	switch x := val.(type) {
	case float32:
		f = float64(x)
	case *float32:
		f = float64(*x)
	case float64:
		f = x
	case *float64:
		f = *x
	default:
		return dst, ErrUnknownType
	}

	dst = strconv.AppendFloat(dst, f, 'f', -1, 64)
	return dst, nil
}
