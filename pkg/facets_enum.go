// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package godrinth

import (
	"fmt"
	"strings"
)

const (
	// OperandTypeEqual is a OperandType of type Equal.
	OperandTypeEqual OperandType = iota
	// OperandTypeNotEqual is a OperandType of type NotEqual.
	OperandTypeNotEqual
	// OperandTypeGreaterThan is a OperandType of type GreaterThan.
	OperandTypeGreaterThan
	// OperandTypeGreaterThanOrEqual is a OperandType of type GreaterThanOrEqual.
	OperandTypeGreaterThanOrEqual
	// OperandTypeLessThan is a OperandType of type LessThan.
	OperandTypeLessThan
	// OperandTypeLessThanOrEqual is a OperandType of type LessThanOrEqual.
	OperandTypeLessThanOrEqual
)

var ErrInvalidOperandType = fmt.Errorf("not a valid OperandType, try [%s]", strings.Join(_OperandTypeNames, ", "))

const _OperandTypeName = "EqualNotEqualGreaterThanGreaterThanOrEqualLessThanLessThanOrEqual"

var _OperandTypeNames = []string{
	_OperandTypeName[0:5],
	_OperandTypeName[5:13],
	_OperandTypeName[13:24],
	_OperandTypeName[24:42],
	_OperandTypeName[42:50],
	_OperandTypeName[50:65],
}

// OperandTypeNames returns a list of possible string values of OperandType.
func OperandTypeNames() []string {
	tmp := make([]string, len(_OperandTypeNames))
	copy(tmp, _OperandTypeNames)
	return tmp
}

var _OperandTypeMap = map[OperandType]string{
	OperandTypeEqual:              _OperandTypeName[0:5],
	OperandTypeNotEqual:           _OperandTypeName[5:13],
	OperandTypeGreaterThan:        _OperandTypeName[13:24],
	OperandTypeGreaterThanOrEqual: _OperandTypeName[24:42],
	OperandTypeLessThan:           _OperandTypeName[42:50],
	OperandTypeLessThanOrEqual:    _OperandTypeName[50:65],
}

// String implements the Stringer interface.
func (x OperandType) String() string {
	if str, ok := _OperandTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("OperandType(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x OperandType) IsValid() bool {
	_, ok := _OperandTypeMap[x]
	return ok
}

var _OperandTypeValue = map[string]OperandType{
	_OperandTypeName[0:5]:   OperandTypeEqual,
	_OperandTypeName[5:13]:  OperandTypeNotEqual,
	_OperandTypeName[13:24]: OperandTypeGreaterThan,
	_OperandTypeName[24:42]: OperandTypeGreaterThanOrEqual,
	_OperandTypeName[42:50]: OperandTypeLessThan,
	_OperandTypeName[50:65]: OperandTypeLessThanOrEqual,
}

// ParseOperandType attempts to convert a string to a OperandType.
func ParseOperandType(name string) (OperandType, error) {
	if x, ok := _OperandTypeValue[name]; ok {
		return x, nil
	}
	return OperandType(0), fmt.Errorf("%s is %w", name, ErrInvalidOperandType)
}

func (x OperandType) Ptr() *OperandType {
	return &x
}

// MarshalText implements the text marshaller method.
func (x OperandType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *OperandType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseOperandType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
