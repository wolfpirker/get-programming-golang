
package main

import (
    "fmt"
    )

 
    
func main(){
    var quote = "L fdph, L vdz, L frqtxhuhg."
    for _, c := range quote {
        switch c {
        case 'a', 'b', 'c':
            fmt.Printf("%c, ", 'x' + c - 'a')    
        default:
            fmt.Printf("%c, ", c - 3)    
        }
        
    }
}
