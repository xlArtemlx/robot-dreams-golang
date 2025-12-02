package main

import (
	"fmt";
	"math";
	"strconv"
)

func FibonacciIterative(n int) int {
	if n >= 0 {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			x, y = y, x+y
		}
		return x
	}
 	return n
}

func FibonacciRecursive(n int) int {
	switch {
		case n == 0:
			return 0
		case n == 1:
			return 1
		case n > 1:
			return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
		default:
			return n
	}
}

func IsPrime(n int) bool {
	for i, s := 2, int(math.Sqrt(float64(n))); i <= s; i++ {
		if n%i == 0 {
			return false
		}
	}
 	return n > 1
}

func IsBinaryPalindrome(n int) bool {
	binary := strconv.FormatInt(int64(n), 2)
	i, j := 0, len(binary)-1
    for i < j {
        if binary[i] != binary[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func ValidParentheses(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{
        '(': ')',
        '[': ']',
        '{': '}',
    }

    for _, symbol := range s {
        if want, ok := pairs[symbol]; ok {
            stack = append(stack, want)
            continue
        }
        if len(stack) == 0 || stack[len(stack)-1] != symbol {
            return false
        }
        stack = stack[:len(stack)-1]
    }
	
    return len(stack) == 0
}

func Increment(num string) int {
    for _, currentNumber := range num {
        if currentNumber != '0' && currentNumber != '1' {
            return 0
        }
    }
	value, error := strconv.ParseInt(num, 2, 64)
	if error != nil {
		return 0
	}
	return int(value + 1)
}

func main() {
 // Невеликі демонстраційні виклики (для наочного запуску `go run .`)
 fmt.Println("FibonacciIterative(10):", FibonacciIterative(10)) // очікуємо 55
 fmt.Println("FibonacciRecursive(10):", FibonacciRecursive(10)) // очікуємо 55

 fmt.Println("IsPrime(2):", IsPrime(2))     // true
 fmt.Println("IsPrime(15):", IsPrime(15))   // false
 fmt.Println("IsPrime(29):", IsPrime(29))   // true

 fmt.Println("IsBinaryPalindrome(7):", IsBinaryPalindrome(7))   // true (111)
 fmt.Println("IsBinaryPalindrome(6):", IsBinaryPalindrome(6))   // false (110)

 fmt.Println(`ValidParentheses("[]{}()"):`, ValidParentheses("[]{}()"))     // true
 fmt.Println(`ValidParentheses("[{]}"):`, ValidParentheses("[{]}"))         // false

 fmt.Println(`Increment("101") ->`, Increment("101")) // 6
}