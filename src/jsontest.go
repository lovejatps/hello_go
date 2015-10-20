package main 

import (
    "fmt"
    "encoding/json"
)



func main() {

	group := ColorGroup{
	    ID:     1,
	    Name:   "Reds",
	    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
	// fmt.Fprintf(w, "测误", r.URL.Path[1:])
	
		fmt.Println("测误")
	   
	}else{
	 //fmt.Fprintf(w, string(b), r.URL.Path[1:])
	 	fmt.Println(string(b))
	}



	var s Serverslice
    str := `{"city":"北京","servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
    json.Unmarshal([]byte(str), &s)
  //  fmt.Println(s)

	bb, err := json.Marshal(s)
	if err != nil {
	// fmt.Fprintf(w, "测误", r.URL.Path[1:])
	
		fmt.Println("错误")
	   
	}else{
	 //fmt.Fprintf(w, string(b), r.URL.Path[1:])
	   fmt.Println(string(bb))
	}
	
	
	
	
}

type Server struct {
    ServerName string
    ServerIP   string
}
type Serverslice struct {
	City string
    Servers []Server
}

type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
}
