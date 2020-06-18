package main



import ("fmt"
	  "math"

)


func iterative(x int) {
root:=math.Ceil(math.Sqrt(float64(x))) 
var flag bool=true
 //fmt.Println(int(root))
for i:=2;i<=int(root);i++{

      if(x%i==0){
       	flag=false
       	break
       }
   
}
 if(flag){
 	fmt.Printf("Iterative: "+"Yes %d is a prime",x)
 } else{
  fmt.Printf("Iterative: "+"Not a Prime") 
}


}



func recursive(n,s,root int) bool{

   if(s>root){
    return true
  }
   if(n%s==0){
    return false
  }
    return recursive(n,s+1,root)
 
}

func main(){

var x int
fmt.Scanln(&x)
 iterative(x)
 fmt.Println()
 root:=math.Ceil(math.Sqrt(float64(x)))
  ans:=recursive(x,2,int(root))
 if(ans){
 	fmt.Printf("Recursive: "+"Yes %d is a prime",x)
 }else{
   fmt.Printf("Recursive: "+"Not a Prime") 
   }

}




