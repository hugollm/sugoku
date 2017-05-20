package main

import "math/rand"

type Individual struct {
    vector []int
    fitness int
    slice float64
}

func NewIndividual(problem [81]int) Individual {
    individual := Individual{}
    individual.vector = make([]int, getInputSizeForProblem(problem))
    for i := 0; i < len(individual.vector); i++ {
        individual.vector[i] = (rand.Int() % 9) + 1
    }
    return individual
}

func getInputSizeForProblem(problem [81]int) int {
    size := 0
    for _, n := range problem {
        if n == 0 {
            size++
        }
    }
    return size
}

func CopyIndividual(individual Individual) Individual {
    newIndividual := Individual{}
    size := len(individual.vector)
    newIndividual.vector = make([]int, size)
    for i := 0; i < size; i++ {
        newIndividual.vector[i] = individual.vector[i]
    }
    return newIndividual
}
