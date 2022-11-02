package main

import "fmt"

func PrimaSegiEmpat(wide, high, start int) string {

// HLine draws a horizontal line
// func HLine(x1, y, x2 int) {
//     for ; x1 <= x2; x1++ {
//         img.Set(x1, y, col)
//     }
// }

// // VLine draws a veritcal line
// func VLine(x, y1, y2 int) {
//     for ; y1 <= y2; y1++ {
//         img.Set(x, y1, col)
//     }
// }

// // Rect draws a rectangle utilizing HLine() and VLine()
// func Rect(x1, y1, x2, y2 int) {
//     HLine(x1, y1, x2)
//     HLine(x1, y2, x2)
//     VLine(x1, y1, y2)
//     VLine(x2, y1, y2)
// }


//-------------------------------
 
   for h := 1; h <= height; h++ {
      for w := 1; w <= width; w++ {
         if h == 1 || h == height {
            fmt.Print("*")
         } else if w == 1 || w == width {
            fmt.Print("*")
         } else {
            fmt.Print(" ")
         }
      }
      fmt.Println()
   }


func main() {

	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	// /*
	// 	17	19
	// 	23	29
	// 	31	37
	// 	156
	// */
	fmt.Println(PrimaSegiEmpat(5, 2, 1))
	// /*
	// 	2	3	5	7	11
	// 	13	17	19	23	29
	// 	129
	// */
}
