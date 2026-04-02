package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isVariable(token string) bool {
	return len(token) == 1 && token[0] >= 'A' && token[0] <= 'Z'
}

func variableIndex(token string) int {
	return int(strings.ToUpper(token)[0]-'A') + 1
}

func emit(out *[]int, values ...int) {
	*out = append(*out, values...)
}

func emitOperand(token string, out *[]int) {
	if isVariable(token) {
		emit(out, BID, variableIndex(token))
	} else {
		value, _ := strconv.Atoi(token)
		emit(out, BCONST, value)
	}
}

func compileLine(tokens []string) ([]int, error) {
	if len(tokens) == 0 {
		return nil, nil
	}

	var out []int

	lineNum, err := strconv.Atoi(tokens[0])
	if err != nil {
		return nil, fmt.Errorf("expected line number, got %q", tokens[0])
	}
	emit(&out, BLINE, lineNum)

	if len(tokens) < 2 {
		return out, nil
	}

	statement := strings.ToUpper(tokens[1])

	switch statement {

	case "STOP":
		emit(&out, BSTOP, 0)

	case "GOTO":
		if len(tokens) < 3 {
			return nil, fmt.Errorf("GOTO requires a target line number")
		}
		target, _ := strconv.Atoi(tokens[2])
		emit(&out, BGOTO, target)

	case "PRINT":
		if len(tokens) < 3 {
			return nil, fmt.Errorf("PRINT requires a variable")
		}
		emit(&out, BPRINT, 0)
		emit(&out, BID, variableIndex(tokens[2]))

	case "IF":
		if len(tokens) < 6 {
			return nil, fmt.Errorf("IF requires: IF <left> <op> <right> <target>")
		}
		left, op, right := tokens[2], tokens[3], tokens[4]
		target, _ := strconv.Atoi(tokens[5])

		emit(&out, BIF, 0)
		emitOperand(left, &out)
		emit(&out, BOP, opValue[op])
		emitOperand(right, &out)
		emit(&out, BGOTO, target)

	default:
		if len(tokens) < 3 {
			return nil, fmt.Errorf("malformed assignment")
		}
		varName := tokens[1]
		expr := tokens[3:]

		emit(&out, BID, variableIndex(varName))
		emit(&out, BOP, opValue["="])

		switch len(expr) {
		case 1:
			emitOperand(expr[0], &out)
		case 3:
			emitOperand(expr[0], &out)
			emit(&out, BOP, opValue[expr[1]])
			emitOperand(expr[2], &out)
		default:
			return nil, fmt.Errorf("invalid expression after '='")
		}
	}

	return out, nil
}
