package main
/*
    Maps are keys and values.
    in Python there are know as Dict
    in Js there are know as Objects
*/

import("fmt")


func main(){

    inMap:=make(map[string]int)

    inMap["one"] = 1
    inMap["two"] = 2
    inMap["three"] = 3
    inMap["four"] = 4
    inMap["five"] = 5

    for key, value :=  range inMap {
        fmt.Println("This is a key => ",key, " And this is the value ->",value)
    }

    // To test if some element is in a map
    el, ok:=inMap["four"]

    if ok {
    fmt.Println(el," is in map")  
    }else{
    fmt.Println(el," is not in map")  
    }

    // To delete in map
    delete(inMap,"four")


    // To change some value in a element
    inMap["two"]=4

    // To print all elements in a map
    fmt.Println("Oi, este é o meu map => ", inMap)
}

