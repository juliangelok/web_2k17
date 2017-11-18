package main
import "strings"
var degr=1;
func PowerGenerator(a int) (func() int) { 
return func() (ret int) {
degr = degr*a
ret = degr
return
}
}
func RemoveEven(b []int) (chetn []int) { 
i := 0
chetn = make([]int, 0)
for i < len(b) {
if b[i]%2==1 {
chetn = append(chetn, b[i])
}
i++
}
return
}
func DifferentWordsCount(c string) (kolvo int) { 
dif := strings.ToLower(c)
d:= strings.FieldsFunc(dif,Prin)	
kolvo = 0;
i:= 0;
for i < len(d) {
if d[i]!="1" {
kolvo++
j:=i+1
for j < len (d) {
if d[i]==d[j] {
d[j]="1"
}
j++
}
d[i]="1"
}
i++
}										
return
}
func Prin(x rune) bool {
if x<='z' && x>='a' {
return false
} else {
return true
}
}

