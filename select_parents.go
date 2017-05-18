package main

import "math/rand"

func SelectParents(population []Individual) [2]Individual {
    size := len(population)
    a := population[rand.Intn(size)]
    b := population[rand.Intn(size)]
    return [2]Individual{a, b}
}
