package main

import "fmt"

func main(){
    
     a:=[]int{1,2,3,3,5,6,341}
     b:=[]int{1,4,5,5,6}
     
     c:= merge(a,len(a),b,len(b))
      fmt.Println(c)


}


func merge(nums1 []int, m int, nums2 []int, n int) []int{

          fmt.Println(m)

    x:=len(nums1)+len(nums2)
    a:=make([]int,x)
    i:=0
    j:=0
    k:=0
    for i<m && j<n {
        if(nums1[i]<nums2[j]){
            a[k]=nums1[i]
             k++
            i++
            
        } else {
            
            a[k]=nums2[j]
            k++
            j++
        }
    }
    for i<m {
        a[k]=nums1[i]
        k++
        i++   
    }
    for j<n{
      a[k]=nums2[j]
         k++
        j++
    }
    
    
    return a
    
    
    
}