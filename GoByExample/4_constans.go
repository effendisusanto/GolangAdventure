package main

import "famt"
import "math"

const s string = "constant"

func main(){
	famt.Println(s)

	const n=500000000

	const d=3e20 /n

	famt.Println(d)

	famt.Println(int64(d))

	famt.Println(math.Sin(n))
}
