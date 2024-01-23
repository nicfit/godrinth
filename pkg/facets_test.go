package godrinth

import (
	"fmt"
	"testing"
)

func TestOperandTypeOperandString(t *testing.T) {
	parameters := []struct {
		input    OperandType
		expected string
	}{
		{OperandTypeEqual, "="}, {OperandTypeNotEqual, "!="},
		{OperandTypeLessThan, "<"}, {OperandTypeLessThanOrEqual, "<="},
		{OperandTypeGreaterThan, ">"}, {OperandTypeGreaterThanOrEqual, ">="},
	}

	for i := range parameters {
		t.Run(fmt.Sprintf("Testing [%v]", parameters[i].input.String()), func(t *testing.T) {
			actual := parameters[i].input.OperandString()
			if actual != parameters[i].expected {
				t.Logf("expected:%s: , actual:%s", parameters[i].expected, actual)
				t.Fail()
			}
		})
	}
}

func TestConcreteFacet(t *testing.T) {
	f := MakeFacet("project_type", OperandTypeEqual, "mod")
	s := f.String()
	fmt.Println(s)
	if f.String() != "\"project_type=mod\"" {
		t.Fail()
	}
}

func TestFacetOr(t *testing.T) {
	f1 := MakeFacet("project_type", OperandTypeEqual, "mod")
	f2 := MakeFacet("version", OperandTypeGreaterThanOrEqual, "1.20")
	or := MakeFacetOr(f1, f2)
	fmt.Println("OR:", or)
	if or.String() != "[\"project_type=mod\", \"version>=1.20\"]" {
		t.Fail()
	}
}

func TestFacetAnd(t *testing.T) {
	f1 := MakeFacet("project_type", OperandTypeEqual, "mod")
	f2 := MakeFacet("version", OperandTypeGreaterThanOrEqual, "1.20")
	f3 := MakeFacet("categories", OperandTypeEqual, "fabric")
	and := MakeFacetAnd(f1, f2, f3)
	fmt.Println("AND:", and)
	if and.String() != "[\"project_type=mod\"],[\"version>=1.20\"],[\"categories=fabric\"]" {
		t.Fail()
	}
}

func TestFacetComplex(t *testing.T) {
	facets := MakeFacetAnd(
		MakeFacetOr(MakeFacet("version", OperandTypeGreaterThanOrEqual, "1.20"),
			MakeFacet("version", OperandTypeNotEqual, "1.20.3")),
		MakeFacetOr(MakeFacet("categories", OperandTypeEqual, "fabric"),
			MakeFacet("categories", OperandTypeEqual, "quilt")),
		MakeFacet("project_type", OperandTypeEqual, "mod"),
	)

	fmt.Println(facets.String())

}
