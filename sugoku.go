package main

import "os"
import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    problem := ReadProblem(os.Args[1])
    const populationSize int = 30
    population := NewPopulation(problem, populationSize)
    best := population[0]
    for generation := 0; generation < 9999999999; generation++ {

        for i := 0; i < len(population); i++ {
            population[i].fitness = Fitness(problem, population[i].vector)
            if &population[i] != &best && population[i].fitness >= best.fitness {
                best = population[i]
            }
        }

        if best.fitness == 243 {
            printStatus(population, best)
            break
        }

        if generation % 100 == 0 {
            printStatus(population, best)
        }

        newPopulation := make([]Individual, populationSize)
        for i := 0; i < (populationSize/2); i++ {
            parents := SelectParents(population)
            offspring := Crossover(parents, 0.6)
            Mutate(offspring, 0.01)
            newPopulation[i+i] = CopyIndividual(offspring[0])
            newPopulation[i+i+1] = CopyIndividual(offspring[1])
        }

        population = newPopulation
        population[0] = best
    }
}

func printStatus(population []Individual, best Individual) {
    for _, individual := range population {
        fmt.Printf("%03d ", individual.fitness)
    }
    fmt.Print(best.vector, best.fitness)
    fmt.Println()
}
