package main

func NewPopulation(problem [81]int, size int) []Individual {
    population := make([]Individual, size)
    for i := 0; i < size; i++ {
        population[i] = NewIndividual(problem)
    }
    return population
}
