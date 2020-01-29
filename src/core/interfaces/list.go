package main

import "fmt"

func main(){
    test_list_operations()
}


type Node struct {
    data int
    next *Node
}

func test_list_operations(){
    var root *Node;

    // Create the list with single node
    root = root.create(10)
    root.display()
    fmt.Println(root.length())
    fmt.Println("---------------------------")
    // insert three times
    root.insert(20)
    root.insert(200)
    root.insert(2000)
    root.insert(20000)
    root.insert(200000)
    root.insert(20000000)
    root.display()
    fmt.Println(root.length())
    fmt.Println("---------------------------")
    // delete
    root.delete(&root, 0)
    root.display()
    fmt.Println(root.length())
    root.delete(&root, 2)
    root.display()
    fmt.Println(root.length())
    root.delete(&root, 4)
    root.display()
    fmt.Println(root.length())
    fmt.Println("---------------------------")
}


func (n *Node) create(val int) (*Node){
    var root *Node = new(Node)
    root.data = val
    root.next = nil
    return root
}

func (root *Node) display(){
    if (root != nil) {
        curr := root
        for ; curr.next != nil; curr = curr.next {
            fmt.Println(curr.data)
        }
        fmt.Println(curr.data)
    }
}

func (root *Node) length() int {
     length := 0
     if (root != nil) {
        curr := root
        length = length + 1
        for ; curr.next != nil; curr = curr.next {
            length = length + 1
        }
    }
    return length
}

func (root *Node) insert(val int) bool {
    if (root != nil) {
        curr := root
        for ; curr.next != nil; {
             curr = curr.next
        }
        var new_node *Node = new(Node)
        new_node.next = nil
        new_node.data = val
        curr.next = new_node
        return true
    } else {
        return false
    }
}


func (root *Node) delete(root_ptr **Node, pos int) bool {
    if (root != nil) {
        if pos >= root.length() {
            return false
        }

        if pos == 0 {
            root = root.next
            *root_ptr = root
            fmt.Println("-")
            root.display()
        } else {
            p1 := root
            p2 := root.next
            for i:=1; i<pos; i++ {
                p1 = p1.next
                p2 = p1.next
            }
            p1.next = p2.next
        }
        return true
    }

    return false
}