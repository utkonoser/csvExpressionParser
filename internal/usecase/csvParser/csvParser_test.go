package csvParser

import (
	"reflect"
	"testing"
)

func TestHashTableOfCsvFile(t *testing.T) {
	const dataPath = "../../../data/myTest.csv"
	HashTableOfCsvFile(dataPath)
	testHash := map[string]map[string]string{
		"A":    {"1": "1", "2": "2", "30": "0"},
		"B":    {"1": "0", "2": "=A1+Cell30", "30": "=B1+A1"},
		"Cell": {"1": "1", "2": "0", "30": "5"}}
	if !reflect.DeepEqual(testHash, csvTable) {
		t.Error("wrong result, testHash and csvTable must be equal")
	}
	testHash2 := map[string]map[string]string{
		"A":    {"11111": "1", "2": "2", "30": "0"},
		"B":    {"1": "0", "2": "=A1+Cell30", "30": "=B1+A1"},
		"Cell": {"1": "1", "2": "0", "30": "5"}}
	if reflect.DeepEqual(testHash2, csvTable) {
		t.Error("wrong result, testHash2 and csvTable does not equal")
	}
}

func TestExpression_Action(t *testing.T) {
	exp := expression{
		arg1:      "5",
		arg2:      "5",
		operation: "-",
	}
	val, _ := exp.Action()
	if val != 0 {
		t.Errorf("wrong result %v not equal %v", val, 0)
	}
	exp.operation = "*"
	val, _ = exp.Action()
	if val != 25 {
		t.Errorf("wrong result %v not equal %v", val, 25)
	}
	exp.operation = "+"
	val, _ = exp.Action()
	if val != 10 {
		t.Errorf("wrong result %v not equal %v", val, 10)
	}
	exp.operation = "/"
	val, _ = exp.Action()
	if val != 1 {
		t.Errorf("wrong result %v not equal %v", val, 1)
	}
	exp.arg2 = "0"
	_, err := exp.Action()
	if err == nil {
		t.Error(err)
	}

	exp.arg1 = "T"
	_, err = exp.Action()
	if err == nil {
		t.Error(err)
	}
}
