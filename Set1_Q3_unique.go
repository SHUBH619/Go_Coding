package main

import (

	"fmt"
      
)

func check(c byte) bool{
         ch:=c
    if((ch>=65&&ch<=92)||(ch>=97&&ch<=122)||(ch>=48&&ch<=57)){
        return false
    } 
    return true  

}


func main(){
   var str string="Hello Team ZettaBytes.My Name is Shubhang.Verma.I hope you are doing well,My code is wrriten in go,I love programming "
    var n int=len(str)
    // for _,ch:=range "Hello World" {
    // 	fmt.Println(string(ch))
    //  }
    var word string=""
     map_:=make(map[string]int)
  for i:=0;i<n;i++ {
      var ch byte=str[i]
       //fmt.Println(rune(ch))
       if(string(ch)==" "||check(str[i])){
        //fmt.Println(ch)
         if(map_[word]==0){
           map_[word]=1
          } else {
           map_[word]=map_[word]+1
          }

         word=""
       } else{
 
          word+=string(str[i])
       }


  }

fmt.Println(len(map_),map_)



}