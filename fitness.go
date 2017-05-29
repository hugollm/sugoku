package main

func Fitness(input []int) float64 {
    solution := makeCandidateSolution(problem, input)
    fit := 0.0
    for row := 0; row < 9; row++ {
        fit += evaluateRow(solution, row)
    }
    for col := 0; col < 9; col++ {
        fit += evaluateCol(solution, col)
    }
    for block := 0; block < 9; block++ {
        fit += evaluateBlock(solution, block)
    }
    return fit
}

func makeCandidateSolution(problem [81]int, input []int) [81]int {
    solution := problem
    for i := 0; i < len(input); i++ {
        for j := 0; j < len(solution); j++ {
            if solution[j] == 0 {
                solution[j] = input[i]
                break
            }
        }
    }
    return solution
}

func evaluateRow(solution [81]int, row int) float64 {
    fit := 0
    set := make(map[int]bool)
    for col := 0; col < 9; col++ {
        i := (row * 9) + col
        if solution[i] != 0 && set[solution[i]] == false {
            fit++
        }
        set[solution[i]] = true
    }
    return float64(fit)
}

func evaluateCol(solution [81]int, col int) float64 {
    fit := 0
    set := make(map[int]bool)
    for row := 0; row < 9; row++ {
        i := (row * 9) + col
        if solution[i] != 0 && set[solution[i]] == false {
            fit++
        }
        set[solution[i]] = true
    }
    return float64(fit)
}

func evaluateBlock(solution [81]int, block int) float64 {
    fit := 0
    set := make(map[int]bool)
    for row := 0; row < 3; row++ {
        for col := 0; col < 3; col++ {
            i := ((block / 3) * 27) + ((block / 3) * 3) + (row * 9) + col
            if solution[i] != 0 && set[solution[i]] == false {
                fit++
            }
            set[solution[i]] = true
        }
    }
    return float64(fit)
}
