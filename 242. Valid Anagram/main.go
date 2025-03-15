package main 

import "fmt"

func isAnagram(s string, t string) bool {
    m := make(map[rune]int)

    for _, el := range s {
        m[el]++
    }

    for _, el := range t {
        m[el]--
    }

    for _, el := range m {
        if el != 0 {
            return false
        }
    }

    return true
}

func main() {
    s1 := "anagram"
    t1 := "nagaram"

    s2 := "rat"
    t2 := "car"

    fmt.Println(isAnagram(s1, t1))
    fmt.Println(isAnagram(s2, t2))
}