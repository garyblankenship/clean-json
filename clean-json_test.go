package main

import (
	"testing"
)

func TestSimpleValidJSONObject(t *testing.T) {
	input := `{"name": "John", "age": 30}`
	expected := `{"name": "John", "age": 30}`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestSimpleValidJSONArray(t *testing.T) {
	input := `[{"name": "John"}, {"name": "Jane"}]`
	expected := `[{"name": "John"}, {"name": "Jane"}]`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestEmbeddedJSONObject(t *testing.T) {
	input := `Random text here and {"key": "value"} more text here`
	expected := `{"key": "value"}`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestMultipleJSONObjects(t *testing.T) {
	input := `{"first": 1} some text {"second": 2}`
	expected := `{"second": 2}` // Assuming the logic prefers the last if same size
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestNestedJSONObject(t *testing.T) {
	input := `{"outer": {"inner": "value"}}`
	expected := `{"outer": {"inner": "value"}}`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestMalformedJSONMissingBrace(t *testing.T) {
	input := `{"incomplete": "test"`
	expected := ``
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestComplexTextWithMultipleJSONFragments(t *testing.T) {
	input := `Here is a {"fragment": "part"}, here is more {"complex": {"nested": "json"}} end of line.`
	expected := `{"complex": {"nested": "json"}}`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestJSONWithArraysAndObjectsCombined(t *testing.T) {
	input := `Start [{"mixed": "array", "obj": {"key": "value"}}] End`
	expected := `[{"mixed": "array", "obj": {"key": "value"}}]`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestEmptyJSONObject(t *testing.T) {
	input := `{}`
	expected := `{}`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestJSONWithSpecialCharactersAndUnicode(t *testing.T) {
	input := `{"emoji": "😊", "text": "special characters like @#&*()"}`
	expected := `{"emoji": "😊", "text": "special characters like @#&*()"}`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestJSONWithEscapedCharacters(t *testing.T) {
	input := `{"data": "Line1\\nLine2"}`
	expected := `{"data": "Line1\\nLine2"}`
	result := findLargestJSON(input)
	if result != expected {
		t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
	}
}

func TestDeeplyNestedJSONObject(t *testing.T) {
    input := `{"level1": {"level2": {"level3": {"level4": "value"}}}}`
    expected := `{"level1": {"level2": {"level3": {"level4": "value"}}}}`
    result := findLargestJSON(input)
    if result != expected {
        t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
    }
}

func TestJSONWithAllDataTypes(t *testing.T) {
    input := `{"string": "text", "number": 100, "boolean": true, "null": null}`
    expected := `{"string": "text", "number": 100, "boolean": true, "null": null}`
    result := findLargestJSON(input)
    if result != expected {
        t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
    }
}

func TestLargeJSONObject(t *testing.T) {
    input := `xx{"key1": "value1", "key2": "value2"}` // Extend this to a very large size
    expected := `{"key1": "value1", "key2": "value2"}` // Expected large JSON
    result := findLargestJSON(input)
    if result != expected {
        t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
    }
}

func TestCorruptedJSONMidway(t *testing.T) {
    input := `{"valid": "yes", "corrupted": {"oops": "no" some invalid json, "validAgain": "yes"}`
    expected := ``
    result := findLargestJSON(input)
    if result != expected {
        t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
    }
}

func TestEmbeddedComplexJSON(t *testing.T) {
    input := `Here is some text before the JSON {"user": {"name": "Jane", "age": 28}, "active": true} and here is some text after.`
    expected := `{"user": {"name": "Jane", "age": 28}, "active": true}`
    result := findLargestJSON(input)
    if result != expected {
        t.Errorf("Test failed, expected '%s', got '%s'", expected, result)
    }
}

