package main

import (
    "reflect"
    "fmt"
    "strings"
)

type Collector interface {
	GetName() string
	Collect() []map[string]interface{}
}

type VMCollector struct {
}

func (collector *VMCollector) Collect() (resources []map[string]interface{}) {
    return
}

// This everyone implements
func (collector *VMCollector) GetName() (string, string, string) {
    return utilGetName(fmt.Sprintf("%s", reflect.TypeOf(collector)))
}


// This goes in utility code
func utilGetName(fqn string) (string, string, string){
    fqn = strings.Replace(fqn, "*", "", -1)
    parsedFqn := strings.Split(fqn, ".")
    package_name, collector_name := parsedFqn[0], parsedFqn[1]
    return fqn, package_name, collector_name
}


func main() {
    vmc := new(VMCollector)
    a, b, c := vmc.GetName()
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}