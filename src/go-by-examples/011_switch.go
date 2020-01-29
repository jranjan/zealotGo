package main

import "time"
import "fmt"

func main() {
        switch time.Now().Weekday() {
        case time.Monday:
                fmt.Println("Lazy mood")
        case time.Tuesday, time.Wednesday, time.Thursday:
                fmt.Println("Very hectic day")
        case time.Friday:
                fmt.Println("Learning day")
        default:
                fmt.Println("Relax")
        }
}
