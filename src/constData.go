package main 

import (

    "fmt"
    "strings"
    "strconv"
    "time"
)

var maps map[string] Mapobj 


type Mapobj struct{
	
	id int
	name string
	age int

}

func main() {

	maps = make(map[string] Mapobj)

	const PI ="huxn"
	var timed string 
	str := "2015-07-01T18:49:20+08:00"
	i := strings.Index(str,"T")
	timed = string((str)[i+1:i+3])
	times,err :=strconv.Atoi(timed)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(times+1)
	}
	//s := strings.SplitN("2015-07-01T16:49:20+08:00",":",2)
    fmt.Println("len:",len(str),"i",i,"time",timed)

	var timedate string
	 t := time.Now().Unix()
	 timedate = time.Unix(t, 0).String()
	 fmt.Println(timedate)
	 fmt.Println(strings.Fields(timedate)[1])
	 
	 if isTime(str){
	 	fmt.Println("时间相等不用更新")
	 }else{
	 	fmt.Println("时间不相等需要用更新")
	 }
	 
	maps["111"]= Mapobj{100,"huxn",33}
	
	mapdata, ok := maps["111"]
	
	if ok{
		fmt.Println("map:",mapdata.name)
	}
	 
}

func isTime(s string) bool{
	str1 := strings.Split(s,":")
	str1 = strings.Split(str1[0],"T")
	date1 := str1[0]
	t := time.Now().Unix()
	timestr := time.Unix(t, 0).String()
	date2 := strings.Fields(timestr)[0]
	if strings.EqualFold(date1,date2){
		time_h,err := strconv.Atoi(str1[1])
		if err != nil{
			fmt.Println(err)
			return false
		}else{
			current_hour,err1 := strconv.Atoi(strings.Split(strings.Fields(timestr)[1],":")[0])
			if err1 != nil{
				fmt.Println(err1)
				return false
			}else{
				if current_hour == time_h{
					return true
				}else{
					return false	
				}
			}
			
		}
		
	}else{
		return false	
	}
	
	

	return false
}


