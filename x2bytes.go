package x2bytes

import "errors"

// Signature of conversion function.
type ToBytesFn func(dst []byte, val interface{}) ([]byte, error)

var (
	// Registry of conversion functions.
	toBytesFnRegistry = make([]ToBytesFn, 0)

	ErrUnknownType = errors.New("unknown type")
)

func init() {
	// Register conversion functions from builtin types.
	RegisterToBytesFn(BytesToBytes)
	RegisterToBytesFn(StrToBytes)
	RegisterToBytesFn(BoolToBytes)
	RegisterToBytesFn(IntToBytes)
	RegisterToBytesFn(UintToBytes)
	RegisterToBytesFn(FloatToBytes)
}

// Register new conversion function.
func RegisterToBytesFn(fn ToBytesFn) {
	for _, f := range toBytesFnRegistry {
		if &f == &fn {
			return
		}
	}
	toBytesFnRegistry = append(toBytesFnRegistry, fn)
}

// Generic conversion function.
//
// Convert val to byte array and append result to the dst.
// Returns dst and conversion error message. Error is nil when succeed.
func ToBytes(dst []byte, val interface{}) ([]byte, error) {
	var err error
	if dst == nil {
		dst = make([]byte, 0)
	}
	for _, fn := range toBytesFnRegistry {
		dst, err = fn(dst, val)
		if err == nil {
			return dst, nil
		}
		if err != ErrUnknownType {
			return dst, err
		}
	}
	return dst, ErrUnknownType
}
