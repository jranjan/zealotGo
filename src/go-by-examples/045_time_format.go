package main

import (
    "time"
    "fmt"
)

func main() {
    now := time.Now()
    fmt.Println(now)

    fmt.Println(now.Format(time.RFC3339))

    // Parsing is better as it reports error if provided input is not
    // in correct or expected format.
    t1, _ := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    fmt.Println(t1)

    // We cooked our format by using elemental values determined from
    // various helper functions.
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        now.Year(), now.Month(), now.Day(),
        now.Hour(), now.Minute(), now.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e := time.Parse(ansic, "8:41PM")
    fmt.Println(e)

    // The summary is that you can mould the time format as you
    // like by using now.Format().
    fmt.Println(now.Format("3:04PM"))
    fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))
    fmt.Println(now.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    fmt.Println(t2)
}
