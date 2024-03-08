package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	A := 0.0
	B := 0.0
	for {
		z := [1760]float64{}
		b := [1760]byte{}
		for j := 0; j < 22; j++ {
			for i := 0; i < 80; i++ {
				b[j*80+i] = ' '
			}
		}
		for j := 0.0; j < 6.28; j += 0.07 {
			for i := 0.0; i < 6.28; i += 0.02 {
				c := math.Sin(i)
				d := math.Cos(j)
				e := math.Sin(A)
				f := math.Sin(j)
				g := math.Cos(A)
				h := d + 2
				D := 1 / (c*h*e + f*g + 5)
				l := math.Cos(i)
				m := math.Cos(B)
				n := math.Sin(B)
				t := c*h*g - f*e
				x := int(40 + 30*D*(l*h*m-t*n))
				y := int(12 + 15*D*(l*h*n+t*m))
				if 0 <= y && y < 22 && 0 <= x && x < 80 {
					o := x + 80*y
					N := int(8 * ((f*e-c*d*g)*m - c*d*e - f*g - l*d*n))
					if D > z[o] {
						z[o] = D
						if 0 <= N && N < len(".,-~:;=!*#$@") {
							b[o] = ".,-~:;=!*#$@"[N]
						}
					}
				}
			}
		}
		fmt.Print("\x1b[H")
		for k := 0; k < len(b); k++ {
			fmt.Printf("%c", b[k])
			if k%80 == 0 {
				fmt.Println()
			}
		}
		time.Sleep(time.Millisecond / 60)
		A += 0.04
		B += 0.02
	}
}
