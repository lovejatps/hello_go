package main 

import( 
//	iconv "github.com/djimenez/iconv-go"
	"fmt"
	"net/http"
	"io/ioutil"

)
func main() {

const Hello = "hello1"

// fmt.Println("hello2")
 fmt.Println(Hello)
 
 const (     a = 1 << iota   // a == 1 (iota在每个const开头被重设为0)    
 			 b = 1 << iota   // b == 2   
 		     c = 1 << iota   // c == 4 
 		     d = 1 << iota   // d == 8 
 		      )   
 fmt.Println(d)	  
 
 //如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式
  const (    aa = 1 << iota   // a == 1 (iota在每个const开头被重设为0)    
 			 bb
 		     cc
 		     dd
 		      )   
 fmt.Println(dd)	
 
 var db=1000    
 db2:=2000  
 db3:= 124|3
 		      
 fmt.Println(db,db2,db3)
 
 var db4 float32 
 db4 = 90.0003
 db5:=90.0004
 fmt.Println("test:",db4,db5)
 
 var  name string
 age:="33"
 name ="huxn"+age

 fmt.Println(name)
 
 //--------方案一　－－－－－－－－－－－－－－－－－－－
 	response,_:=http.Get("http://www.baidu.com")
 	defer response.Body.Close()
 	body,_:=ioutil.ReadAll(response.Body)
 	fmt.Println(string(body))
 	
 	
  //--------方案二－－－－－－－－－－－－－－－－－
  
  client :=&http.Client{}
  reqest,_:=http.NewRequest("GET","http://www.baidu.com",nil)
  reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
  reqest.Header.Set("Accept-Charset","utf-8,GBK;q=0.7,*;q=0.3")
  reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
  reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
  reqest.Header.Set("Cache-Control","max-age=0")
  reqest.Header.Set("Connection","keep-alive")
  responses,_:=client.Do(reqest)
  if(responses.StatusCode==200){
 // 	body,_:=ioutil.ReadAll(responses.Body)
//  	out := make([]byte, len(input))
//    out = out[:]
//    iconv.Convert(input, out, "gb2312", "utf-8")
//  	ioutil.WriteFile("out.html", out, 0644)
  //	fmt.Println(bodystr)
  	}
  
  
  
  
}

