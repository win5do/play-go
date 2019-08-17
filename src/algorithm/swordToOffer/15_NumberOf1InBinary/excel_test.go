package main

import (
	"testing"
)

func TestExcelNumber(t *testing.T) {
	t.Log(excelNumber("Z"))
	t.Log(excelNumber("AA"))
	t.Log(excelNumber("AB"))
	t.Log(excelNumber("AZ"))
	t.Log(excelNumber("AAA"))
	t.Log(excelNumber("123"))
}
