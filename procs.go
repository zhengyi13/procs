package main

import (
	"fmt"
	"github.com/zhengyi13/procs/procdata"
)

func main() {
	mypid, err := procdata.GetMyPid()
	if err != nil { fmt.Printf("Fucked up: %v", err)}
	fmt.Printf("My PID data: %+v\n", mypid)
	myppid, err := procdata.GetParentPid(mypid)
	if err != nil { fmt.Printf("Fucked up: %v", err)}
	fmt.Printf("My parent is %d\n", myppid)
}
