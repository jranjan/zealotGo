package average

import "os"

func ComputeAverage(nums []int) int {
    sum := 0
	for _, num := range nums {
        sum += num
    }

    avg := sum / int(len(nums))

    PersistResult("C:\\tmp\\data", nums, avg)
    return avg
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func PersistResult(fname string, nums []int, avg int) {
    f, err := os.Create(fname)
    check(err)
    defer f.Close()

    n2, err := f.WriteString("numbers: %v, average: %d", nums, avg)
    check(err)
    f.Sync()
}