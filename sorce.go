package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
//init
    a := make([][]uint8,dy)
    for i:=0; i < dy; i++{
        a[i] = make ([]uint8,dx)
    }
    
    return a
}

func main() {
    pic.Show(Pic)
}
