package main

import "fmt"

const NoOfDaysInNonLeapYear int = 365
const NoOfDaysInLeapYear int = 366
const NoOfDaysInFourWeeks int = 28
const NoOfDaysInFebLeapYear int = 29
var Days = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday","Saurday","Sunday"}
var DaysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var StartDayVsSundaysNonleap = make(map[string]int)
var StartDayVsSundayleap = make(map[string]int)
//Jan 1st 1901 was a tuesday
var StartDayOfTheYear int = 1
var IsLeapYear bool = false

func main() {
	var TotalNumberofSundays int = 0
	for index, _ := range Days {
		StartDayVsSundaysNonleap[Days[index]] = getnumberofsundays(index, false)
		StartDayVsSundayleap[Days[index]] = getnumberofsundays(index, true)
	}

	for i := 1901; i <= 2000; i++{
	 IsLeapYear := false;
	if ((i % 4 == 0) || ((i % 4 == 0) && (i % 400 == 0))){
	  IsLeapYear = true;
	}
	var SundaysInThisYear int
	if(IsLeapYear){
	  SundaysInThisYear = StartDayVsSundayleap[Days[StartDayOfTheYear]]
	}else{
	   SundaysInThisYear = StartDayVsSundaysNonleap[Days[StartDayOfTheYear]]
	}
	TotalNumberofSundays = TotalNumberofSundays + SundaysInThisYear;
	var NumberOfDaysInTheYear int
        if(IsLeapYear){
	    NumberOfDaysInTheYear = NoOfDaysInLeapYear
	}else{
		NumberOfDaysInTheYear = NoOfDaysInNonLeapYear
	}
	StartDayOfTheYear = (StartDayOfTheYear + NumberOfDaysInTheYear) % 7;
	}
	fmt.Printf("Number Of Sundays %d", TotalNumberofSundays )
}

	func getnumberofsundays(StartDay int,IsleapYear bool) int {
		var TotalSundays int = 0
		for i,_:= range DaysInMonth {
			if (i != 0) {
			  if (Days[StartDay] == "Sunday"){
		                TotalSundays = TotalSundays+1;
		           }
				// each month will have 4 complete weeks and some remaining days depending upon the number of days. Based on the Start day this remaining
				//days will determine if the start day of the next month was a Sunday.
				var NumberOfDaysinMonth int
				if (i == 2 && IsleapYear) {
					NumberOfDaysinMonth = NoOfDaysInFebLeapYear
				}else {
					NumberOfDaysinMonth = DaysInMonth[i]
				}
				var NumberOfNonWeekDays = NumberOfDaysinMonth - NoOfDaysInFourWeeks;
				StartDay = (StartDay + NumberOfNonWeekDays) % 7;

			}
		}
		return TotalSundays;
	}
