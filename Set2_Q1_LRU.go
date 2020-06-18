package main

import ("fmt"
	 "strconv"
)
type Node struct{
 key int
 val int
 prev *Node 
 next *Node
}



type LRU struct{
head *Node
tail*Node
m map[int]*Node
capacity int
size int
}

func constructLRU ()LRU {
      h:=&Node{key:-1,val:-1,prev:&Node{},next:&Node{}}
      t:=&Node{key:-1,val:-1,prev:&Node{},next:&Node{}}
      h.next=t
      t.prev=h
   var cache LRU=LRU{head:h,tail:t,m:make(map[int]*Node),capacity:3,size:0}
   return cache

}



func (c *LRU) add(key,val int) {
    
    if(c.m[key]==nil){
        if(c.size==c.capacity){
        	//fmt.Println("lozl")
        	   delete(c.m,c.tail.prev.key)
        	c.remove(c.tail.prev)
        	 temp:=&Node{key:key,val:val,prev:nil,next:nil}
            c.setHead(temp)
            c.m[key]=temp
        } else{
            temp:=&Node{key:key,val:val,prev:nil,next:nil}
            c.setHead(temp)
            c.m[key]=temp
              c.size++
             } 

    } else{
      temp:=c.m[key]
      c.remove(temp)
      temp=&Node{key:key,val:val,prev:nil,next:nil}
            c.setHead(temp)
           c.m[key]=temp


  }


}

func (c *LRU)  get(key int) string{
      if(c.m[key]==nil){
        return "Not Present in Cache"
        }else{
         temp:=c.m[key]
           str:=strconv.Itoa(temp.val)
           return str
        }


}
func (c *LRU)  show() {
	
     var temp *Node=c.head
     for temp!=nil {
     	x:=temp.key
     	if(x==-1||x==0){
         temp=temp.next
         	continue
         }
         y:=temp.val
         
         fmt.Println(x,y)
         temp=temp.next
     }

}

func (c *LRU)  length() int{
	return c.size
}

func (c *LRU) setHead(temp *Node){
    
    var right*Node =c.head.next
     right.prev=temp
     c.head.next=temp
     temp.prev=c.head
     temp.next=right


}

func (c*LRU) remove(temp *Node){
  //fmt.Println("lolz")
    var right*Node =temp.next
	    var left*Node =temp.prev

	    left.next=right
	    right.prev=left
	    temp.next=nil
	    temp.prev=nil

}


 




func main(){

 cache:=constructLRU()
  cache.add(2,4)
  cache.add(4,6)
  cache.add(6,4)
  fmt.Println("Initial Cache")
  cache.show()
  cache.add(7,9)
  cache.add(4,3)
  cache.add(10,5)
  fmt.Println("Initial Cache")
  fmt.Println("Cache after additions")
  cache.show()
  fmt.Println(cache.length())
  fmt.Println(cache.get(7))
  cache.add(12,6)
  fmt.Println(cache.get(5))
   cache.show()
  fmt.Println(len(cache.m))




}