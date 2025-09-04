package main

import (
	"fmt"
)


func isValid(s string) bool {
    ms := make(map[string]string)
    ms["}"] = "{"
    ms[")"] = "("
    ms["]"] = "["

    var sq Stack
    for _, t := range s {
        r,ok := ms[string(t)]
        if ok {
            if len(sq.data) <= 0 || sq.data[len(sq.data)-1] != r {
                return false
            }
            sq.pop()
        }else {
            sq.push(string(t))
        }
    }
    _, res := sq.pop()
    return !res
}

type Stack struct {
    data [] string
}

func (s *Stack) push(val string) {
    s.data = append(s.data, val)
}

func (s *Stack) pop() (string,bool) {
    if len(s.data) <= 0 {
        return "", false
    }
    val := s.data[len(s.data) - 1]
    s.data = s.data[:len(s.data) - 1]
    return val, true
}

func main() {
	// ms := make(map[string]string)
    // ms["}"] = "{"
    // ms[")"] = "("
    // ms["]"] = "["
	// s := "{([])}"
	// for _, t := range s {
    //     // r,ok := ms[t]
	// 	fmt.Printf("helloworld%c,", t)
	// }

	// p := "{([])}"
	// for _, s := range p {
	// 	fmt.Printf("索引 %d, 字符 ", s)
	// }


	// for i, r := range p {
	// 	fmt.Printf("索引 %d, 字符 %c\n", i, r)
	// }
	fmt.Printf("val:%t\n", isValid("{([])}"))
	fmt.Printf("helloworld")
}

