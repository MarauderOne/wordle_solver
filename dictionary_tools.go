package main

import (
    "regexp"
    "github.com/ddliu/go-dict/util"
)

func newSimpleDict() *mySimpleDict {
    return &mySimpleDict{}
}

type mySimpleDict struct {
    words []string
}

func (this *mySimpleDict) addWordsList(words []string) {
    for _, w := range words {
        this.words = append(this.words, w)
    }
}

func (this *mySimpleDict) count() int {
    return len(this.words)
}

func (this *mySimpleDict) load(dict string) {
    util.WalkFileLines(dict, func(line string) bool {
        this.words = append(this.words, string(line[:]))
        return true
    })
}

func (this *mySimpleDict) lookup(pattern string, offset int, limit int) []string {
    if pattern == "" {
        if offset == 0 && limit == 0 {
            return this.words[:]
        }
        if offset >= len(this.words) {
            return nil
        }

        var end int
        if limit <= 0 {
            end = len(this.words)
        } else {
            end = offset + limit
        }

        return this.words[offset:end]
    }

    var compiledPattern = regexp.MustCompile("^" + pattern + "$")
    var matched []string
    found := 0
    this.walk(func(word string) bool {
        if compiledPattern.MatchString(word) {
            found++

            if found < offset + 1 {
                return true
            }

            matched = append(matched, word)

            if limit > 0 && found >= offset + limit  {
                return false
            }            
        }

        return true
    })

    return matched
}

func (this *mySimpleDict) walk(f func(string) bool) {
    for i := 0; i < len(this.words); i++ {
        if !f(this.words[i]) {
            break
        }
    }
}
