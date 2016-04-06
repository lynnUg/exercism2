// Build constraint.  It's like when you buy something new and you find a
// tag that says "Inspected by #12."  Nice to know it was inspected, but
// you don't need the tag.  Take it off.
// +build !example

// Package clause.
package gigasecond
import ( 
"time"
"math"
)
// Constant declaration.
const TestVersion = 2 // find the value in gigasecond_test.go

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time{
return t.Add(time.Duration(math.Pow(10,18)))
}

// API variable.  It needs a type at least.  A value would be nice too.
// See gigasecond_test.go.
var Birthday time.Time

// Reviewers don't think much of stub comments.  Replace or remove.
