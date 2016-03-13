package main

import "fmt"

func main() {

OneDigits := [20]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven","twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

TenDigits := [8]string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

const OneThousand string = "OneThousand"

const HundredAnd string = "HundredAnd"

var SumUntilNine int = 0

var SumFromTenToNineteen int = 0

var SumOfTens int = 0

var HundredSum int = 0

// take the sum from one to nine and store it
SumUntilNine = GetCount(OneDigits[1:10])
// take the sum from 10 to 19 and store it
SumFromTenToNineteen = GetCount(OneDigits[10:20])
// take the sum of all tens
SumOfTens = GetCount(TenDigits[0:8])
// Digits one from nine occur in all 21 to 99 and so we can take the sum from
// one to nine and multiply it by 8 (21-29, 31-39, 41-49, 51-59,61-69,71-79,81-89,91-99 = 8 times)
// each of twenty, thirty, forty occur ten times in 20-29, 30-39 and hence we take the sum of tens
// and multiply it with 10
var SumFromTwentyToNinetyNine int = SumOfTens * 10 + SumUntilNine * 8;
// sum from one to 99 is = sum from one to 9,ten to nineteen and twenty to 99
var total int = SumUntilNine + SumFromTenToNineteen + SumFromTwentyToNinetyNine;

 for i := 1; i <= 9; i++ {
   // the word onehundredand, twohundredand, threehundred and appear in 100-199 , hundred times
   // that is why, we take each of onehundredand , twohundredand and multiply it with 100. Since the
   // first onehundred, two hundred do not contain and, we have an extra "and" added and that is why we subtract
  // the length of and from the total of 100-199, 200-299. The rest of the values 1-99 have the same sum as
  // previously calculated and hence we add that.
   val := len(OneDigits[i]) + len(HundredAnd) ;
   HundredSum = (HundredSum + val*100 + total)-3;            
  }
 // add the last one thousand to the total sum
 total = total + HundredSum + len(OneThousand);
 fmt.Println(total)  
}

func GetCount(value[] string) int {
    sum := 0
    for _,val := range value {
      sum += len(val)
   }
   return sum
}






