package main

import "fmt"


func main(){

//Here I am creating map of indices for arr1 in arr2

arr1:=[]int{2,3,4,5,6,7}
arr2:=[]int{4,5,2,3,7,6}
  

map_:=make(map[int]int)
ans:=make(map[int]int)
 
 for i,val:=range arr2{

 	map_[val]=i
 }

 for i,val:=range arr1{
 	ans[i]=map_[val]
 }

 for key,val :=range ans{
 	fmt.Println(key,val)
 }

}