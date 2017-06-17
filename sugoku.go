package main

import "os"
import "fmt"
import "github.com/hugollm/exodus"

var problem [81]int

func main() {
    problem = ReadProblem(os.Args[1])
    search := exodus.Search{
        IndividualSize: InputSize(problem),
        PopulationSize: 20,
        CrossoverRate: 0.6,
        MutationRate: 0.005,
        MigrationRate: 0.005,
        NewGene: newGene,
        Fitness: Fitness,
        OnGeneration: onGeneration,
    }
    search.Start()
}

func newGene() float64 {
    return float64(exodus.RandomInt(1, 9))
}

func onGeneration(search *exodus.Search) {
    if search.Best.Fitness == 243 {
        printStatus(search)
        search.Stop()
    } else if search.Generation % 1000 == 0 {
        printStatus(search)
    }
}

func printStatus(search *exodus.Search) {
    for _, individual := range search.Population.Individuals {
        fmt.Printf("%03d ", int(individual.Fitness))
    }
    fmt.Print(search.Best.Genome, int(search.Best.Fitness), search.Generation)
    fmt.Println()
}
