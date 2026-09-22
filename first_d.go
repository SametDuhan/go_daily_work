package main 

import "fmt"
const PI = 3.14
func main(){
  fmt.Println("Go is working well")

//This is a comment row
var name string = "Alia"  // type is string
var surname = "Atreides"  // type is inferred
a := 14 

const CONSTYOE type = value
var arr1 = [3]int{1,2,3}


fmt.Println(name)
fmt.Println(surname)
fmt.Println(a)

myslice1 := []int{}
fmt.Println(len(myslice1))
fmt.Println(cap(myslice1))
fmt.Println(myslice1)

myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
fmt.Println(len(myslice2))
fmt.Println(cap(myslice2))
fmt.Println(myslice2)

fmt.print("Enter Your Name : ")
fmt.scan(&name)

fmt.print("enter your age : ")
fmt.scan(&age)

fmt.printf("hello %s , you are %d years old",name,age)

package main

import "fmt"

func main() {
    puan := 75

    if puan >= 90 {
        fmt.Println("Harf Notu: AA")
    } else if puan >= 80 {
        fmt.Println("Harf Notu: BA")
    } else if puan >= 70 {
        fmt.Println("Harf Notu: BB")
    } else {
        fmt.Println("Daha çok çalışmalısın.")
    }
}




  
}
