package main

import "os"
import "fmt"

func main() {
    problem := ReadProblem(os.Args[1])
    const populationSize int = 30
    population := NewPopulation(problem, populationSize)
    best := population[0]
    for generation := 0; generation < 9999999; generation++ {
        for _, individual := range population {
            individual.fitness = Fitness(problem, individual.vector)
            if &individual != &best && individual.fitness >= best.fitness {
                best = individual
            }
        }
        newPopulation := NewPopulation(problem, populationSize)
        for i := 0; i < (populationSize/2); i++ {
            parents := SelectParents(population)
            offspring := Crossover(parents, 0.7)
            offspring = Mutate(offspring, 0.001)
            newPopulation[i+i] = offspring[0]
            newPopulation[i+i+1] = offspring[1]
        }
        population = newPopulation
        population[0] = best
        if generation % 100 == 0 {
            printPopulationFitness(population)
        }
    }
}

func printPopulationFitness(population []Individual) {
    for _, individual := range population {
        fmt.Printf("%02d ", individual.fitness)
    }
    fmt.Println()
}
