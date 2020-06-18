package main

import ("fmt"
   "strconv"

)
type Node struct{
next *Node
key int
val int
move *Node

}



type LinkedHashMap struct{

arr []*Node
 head *Node
 tail *Node

 size int
 cap int

}

func(m*LinkedHashMap) printOrder(){

 temp:=m.head
 for temp!=nil{
     a:=strconv.Itoa(temp.key)
     b:=strconv.Itoa(temp.val)
 	fmt.Println(a+" "+b)
   temp=temp.move

 }





}

func (m *LinkedHashMap) contains(key int) bool{

  index:=hashcode(key)
  
  temp:=m.arr[index]
  for temp!=nil{
     
     if(key==temp.key){
     	return true
     }   
     temp=temp.next

  }

   return false

}


func (m *LinkedHashMap) get(key int) string{

  index:=hashcode(key)
  
  temp:= m.arr[index]
  
   for temp!=nil{
     
     if(key==temp.key){
     	x:=strconv.Itoa(temp.val)
     	//fmt.Println(temp.val)
     	return string(x)
     }   
       temp=temp.next
  }
  return "Not found"

 

   }
 


func(m *LinkedHashMap) put(key ,val int){


 root:=&Node{key:key,val:val}

  index:=hashcode(key)

  if(m.contains(key)){
  	temp:= m.arr[index]
  
   for temp.key!=key{
     
       temp=temp.next
  }

  temp.val=val



  }else{

 if(m.arr[index]==nil){
    m.arr[index]=root
 }else{

   root.next=m.arr[index]
   m.arr[index]=root
 }

  
  if(m.head==nil){
    m.head=root
    m.tail=root
  }else  {
      m.tail.move=root
       m.tail=(m.tail.move)
    }
    m.size++
   } 

  

}


func createMap() LinkedHashMap{
 
 var map_ LinkedHashMap
  map_.arr=make([]*Node,100,100)
 map_.size=0
 map_.cap=100
 
 return map_

}


func  hashcode(key int) int{

 return (key)%100
}
   



func main(){
map_:=createMap()
map_.put(5,3)
map_.put(10,2)
map_.put(16,4)
map_.put(16,32)
map_.put(14,58)
map_.put(432,53)
map_.put(1,3)

fmt.Println(map_.get(10))
fmt.Println(map_.get(16))
fmt.Println(map_.get(6))
map_.printOrder()
fmt.Printf("The size of the LinkedHashMap is %d",map_.size)


}


