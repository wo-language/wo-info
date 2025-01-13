package _go

import "fmt"

// z := 1 // not possible at package level
func declares() {
	var a int = 1
	var c = 1
	d := 1
	var e int
	var m, n int = 1, 2
	var q, r = 1, 2
	s, t := 1, 2
	var u, v int
	e = 2
	e, n = 0, 1
	e, y := 2, 3
	{
		d := "other type"
		e = 3
		var (
			a int = 3
			c = 2
			//d := "other type"
			e int16
			m, n int = 1, 2
			q, r = 1, 2
			//s, t := 1, 2
			u, v int
			f = 2
			//e, y := 2, 3
		)
		fmt.Println(a, c, d, e, f, m, n, s, t, u, v, y, q, r)
	}
	// at this point, d == 1

	// note: e was already declared (we wouldn't need this note if shadowing were more explicit)
	//e := 2       // not possible
	//e, _ := 2, 3 // not possible

	//b int := 1 // not possible
	//f int // not possible (unlike C style int f;)
	//var X // not possible
	//o, p int := 1, 2 // not possible
	//w, x int // not possible

	fmt.Println(a, c, d, e, m, n, q, r, s, t, u, v, x, y)
}

func constDeclares() {
	const a int = 1
	const c = 1
	//const d := 1
	//const e int
	const m, n int = 1, 2
	const q, r = 1, 2
	s, t := 1, 2
	//const u, v int

	//e = 2
	//e, n = 0, 1
	//e, y := 2, 3
	{
		d := "other type"
		const (
			a int = 3
			c = 2
			//d := "other type"
			//e int16
			m, n int = 1, 2
			q, r = 1, 2
			//s, t := 1, 2
			//u, v int
			f = 2
			//e, y := 2, 3
		)
		fmt.Println(a, c, d, f, m, n, s, t, q, r)
	}

	fmt.Println(a, c, m, n, q, r, s, t, x)
}

//x int = 3
//var a = 8
//const X int = 9
//const b = X
//x = b
//x := b // not allowed; doesn't shadow, use var x = for init
//{ // shadow
//  x int := 0
//  var a := "other type"
//  const X int := 99
//  const b := a
//  x = b  // assign
//  x := 1 // not allowed
//}
//var ( z, _ (int, error) = count() ) // has types mixed in, or even consts
//const ( h string = "永", e error = nil )

func declaresWo() {
	//var a int = 1
	var c = 1
	//d := 1
	//var e int
	//var m, n int = 1, 2
    m, n int = 1, 2
    // OR: o, p (int, uint) = 1, 2
	var q, r = 1, 2
	//s, t := 1, 2
	//var u, v int
	c = 2
	//e, n = 0, 1
	//e, y := 2, 3
	{
		var (
			//a int := 3
			c := 2
			//d := "other type"
			//e int16
			//m, n int = 1, 2
			q, r := 1, 2
			//s, t := 1, 2
			//u, v int
			f = 2
			//e, y := 2, 3
		)
		fmt.Println(a, c, d, f, m, n, s, t, q, r)
	}
	// at this point, d == 1

	// note: e was already declared (we wouldn't need this note if shadowing were more explicit)
	//e := 2       // not possible
	//e, _ := 2, 3 // not possible

	//b int := 1 // not possible
	//f int // not possible (unlike C style int f;)
	//var X // not possible
	//o, p int := 1, 2 // not possible
	//w, x int // not possible

	fmt.Println( c, q, r, s, t, x)
}

func constDeclaresWo() {
	const a int = 1
	const c = 1 // const var
	//d := 1
	//const e int
	const m, n int = 1, 2
	const q, r = 1, 2
	s, t := 1, 2
	//const u, v int

	//e = 2
	//e, y := 2, 3
	{
		const (
			a string := 3
			c := 2
			//d := "other type"
			//e int16
			m, n int := 1, 2
			q, r := 1, 2
			//s, t := 1, 2
			//u, v int
			f = 2
			//e, y := 2, 3
		)
		fmt.Println(a, c, f, m, n, s, t, q, r)
	}

	fmt.Println(a, c, m, n, q, r, s, t, x)
}
