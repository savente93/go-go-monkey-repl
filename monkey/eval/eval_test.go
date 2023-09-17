package eval

import (
	"fmt"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 - 50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for i, tt := range tests {
		fmt.Println(tt.input)
		evaluated := testEval(tt.input)
		testIntegerObject(t, i, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, i int, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object of test %v has wrong value. got=%d, want=%d", i, result.Value, expected)
		return false
	}
	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"true == true", true},
		{"flase == false", false},
		{"true = false", false},
		{"true != false", true},
		{"flase != true", true},
		{" (1 < 2) == true", true},
		{"( 1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1> 2) == false", true},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		p := testBooleanObject(t, i, evaluated, tt.expected)
		if !p {
			t.Errorf("%v", tt.input)
		}
	}
}

func testBooleanObject(t *testing.T, i int, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)

	if !ok {
		t.Errorf("object of test %v is not Boolean. got=%T (+%v)", i, obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object of test %v has wrong value. got=%t, want=%t", i, result.Value, expected)
		return false
	}
	return true

}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, i, evaluated, tt.expected)

	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) {10}", 10},
		{"if (false) {10}", nil},
		{"if (1) { 10 }", 10},
		{"if (1<2){ 10}", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, i, evaluated, int64(integer))
		} else {
			testNullObject(t, i, evaluated)
		}
	}

}

func testNullObject(t *testing.T, i int, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object of test %v is not NULL. got=%T (%+v)", i, obj, obj)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2*5; 9;", 10},
		{"9; return 2*5; 9;", 10},
		{
			`
if (10 > 1) {
	if (10 > 1) {
		return 10;
	}

	return 1;
}
`,
			10,
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, i, evaluated, tt.expected)
	}
}
