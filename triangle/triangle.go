//+build !example

package triangle

// Code this function.
func KindFromSides(a, b, c float64) Kind {
if (a+b)>c && (b+c)>a && (a+c)>b {
if a==b && b==c  && a==c {
return 3
}

if a==b || b==c || a==c {
return 2
}
return 0
}

return -1
}

// Notice it returns this type.  Pick something suitable.
type Kind int

// Pick values for the following identifiers used by the test program.
var Equ Kind =3 // equilateral
var Iso Kind =2// isosceles
var Sca Kind =0// scalene
var NaT Kind =-1// not a triangle

// Organize your code for readability.
