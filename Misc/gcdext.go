package main

import "fmt"

func extendedGCD(a, b int) (int, int, int) {
    if b == 0 {
        return a, 1, 0
    }
    gcd, x1, y1 := extendedGCD(b, a%b)
    x := y1
    y := x1 - (a/b)*y1
    return gcd, x, y
}

func main() {
    a := 48
    b := 18
    gcd, x, y := extendedGCD(a, b)
    fmt.Printf("GCD(%d, %d) = %d\n", a, b, gcd)
    fmt.Printf("x = %d, y = %d\n", x, y)
}