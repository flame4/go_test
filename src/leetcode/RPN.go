package leetcode

import (
	_ "container/list"
	_ "reflect"
	"fmt"
	_ "strings"
)

// 逆波兰表达式
// ((3+5*2)+2)/5+6/3*2+3
// 352*+2+5/63/2*+3+
// 规则:
// 1. 如果是数字,直接入
// 2. 如果是 + -

type Stack struct {
	list []interface{}
	len int
}

func (s *Stack) Pop() interface{} {
	if(s.len <= 0) {
		return nil
	} else {
		result := s.list[s.len-1]
		s.len--
		return result
	}
}

func (s *Stack) PushBack(v interface{}) {
	s.list = append(s.list, v)
	s.len++
}

func (s *Stack) Back() interface{} {
	return s.list[s.len]
}

func (s *Stack) Len() int {
	return s.len
}


func rpn(mid_express string) string {
	stack_symbol := new(Stack)
	stack_out := new(Stack)

	for _, bt := range mid_express {
		s := string(bt)
		fmt.Println(s)
		switch s {
		case "(":
			stack_symbol.PushBack(s)
		case ")":
			for stack_symbol.Back().(string) != "(" {
				stack_out.PushBack(stack_symbol.Back())
				stack_symbol.Pop()
			}
			stack_symbol.Pop()
		case "+", "-":
			if(stack_symbol.Len() == 0) {
				stack_symbol.PushBack(s)
			} else {
				for stack_symbol.Back().(string) != "(" {
					stack_out.PushBack(stack_symbol.Back().(string))
					fmt.Println("before remove: ", stack_symbol.Back().(string))
					stack_symbol.Pop()
					fmt.Println("After remove: ", stack_symbol.Back().(string))
				}
				stack_symbol.PushBack(s)
			}
		case "*", "/":
			for stack_symbol.Back().(string) == "*" && stack_symbol.Back().(string) == "/" {
				stack_out.PushBack(stack_symbol.Back().(string))
				stack_symbol.Pop()
			}
			stack_symbol.PushBack(s)
		default:
			stack_out.PushBack(s)
		}
	}
	for stack_symbol.Len() > 0 {
		stack_out.PushBack(stack_symbol.Back().(string))
		stack_symbol.Pop()
	}
	result := ""
	for _, s := range stack_out.list {
		result = s.(string) + result
	}
	return result
}




