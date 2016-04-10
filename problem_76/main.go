package main

import "fmt"

var TargetVal int = 100
var MaxVal int = 99
var TotalWays int = 0


func main() {
     for i := 1; i <= MaxVal; i++{
         if (TargetVal % i == 0){
              TotalWays++
         }
         for j := (TargetVal/i-1); j >= 1; j--{
            RemainingVal := TargetVal - (j*i)
            if(RemainingVal < (i+1)){
                continue;
            }
            GetWays(i+1, RemainingVal)
         }
     }
     fmt.Println(TotalWays)
}

func GetWays(CurrentIndex int, ValToForm int){
       for j := CurrentIndex; j <= MaxVal; j++{
            if (ValToForm < j){
                continue
            }
            for k := (ValToForm / j); k >= 1; k--{
                 RemainVal := ValToForm - (k * j)
                 if (RemainVal == 0){
                    TotalWays++
                 }else{
			 // if remaining value is less than the next denomination to be formed
			 // simply continue as this value can no way be formed with the remaining denomination
                     if (RemainVal < (j+1)){
                        continue
                      }
                      GetWays(j + 1, RemainVal)
                    }
                 }
       }
  }


