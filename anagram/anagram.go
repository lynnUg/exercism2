package anagram

import (
"strings"
"sort"
)
<<<<<<< HEAD
func Count(s string) a []string {
  for i,_ := range s {
  	z:=string(s[i])
  	if _,ok:=a[] ; ok {
  		a[z]+=1
  	} else {
  		a[z]=1
  	}
  }
  return 
}
func Detect(subject string , candidates []string ) output []string {
	subject_count:=Count(subject)
	
    for i,_ := range candidates {
    	is_anagram:=false
    	cand_count:=Count(candidates[i])
    	for i,_:= range subject {
    		a:=string(subject[i])
    		if subject_count[a] == cand_count[a] {
    			is_anagram = true

    		} else {
    			is_anagram = false
    			break
    		}
    	}

    	if is_anagram {
    		if len(subject) == candidates[i] {
                output= append(output,candidates[i] )

    		}
    	}
    }
	return 
}
