package main
import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
//init
a := make([][]uint8,dy)
for i:=0; i < dy; i++{
a[i] = make ([]uint8,dx)
}

    for k:=0; k < dy; k = k + 5{  
        for h:=0; h < dx ;h = h +5 {
        a[k][h] = 255
    	}
    }
    
return a  
}

func main() {
pic.Show(Pic)
}
//てんてん
