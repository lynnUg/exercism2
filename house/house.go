package house


func Embed(s , p string) string {

return s+" "+p
}

func Verse(s string , x []string , p string) string{

if len(x)==0 {
return Embed(s,p)
}
return Verse(s ,x[:len(x)-1],Embed(x[len(x)-1],p))
}

func Song() string {
start:=""
verse :=[] string {
"that belonged to the farmer sowing his corn",
"that kept the rooster that crowed in the morn",
"that woke the priest all shaven and shorn",
"that married the man all tattered and torn",
"that kissed the maiden all forlorn",
"that milked the cow with the crumpled horn",
"that tossed the dog",
"that worried the cat",
"that killed the rat",
"that ate the malt",
}

return ""
}
