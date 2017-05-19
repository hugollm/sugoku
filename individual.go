package main

import "math/rand"
import "time"

type Individual struct {
    vector []int
    fitness int
    slice float64
}

func NewIndividual(problem [81]int) Individual {
    individual := Individual{}
    individual.vector = make([]int, getInputSizeForProblem(problem))
    rand.Seed(time.Now().UTC().UnixNano())
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
