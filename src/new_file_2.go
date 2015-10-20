package main 

import (
	"encoding/json"
    "net/http"
    "fmt"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	group := ColorGroup{
	    ID:     1,
	    Name:   "Reds",
	    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	
//	var jsonBlob = []byte(`{"city":"北京","animalobj":[{"Name": "Platypus", "Order": "huxn"},{"Name": "Quoll", "Order": "Dasyuromorphia"} ]}`)

//	var jsonBlob = []byte(`[{"Name": "Platypus", "Order": "huxn"},{"Name": "Quoll", "Order": "Dasyuromorphia"} ]`)
	
	
	var jsonBlob = []byte(`{"city":"北京","dataobj":{"Name": "Platypus", "Order": "huxn"}}`)
	
	var animals Datas  
	err1 := json.Unmarshal(jsonBlob, &animals)  
	if err1 != nil {  
    	fmt.Println("error:", err1)  
	}  
	
	fmt.Printf(animals.City)  
	
	
	b, err := json.Marshal(group)
	if err != nil {
	// fmt.Fprintf(w, "测误", r.URL.Path[1:])
	
		w.Write([] byte("错误"))
	   
	}else{
	 //fmt.Fprintf(w, string(b), r.URL.Path[1:])
	   w.Write(b)
	}
	
    
}

func main() {
    http.HandleFunc("/test", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
}


type Datas struct {
	City string
	Dataobj Animal

}

type Animal struct {  
    Name  string  
    Order string  
}

	type Marshaler interface {  
	    MarshalJSON() ([]byte, error)  
	}  
	type Unmarshaler interface {  
	    UnmarshalJSON([]byte) error  
	}  
