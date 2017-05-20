package main

import "math/rand"

func Mutate(offspring [2]Individual, probability float64) {
    for i := 0; i < len(offspring[0].vector); i++ {
        if shouldMutate(probability) {
            offspring[0].vector[i] = randomValue()
        }
        if shouldMutate(probability) {
            offspring[1].vector[i] = randomValue()
        }
    }
}

func shouldMutate(probability float64) bool {
    return rand.Float64() < probability
}

func randomValue() int {
    return (rand.Int() % 9) + 1
}
