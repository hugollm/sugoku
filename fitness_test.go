package main

import "testing"

func TestMakeCandidateSolution(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    input := []int{3, 5, 6}
    solution := makeCandidateSolution(problem, input)
    expected := SudokuFromString("138529674600000500020003006201040000500301008000060703900700030002000007340280105")
    if solution != expected {
        t.Fail()
    }
}

func TestFitness1(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    input := []int{3, 5, 6}
    if Fitness(problem, input) != 1 {
        t.Fail()
    }
}

func TestFitness27(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    input := VectorFromString("5639347821741598387659769244895216514828193546769")
    if Fitness(problem, input) != 27 {
        t.Fail()
    }
}

func TestFitness0(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    input := VectorFromString("1111111111111111111111111111111111111111111111111")
    if Fitness(problem, input) != 0 {
        t.Fail()
    }
}
