package main

func SudokuFromString(str string) [81]int {
    var vector [81]int
    for i := 0; i < 81; i++ {
        vector[i] = int(str[i] - '0')
    }
    return vector
}

func VectorFromString(str string) []int {
    vector := make([]int, len(str))
    for i := 0; i < len(str); i++ {
        vector[i] = int(str[i] - '0')
    }
    return vector
}
