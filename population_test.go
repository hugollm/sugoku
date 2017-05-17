package main

import "testing"

func TestNewPopulationHasCorrectSize(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    population := NewPopulation(problem, 30)
    if len(population) != 30 {
        t.Fail()
    }
}

func TestNewPopulationHasInitializedIndividuals(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    population := NewPopulation(problem, 30)
    for _, individual := range population {
        if individual.vector == nil {
            t.Fail()
        }
    }
}
