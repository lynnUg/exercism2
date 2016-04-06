// +build !example

package foodchain
import (
"fmt"
"strings"
)
const TestVersion = 1

var object = []string{"fly","spider","bird","cat","dog","goat","cow","horse"}

var verseStart = "I know an old lady who swallowed a %v.\n"
var verseMid =[]string {
  "",
  "It wriggled and jiggled and tickled inside her.\n",
 "How absurd to swallow a %v!\n",
  "Imagine that, to swallow a %v!\n",
 "What a hog, to swallow a %v!\n",
  "Just opened her throat and swallowed a %v!\n",
 "I don't know how she swallowed a %v!\n",
  "She's dead, of course!"}
var verseRepeat= "She swallowed the %v to catch the %v.\n"

var verseLast= "I don't know why she swallowed the %v. Perhaps she'll die."

const wriggle =" that wriggled and jiggled and tickled inside her.\n"

func Verse(n int) string{
  if n < 1 || n > 8 {
    return "I know nothing"
  }
  if n==1 {
    return fmt.Sprintf(verseStart,object[n-1])+ fmt.Sprintf(verseLast,object[n-1])
  }
  if n==8 {
    return fmt.Sprintf(verseStart,object[n-1])+ verseMid[n-1]

  }
  var fullVerse string

  for i:=1 ;i<n;i++ {
    if i==2{
      fullVerse = fmt.Sprintf("She swallowed the %v to catch the %v",object[i],object[i-1])+ wriggle +fullVerse
    } else {
      fullVerse = fmt.Sprintf(verseRepeat,object[i],object[i-1])+fullVerse
    }

  }
  if n== 2{
   fullVerse = fmt.Sprintf(verseStart,object[n-1]) +verseMid[n-1] +fullVerse
  } else {
     fullVerse = fmt.Sprintf(verseStart,object[n-1]) +fmt.Sprintf(verseMid[n-1],object[n-1])+fullVerse
  }
  fullVerse += fmt.Sprintf(verseLast,object[0])

  return fullVerse

}

func Verses (x,y int) string {
  return Verse(x)+"\n\n"+Verse(y)
}

func Song() string  {
output :=""
for i:=1 ; i<9 ;i++ {
output+=Verse(i)+"\n\n"
}
return strings.TrimSpace(output)
}
