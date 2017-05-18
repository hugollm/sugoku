package main

import "testing"

func TestSelectParents(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    population := NewPopulation(problem, 30)
    parents := SelectParents(population)
    if parents[0].vector == nil || parents[1].vector == nil {
        t.Fail()
    }
}
