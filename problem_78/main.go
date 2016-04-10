package main

import "fmt"

func main() {
        const TargetVal int = 100000
        var Smallestval int = 0
        var ways[TargetVal+1] int
        ways[0] = 1
        for i := 1; i < TargetVal; i++{
           for j := i; j <= TargetVal; j++{
              ways[j] = (ways[j] + ways[j - i]) % 1000000;
           }
           if(ways[i] == 0){
                   Smallestval = i
                   break;
           }
         }
        fmt.Println(Smallestval)
    }




