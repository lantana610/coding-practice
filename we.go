package main

import (
    "strings"
)

func FixArticle(s string) string {
    word := strings.Fields(s)
    for i := 0; i < len(word)-1; i++ {
        current := strings.ToLower(word[i])
        nextword := strings.ToLower(word[i+1])
        if len(nextword) > 0 {
            switch nextword[0] {
            case 'a', 'i', 'e', 'u', 'h', 'o':
                if current == "a" {
                    word[i] = "an"
                }
        
            default:
                if current == "an" {
                    word[i] = "a"
                }
            }
        }
    }
    return strings.Join(word, " ")
}
