package main

import (
    "net/url"
    "fmt"
    "net"
)

// Little bit about url...
// url can have following elements:
//      1. host name
//      2. port
//      3. query params:
//      4. scheme (What is it?) - https, http, postgres etc are scheme. :-) Learnt something more.
//      5. authentication info
func main() {
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery) // Got the piece '/path'
    m, _ := url.ParseQuery(u.RawQuery) // Got the map of keyvalue:pair
    fmt.Println(m)
    fmt.Println(m["k"][0])

    // Repeat with other url
    do_for_google()
}


func do_for_google(){
    s := "http://www.moneycontrol.com/bestportfolio/wealth-management-tool/investments#port_top"
//            |               |         -----------------------------------------------
//            |               |                       |                               ---------
//            |               |                       |                                     |
//            |               |                       |                                     |_______ fragment
//            |___ scehme     |                       |
//                            |_____ ?                |
//                                                    |
//                                                    |____________ path
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println("----------------------------------------------------")
    fmt.Println(s)
    fmt.Println(u.Scheme)

    fmt.Println("path:", u.Path)
    fmt.Println("fragment:", u.Fragment)
    fmt.Println("raw query:", u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery) // Got the map of keyvalue:pair
    fmt.Println(m)


    //if len(u.User()) != 0 {
    //        fmt.Println("No user name specified")
    //}

    // It pancked? :-)

    // fmt.Println(u.User)
    // fmt.Println(u.User.Username())
    // p, _ := u.User.Password()
    // fmt.Println(p)
}