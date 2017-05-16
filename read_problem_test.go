package main

import "testing"

func TestReadProblem(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    expected := SudokuFromString("108029074600000500020003006201040000500301008000060703900700030002000007340280105")
    if problem != expected {
        t.Fail()
    }
}
