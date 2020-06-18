package main

import ("fmt"
)


func main(){


// var int n

a:=[10]int{1,0,2,0,13,14,0,0,1,5}

low:=0
mid:=0

for mid<10 {
  if(a[mid]!=0) {
    temp:=a[low]
    a[low]=a[mid]
    a[mid]=temp
     mid=mid+1
     low=low+1
     
  } else {
    mid=mid+1
  }

}

fmt.Println(a)




}
