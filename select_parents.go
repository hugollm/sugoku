package main

import "math/rand"

func SelectParents(population []Individual) [2]Individual {
    calculateSlices(population)
    a := selectOne(population)
    b := selectOne(population)
    return [2]Individual{a, b}
}

func calculateSlices(population []Individual) {
    sumFitness := sumFitness(population)
    for i := 0; i < len(population); i++ {
        previous_slice := 0.0
        if i > 0 {
            previous_slice = population[i-1].slice
        }
        population[i].slice = previous_slice + float64(population[i].fitness) / float64(sumFitness)
    }
}

func selectOne(population []Individual) Individual {
    r := rand.Float64()
    for _, individual := range population {
        if r < individual.slice {
            return individual
        }
    }
    return population[len(population)-1]
}

func sumFitness(population []Individual) int {
    sum := 0
    for _, individual := range population {
        sum += individual.fitness
    }
    return sum
}
