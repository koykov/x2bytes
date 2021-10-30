package x2bytes

import "errors"

// ToBytesFn is a signature of conversion function.
type ToBytesFn func(dst []byte, val interface{}) ([]byte, error)

var (
	// Registry of conversion functions.
	toBytesFnRegistry = make([]ToBytesFn, 0)
	ErrUnknownType    = errors.New("unknown type")

	_ = ToBytesWR
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

// RegisterToBytesFn registers new conversion function.
func RegisterToBytesFn(fn ToBytesFn) {
	for _, f := range toBytesFnRegistry {
		if &f == &fn {
			return
		}
	}
	toBytesFnRegistry = append(toBytesFnRegistry, fn)
}

// ToBytes is a generic conversion function.
//
// Convert val to byte array and append result to the dst.
// Returns dst and conversion error message. Error is nil when succeeded.
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

// ToBytesWR converts val to byte array with preliminary reset length of the dst.
func ToBytesWR(dst []byte, val interface{}) ([]byte, error) {
	dst = dst[:0]
	return ToBytes(dst, val)
}
