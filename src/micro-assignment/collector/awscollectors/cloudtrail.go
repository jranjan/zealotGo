package awscollectors

import (
	"fmt"
	"math/rand"
    "../validator/validator.go"
)

type CloudTrailCollector struct {
	defaultSchemaValidator DefaultSchemaValidator
}

// Collect gets CloudTrail resources
func (collector *CloudTrailCollector) Collect() int {
    n := rand.Intn(100)
    fmt.Print(n, ",")
    return n
}