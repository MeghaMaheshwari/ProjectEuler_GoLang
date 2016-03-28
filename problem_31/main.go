package main

import "fmt"

var TargetVal int = 200
var coins = [8]int {1, 2, 5, 10, 20, 50, 100, 200 }
var TotalWays int = 0


func main() {
     for i := 0; i < len(coins); i++{
         if (TargetVal % coins[i] == 0){
              TotalWays++
         }
         for j := (TargetVal/coins[i]-1); j >= 1; j--{
            RemainingVal := TargetVal - (j*coins[i])
            if(RemainingVal < coins[i+1]){
                continue;
            }
            GetWays(i+1, RemainingVal)
         }
     }
     fmt.Println(TotalWays)
}

func GetWays(CurrentIndex int, ValToForm int){
       for j := CurrentIndex; j < len(coins); j++{
            if (ValToForm < coins[j]){
                continue
            }
            for k := (ValToForm / coins[j]); k >= 1; k--{
                 RemainVal := ValToForm - (k * coins[j])
                 if (RemainVal == 0){
                    TotalWays++
                 }else{
			 // if remaining value is less than the next denomination to be formed
			 // simply continue as this value can no way be formed with the remaining denomination
                     if (RemainVal < coins[j + 1]){
                        continue
                      }
                      GetWays(j + 1, RemainVal)
                    }
                 }
       }
  }


