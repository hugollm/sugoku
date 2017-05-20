package main

import "testing"
import "reflect"

func TestMutateWithMinProbability(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    offspring := [2]Individual{NewIndividual(problem), NewIndividual(problem)}
    offspring_copy := copyOffspring(offspring)
    Mutate(offspring, 0)
    if ! reflect.DeepEqual(offspring[0].vector, offspring_copy[0]) || ! reflect.DeepEqual(offspring[1].vector, offspring_copy[1]) {
        t.Fail()
    }
}

func TestMutateWithMaxProbability(t *testing.T) {
    problem := ReadProblem("problems/simple/problem")
    offspring := [2]Individual{NewIndividual(problem), NewIndividual(problem)}
    offspring_copy := copyOffspring(offspring)
    Mutate(offspring, 1)
    if reflect.DeepEqual(offspring[0].vector, offspring_copy[0]) || reflect.DeepEqual(offspring[1].vector, offspring_copy[1]) {
        t.Fail()
    }
}

func copyOffspring(offspring [2]Individual) [2][]int {
    size := len(offspring[0].vector)
    copies := [2][]int{make([]int, size), make([]int, size)}
    copy(copies[0], offspring[0].vector)
    copy(copies[1], offspring[1].vector)
    return copies
}
