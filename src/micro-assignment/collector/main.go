package collector

import (
    "./awscollectors/cloudtrail.go"
    "fmt"
)


func main() {
    fmt.Println("trying package inclusion...")

    c := CloudTrailCollector{}
    //result := c.defaultSchemaValidator.Validate()
    //fmt.Printf("validatino:" + result)
}