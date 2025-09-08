package main

import (
	"fmt"
    "sort"
)


// 流程控制
// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
    m := make(map[int]int)
    for _,v := range nums {
        _,ok := m[v]
        if ok {
            m[v] = m[v] + 1
        }else {
            m[v] = 1
        }
    }
    for k,v := range m {
        if v == 1 {
            return k
        }
    }
    return -1

}


// 字符串
// 有效的括号 

// 考察：字符串处理、栈的使用

// 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
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


// 最长公共前缀
// 考察：字符串处理、循环嵌套
// 题目：查找字符串数组中的最长公共前缀
func longestCommonPrefix(strs []string) string {
    
    l := 0
hwloop:
        for i,s := range strs[0] {
            for _, os := range strs[1:] {
                if i == len(os) {
                    break hwloop
                }
                if string(s) != string(os[i]) {
                    break hwloop
                }
            }
            l = i + 1
        }
    return strs[0][:l]
}

// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
    l := len(digits)
    v := 1
    t := false
    for i:=l-1; i>=0; i-- {
        if digits[i] + v >= 10 {
            digits[i] = 0
            v = 1
            if i == 0 {
                t = true
            }
        }else {
            digits[i] = digits[i] + 1

            return digits
        }

    }
    if t {

        return append([]int{1},digits...)
    }
    return digits
}


// 删除有序数组中的重复项
// 通过set判断重复的写法，有点垃圾
func removeDuplicates(nums []int) int {
    data := make([]int, 0)
    set := make(map[int]struct{})
    // rdata int [len(nums)]
    
    for _,v := range nums {
        _,ok := set[v]
        if ok {
            // append(rdata)
            continue
        }else {
            set[v] = struct{}{}
            fmt.Printf("data:%v,v:%v\n", data, v)
            data = append(data, v)
        }
    }
    for i,k := range data {
        if k == 0 {
            break
        }
        nums[i] = k
    }
    return len(set)
}

// 删除有序数组中的重复项
// 双指针的写法，简洁高效
func removeDuplicates2(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    f := 1
    s := 1
    for f < len(nums) {
        if nums[f] != nums[f-1] {
            nums[s] = nums[f]
            s+=1
        }
        f += 1
    }
    return s
}



// 合并区间，腾讯一面的题目
func mergeRange(intervals [][]int) [][]int  {

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    merged := make([][]int, 0)
    for _,interval := range intervals {
        if len(merged) > 0 && interval[0] <= merged[len(merged) - 1][1] {

            merged[len(merged) - 1][1] = max(interval[1],merged[len(merged) - 1][1])
        }else {
            merged = append(merged, interval)
        }
    }
    return merged
}
// 插入区间 腾讯面试题
func mergeRange2(intervals [][]int, newInterval []int) [][]int {
    
    merged := make([][]int, 0)

    i := 0
    for intervals[i][1] < newInterval[0] {
        merged = append(merged, intervals[i])
        i += 1
    }
    for i < len(intervals) && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(intervals[i][0], newInterval[0])
        newInterval[1] = max(intervals[i][1], newInterval[1])
        i +=1
    }
    merged = append(merged, newInterval)
    merged = append(merged, intervals[i:]...)

    return merged
}

func main() {
	// fmt.Printf("val:%t\n", isValid("{([])}"))
	fmt.Printf("helloworld\n")
    // strs := [...]string{"abcd","flow","flight"}
    // for _, os := range strs {
    //     fmt.Printf("char: %v", os[0])
    // }
    // data :=  [6] int{0,1,1,2,2,3}
    // sdata := data[:]
    // fmt.Printf("rtn: %v",removeDuplicates2(sdata))
    // intervals := [4][2] int {{1,3},{2,6},{8,10},{15,18}}
    intervals := make([][]int, 3)
    intervals[0] = make([]int, 0)
    intervals[0] = append(intervals[0], 1,3)


    intervals[1] = make([]int, 0)
    intervals[1] = append(intervals[1], 2,6)

    intervals[2] = make([]int, 0)
    intervals[2] = append(intervals[2], 8,10)
    // sintervals := intervals[:]
    fmt.Printf("v-> %v\n", intervals)
    result := mergeRange(intervals)
    for _,v := range result {
        fmt.Printf("v-> %v\n", v)
    }
}

