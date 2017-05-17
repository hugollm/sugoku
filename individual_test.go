package main

import "testing"

func TestNewIndividualHasProperSize(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    individual := NewIndividual(problem)
    if len(individual.vector) != 49 {
        t.Fail()
    }
}

func TestNewIndividualHasDifferentValues(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    individual := NewIndividual(problem)
    set := make(map[int]bool)
    different_numbers_count := 0
    for _, cell := range individual.vector {
        if ! set[cell] {
            different_numbers_count++
        }
        set[cell] = true
    }
    if different_numbers_count <= 1 {
        t.Fail()
    }
}
