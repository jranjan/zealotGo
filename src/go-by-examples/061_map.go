package main

import (
    "fmt"
    "errors"
    "math/rand"
)

type collectAttributes struct {
    name string
    constructor func(hint int, name string) interface{}
}

var collectorMap = map[string]collectAttributes{
    "BlobStorageAccountCollectorName" : {"Azure BlobStorageAccountCollector", NewBlobStorageAccountCollector},
    "DiskCollectorName" : {"Azure DiskCollector", NewDiskCollector},
    "MySQLCollectorName" : {"Azure MySQLCollector", NewVMCollector},
    "SnapshotCollectorName" : {"Azure SnapshotCollector", NewVirtualNetCollector},
    "SQLCollectorName" : {"Azure SQLCollector", NewSnapshotCollector},
    "VMCollectorName" : {"Azure VMCollector", NewSQLCollector},
    "VirtualNetCollectorName" : {"Azure VirtualNetCollector", NewMySQLCollector},
}




// NewBlobStorageAccountCollector returns an instance of BlobStorageAccountCollector
func NewBlobStorageAccountCollector(hint int, name string) interface{} {
	fmt.Println("Created collector: %s with attributes", name, hint)
	return nil
}

// NewBlobStorageAccountCollector returns an instance of BlobStorageAccountCollector
func NewDiskCollector(hint int, name string) interface{} {
	fmt.Println("Created collector: %s with attributes", name, hint)
	return nil
}

// NewBlobStorageAccountCollector returns an instance of BlobStorageAccountCollector
func NewVMCollector(hint int, name string) interface{} {
	fmt.Println("Created collector: %s with attributes", name, hint)
	return nil
}

// NewBlobStorageAccountCollector returns an instance of BlobStorageAccountCollector
func NewVirtualNetCollector(hint int, name string) interface{} {
	fmt.Println("Created collector: %s with attributes", name, hint)
	return nil
}

// NewBlobStorageAccountCollector returns an instance of BlobStorageAccountCollector
func NewSnapshotCollector(hint int, name string) interface{} {
	fmt.Println("Created collector: %s with attributes", name, hint)
	return nil
}

// NewBlobStorageAccountCollector returns an instance of BlobStorageAccountCollector
func NewSQLCollector(hint int, name string) interface{} {
	fmt.Println("Created collector: %s with attributes", name, hint)
	return nil
}

// NewBlobStorageAccountCollector returns an instance of BlobStorageAccountCollector
func NewMySQLCollector(hint int, name string) interface{} {
	fmt.Println("Created collector: %s with attributes", name, hint)
	return nil
}


// newCollectorsGeneral implements a general provider level collector for Azure
var newCollectorsGeneral = func() (collectors []interface{}, err error) {
	// Loop through each collector function and add to the collectors list
	for _, collectorattributes := range collectorMap {
		collectors = append(collectors, collectorattributes.constructor(rand.Intn(900), collectorattributes.name))
	}

	if len(collectors) == 0 {
		return collectors, errors.New("No collectors found for Azure provider")
	}
	return
}

// NewAzureCollectors returns a map of collectors for azure provider
func NewAzureCollectors() (collectors []interface{}, err error) {
	newCollectors, err := newCollectorsGeneral()
	if err != nil {
		return nil, err
	}
	collectors = append(collectors, newCollectors...)
	return
}


func main() {
    fmt.Println("Hello world")
    NewAzureCollectors()
	return
}