package main

import "io/ioutil"
import "strings"

func ReadProblem(path string) [81]int {
    bytes, e := ioutil.ReadFile(path)
    if (e != nil) {
        panic(e)
    }
    str := string(bytes)
    str = strings.Replace(str, "\n", "", -1)
    str = strings.Replace(str, ".", "0", -1)
    var vector [81]int
    for i := 0; i < len(str); i++ {
        vector[i] = int(str[i] - '0')
    }
    return vector
}
