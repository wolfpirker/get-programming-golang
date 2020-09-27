
package main

import (
    "fmt"
    )
    
func main(){
    const distance = 56000000
    const h_per_day = 24.0
    const days = 28.0
    var required_speed = distance/(days*h_per_day)
    fmt.Printf("required speed %28v\n", required_speed)    
    fmt.Println(distance/(days*h_per_day), "km/h")
}

