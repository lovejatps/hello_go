package main

import  (
//"unicode/utf8"
"fmt"
"runtime"
"strings"
)
func main() {

	const(
		Sunday= iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	
	os_ := runtime.GOOS
	
	if strings.EqualFold(os_,"windows"){
		fmt.Println("windows")
	}
	

	
	list := [] string {"teset1","teset2","ccc"}
	for i:=0; i < 3; i++{
		a := list[i]
		switch a {
		case "test1":
			fmt.Println("teset1")
		case "test2","teset3":
			fmt.Println("test_N")
		default:
			fmt.Println("default")
	
	  }
   }
   
   
   for i:=0 ; i<10 ; i++{
   		fmt.Println("I:",i)
   }


	arraydata := [] int {0,1,2,3,4,5,6,7,8,9,10}
	for _, v := range arraydata{
		fmt.Println(v)	
	}

	fmt.Println("------------------------")	
	ForGoto()
	fmt.Println("------------------------")	
	
	Revert("A test 测试中文")
	
	fmt.Println(Reverse("A test 测试中文"))
	

}


	func ForGoto(){
		i:=0
		I:
		fmt.Println(i)
		i++
		if(i<10){
			goto I
		} 
	}
	
func Revert(vals string){
 	var revelts string
 	for _, val:=range vals{
 		revelts = string(val)+revelts
 	}
	fmt.Println(revelts) 
 }
 
 func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

//func Reverses(s string){
//    o := make([]int, utf8.RuneCountInString(s));
//    for i, j := 0, len(o)-1; i < j; i, j = i+1, j-1 {
//        o[i], o[j] = o[j], o[i]
//    }
//    
//  
//}
