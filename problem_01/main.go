package main

import "fmt"

func main(){
	var count int
	for i := 0; i < 1000; i++{
		if( i % 3 == 0){
			count += i
		}else if(i % 5 == 0){
			count += i
		}
	}
	fmt.Printf("The total count is %v", count)
}