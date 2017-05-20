package main

import "math/rand"

func Crossover(parents [2]Individual, probability float64) [2]Individual {
    if rand.Float64() > probability {
        return [2]Individual{CopyIndividual(parents[0]), CopyIndividual(parents[1])}
    }
    size := len(parents[0].vector)
    offspring := makeEmptyIndividuals(size)
    point := rand.Intn(size)
    for i := 0; i < size; i++ {
        if i < point {
            offspring[0].vector[i] = parents[0].vector[i]
            offspring[1].vector[i] = parents[1].vector[i]
        } else {
            offspring[0].vector[i] = parents[1].vector[i]
            offspring[1].vector[i] = parents[0].vector[i]
        }
    }
    return offspring
}

func makeEmptyIndividuals(size int) [2]Individual {
    a := Individual{}
    b := Individual{}
    a.vector = make([]int, size)
    b.vector = make([]int, size)
    return [2]Individual{a, b}
}
