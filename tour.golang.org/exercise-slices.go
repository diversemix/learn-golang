package main

import ( 
	//"fmt" 
	"golang.org/x/tour/pic"
)

func f(x, y int) int {
 	return x*y
}

func Pic(dx, dy int) [][]uint8 {
	retval := make([][]uint8, dy)
	for i:=0; i<dy; i++ {
		// fmt.Println("i", i)
		retval[i] = make([]uint8, dx)
		for j:=0; j<dx; j++ {
			// fmt.Println("j", j)
			retval[i][j] = uint8(f(i,j))
		}
	}
	return retval
}

func main() {
	pic.Show(Pic)
}

