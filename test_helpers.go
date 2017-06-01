package main

func SudokuFromString(str string) [81]int {
    var vector [81]int
    for i := 0; i < 81; i++ {
        vector[i] = int(str[i] - '0')
    }
    return vector
}

func VectorFromString(str string) []float64 {
    vector := make([]float64, len(str))
    for i := 0; i < len(str); i++ {
        vector[i] = float64(int(str[i] - '0'))
    }
    return vector
}
