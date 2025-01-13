package wo_info


//z := 1 // not possible at package level
func declares() {
 var a int = 1
 var c = 1
 d := 1
 var e int
 var (
     // all the same things it could already do
 )
 var m, n int = 1, 2
 var q, r = 1, 2
 s, t := 1, 2
 var u, v int
 
 e = 2
 e, y := 2, 3
 if d == 1 {
     d := 2
     d := "other type"
     var d []File
     m, n := 3, 4
 }
 // at this point,
 // d == 1
 // m, n == 1, 2 
 
 // note: e was already declared (we wouldn't need this note if shadowing were more explicit)
 e := 2       // not possible
 e, _ := 2, 3 // not possible
 
 b int := 1 // not possible
 f int // not possible (unlike C style int f;)
 var X // not possible
o, p int := 1, 2 // not possible
w, x int // not possible
 
 fmt.Println(a, b, ... x, y) // haha
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
