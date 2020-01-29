package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)



func main() {
    // DisplayUserMatrix(2, 6)
    // AddTwoMatrix(2, 4)
    AddMultipleMatrix(2,4,3)
}

func AddMultipleMatrix(row int, col int, count int){
    m, err := ReadMultipleMatrix(row, col, 3)
    if err != nil {
        fmt.Printf("Error occured!")
    }

    am, err := create_2D_matrix(row, col)
    if err != nil {
        fmt.Printf("Error occured!")
    }

    for c:=0; c<count; c++{
        for i:=0; i<row; i++{
            for j:=0; j<col; j++{
                am[i][j] = am[i][j] + m[c][i][j]
            }
        }
    }

    fmt.Println(am)
}

func AddTwoMatrix(row int, col int){
    m1, err := ReadMatrix(row, col)
    if err != nil {
        fmt.Printf("Error occured!")
    }
    fmt.Println(m1)

    m2, err := ReadMatrix(row, col)
    if err != nil {
        fmt.Printf("Error occured!")
    }
    fmt.Println(m2)

    am, err := create_2D_matrix(row, col)
    for i:=0; i<row; i++{
        for j:=0; j<col; j++{
            am[i][j] = m1[i][j] + m2[i][j]
        }
    }
    fmt.Println(am)
}


func DisplayUserMatrix(row int, col int){
    m, err := ReadMatrix(2,4)
    if err != nil {
        fmt.Printf("Error occured!")
    }
    DisplayMatrix(m)
}

func DisplayMatrix(m [][]int) {
    rows := len(m)
    cols := len(m[0])
    fmt.Printf("Matrix [%d X %d]:\n", rows, cols)
    fmt.Println(m)
}

func create_2D_matrix(row int, col int) ([][]int, error) {
    m := make([][]int, row)
    for i:=0; i<row; i++ {
        m[i] = make([]int, col)
    }
    return m, nil
}

func create_3D_matrix(row int, col int, count int) ([][][]int, error) {
    var err error
    m := make([][][]int, count)
    for i:=0; i<count; i++ {
        m[i], err = create_2D_matrix(row, col)
        if err != nil {
            fmt.Printf("Error occured!")
        }
    }
    return m, nil
}


func ReadMatrix(row int, col int) ([][]int, error) {
    var err error
    m, _ := create_2D_matrix(row, col)

    for i:=0; i<row; i++ {
        fmt.Printf("Enter row %d:\n", i)
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        line := scanner.Text()
        nums := strings.Split(line, " ")
        for j:=0; j<col; j++{
            m[i][j], err = strconv.Atoi(nums[j])
            if err != nil {
                return nil, err
            }
        }
    }

    return m, nil
}

func ReadMultipleMatrix(row int, col int, count int) ([][][]int, error){
    var err error
    m, _ := create_3D_matrix(row, col, count)

    for c:=0; c<count; c++{
        fmt.Printf("Enter matrix ID = %d\n", c);
        for i:=0; i<row; i++ {
            fmt.Printf("Enter row %d:\n", i)
            scanner := bufio.NewScanner(os.Stdin)
            scanner.Scan()
            line := scanner.Text()
            nums := strings.Split(line, " ")
            for j:=0; j<col; j++{
                m[c][i][j], err = strconv.Atoi(nums[j])
                if err != nil {
                    return nil, err
                }
            }
        }
    }

    return m, nil
}




// Read a square matrix of order n. And swap rows and colums of matrix as per user choice.



// Write a program to remove duplicate words from a string


// Write a program to find frequency occurence of words from a string


// Write a program to read n strings and search a word in list of string