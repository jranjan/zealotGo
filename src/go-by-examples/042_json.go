package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// This is an interesting area. We can convert built-in as well as custom data
// to json and vice versa. Points to be noted are:
//      1. Need to call Marshal() or Unmarshal()
//      2. If you are doing unmarshling to custom types then you get type safety. No need to apply cast operator
//      3. If you read generic data then you might need to apply cast. See example for better understanding.
//      4. While marshaling built-in or custom data type to json, you might specify custom key name

type Response1 struct {
    Page   int
    Fruits []string
}
type Response2 struct {
    Page   int      `json:"page_details"` // Interesting, it allows to pick the name of our choice.
    Fruits []string `json:"contains_fruits"`
}

func main() {
    encode()
    decode()
}

func encode(){
    // json.Marshal() will encode the data type to json
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))
}


func decode(){
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    num := dat["num"].(float64)
    fmt.Println(num)

    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    str := `{"page_details": 1, "contains_fruits": ["apple", "peach"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}