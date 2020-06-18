
package main
import "fmt"
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
var post int=-1

func main(){
a:=[]int{9,3,15,20,7}
b:=[]int{9,15,7,20,3}

root:=buildTree(a,b)

inorder(root)

}

func preoder(root *TreeNode ){
  if(root==nil){
    fmt.Printf("null ")
    return 
  }
   fmt.Printf("%d ",root.Val)
  
  inorder(root.Left)
   inorder(root.Right)


}


func buildTree(inorder []int, postorder []int) *TreeNode {
    var n int=len(inorder)
    post=n-1
    
 root:=build(inorder,0,n-1,postorder)
    return root
        
    }  
    

func build(inorder []int, s int, e int,postorder []int) *TreeNode{
    
    if(s>e){
     return nil
    }   
    if(s==e){
        post--
    return &TreeNode{Val:inorder[s]}
    }
    var val int=postorder[post]
    post--
    
    loc:=-1
    for i:=s;i<=e;i++ {
        if(inorder[i]==val){
            loc=i
            break
        }  
    } 
          
        
    root:=&TreeNode{Val:val,Left:nil,Right:nil}
        
        root.Right=build(inorder,loc+1,e,postorder)
        root.Left=build(inorder,s,loc-1,postorder)        
        return root

    
}