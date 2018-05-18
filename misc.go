package main

import (
	"fmt"
	"strings"
)

func StrFormat(format string, args ...interface{}) string {
    args2 := make([]string, len(args))
    for i, v := range args {
        if i%2 == 0 {
            args2[i] = fmt.Sprintf("{%v}", v)
        } else {
            args2[i] = fmt.Sprint(v)
        }
    }
    r := strings.NewReplacer(args2...)
    return r.Replace(format)
}