package dictionary_tools

import (
    "regexp"
    "github.com/ddliu/go-dict/util"
)

func NewSimpleDict() *MySimpleDict {
    return &MySimpleDict{}
}

type MySimpleDict struct {
    Words []string
}

func (this *MySimpleDict) AddWordsList(words []string) {
    for _, w := range words {
        this.Words = append(this.Words, w)
    }
}

func (this *MySimpleDict) Count() int {
    return len(this.Words)
}

func (this *MySimpleDict) Load(dict string) {
    util.WalkFileLines(dict, func(line string) bool {
        this.Words = append(this.Words, string(line[:]))
        return true
    })
}

func (this *MySimpleDict) Lookup(pattern string, offset int, limit int) []string {
    if pattern == "" {
        if offset == 0 && limit == 0 {
            return this.Words[:]
        }
        if offset >= len(this.Words) {
            return nil
        }

        var end int
        if limit <= 0 {
            end = len(this.Words)
        } else {
            end = offset + limit
        }

        return this.Words[offset:end]
    }

    var compiledPattern = regexp.MustCompile("^" + pattern + "$")
    var matched []string
    found := 0
    this.Walk(func(word string) bool {
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

func (this *MySimpleDict) Walk(f func(string) bool) {
    for i := 0; i < len(this.Words); i++ {
        if !f(this.Words[i]) {
            break
        }
    }
}
