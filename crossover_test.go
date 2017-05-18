package main

import "testing"
import "reflect"

func TestCrossoverWithMinProbability(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    parents := [2]Individual{NewIndividual(problem), NewIndividual(problem)}
    offspring := Crossover(parents, 0)
    if ! reflect.DeepEqual(offspring[0].vector, parents[0].vector) || ! reflect.DeepEqual(offspring[1].vector, parents[1].vector) {
        t.Fail()
    }
}

func TestCrossoverWithMaxProbability(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    parents := [2]Individual{NewIndividual(problem), NewIndividual(problem)}
    offspring := Crossover(parents, 1)
    if reflect.DeepEqual(offspring[0].vector, parents[0].vector) || reflect.DeepEqual(offspring[1].vector, parents[1].vector) {
        t.Fail()
    }
}

func TestCrossoverOffspringHaveValidVectors(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    parents := [2]Individual{NewIndividual(problem), NewIndividual(problem)}
    offspring := Crossover(parents, 1)
    for _, individual := range offspring {
        for _, value := range individual.vector {
            if value == 0 {
                t.Fail()
            }
        }
    }
}
