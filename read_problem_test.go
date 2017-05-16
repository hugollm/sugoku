package main

import "testing"

func TestReadProblem(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    var expected_vector [81]int
    var expected_string = "108029074600000500020003006201040000500301008000060703900700030002000007340280105"
    for i := 0; i < len(expected_vector); i++ {
        expected_vector[i] = int(expected_string[i] - '0')
    }
    if problem != expected_vector {
        t.Fail()
    }
}
