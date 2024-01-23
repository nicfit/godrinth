package godrinth

import (
	"fmt"
	"strings"
)

// OperandType operation between facet type and value.
// ENUM(Equal, NotEqual, GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual)
type OperandType int

func (o OperandType) OperandString() string {
	switch o {
	case OperandTypeEqual:
		return "="
	case OperandTypeNotEqual:
		return "!="
	case OperandTypeLessThan:
		return "<"
	case OperandTypeLessThanOrEqual:
		return "<="
	case OperandTypeGreaterThan:
		return ">"
	case OperandTypeGreaterThanOrEqual:
		return ">="
	default:
		panic("Unknown OperandType")

	}
}

type Facet interface {
	String() string
}

type ConcreteFacet struct {
	Type    string
	Operand OperandType
	Value   string
}

func (f ConcreteFacet) String() string {
	return fmt.Sprintf("\"%s%s%s\"", f.Type, f.Operand.OperandString(), f.Value)
}

type facetsAnd struct {
	Facets []Facet
}

func (f facetsAnd) String() string {
	strs := make([]string, 0, len(f.Facets))
	for i := range f.Facets {
		strs = append(strs, f.Facets[i].String())
	}
	return fmt.Sprintf("[[%s]]", strings.Join(strs, "],["))
}

type facetsOr struct {
	Facets []Facet
}

func (f facetsOr) String() string {
	strs := make([]string, 0, len(f.Facets))
	for i := range f.Facets {
		strs = append(strs, f.Facets[i].String())
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}

func MakeFacet(name string, operand OperandType, value string) Facet {
	return ConcreteFacet{
		Type:    name,
		Operand: operand,
		Value:   value,
	}
}

func MakeFacetAnd(facets ...Facet) Facet {
	f := facetsAnd{Facets: make([]Facet, 0)}
	for i := range facets {
		f.Facets = append(f.Facets, facets[i])
	}
	return &f
}

func MakeFacetOr(facets ...Facet) Facet {
	f := facetsOr{Facets: make([]Facet, 0)}
	for i := range facets {
		f.Facets = append(f.Facets, facets[i])
	}
	return &f
}
