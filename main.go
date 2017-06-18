package main

import (
	"flag"
	"fmt"
	"strconv"
	"math/rand"
	"math"
	"os"
	"time"
)

func cacu(a int)float64{
	var sum float64 = 0
	for i:= 0;i<a;i++{
		x1 := rand.Float64()
		x2 := rand.Float64()
		y1 := rand.Float64()
		y2 := rand.Float64()
		dx:=x2-x1
		dx = dx*dx
		dy:=y2-y1
		dy = dy*dy
		distence := math.Sqrt(dy+dx)
		sum = sum + distence
		}
	times := float64(a)
	avg := sum/times
	return avg
}
func main(){
	startTime := time.Now()
	flag.Parse()
	if flag.NArg() != 1{
		fmt.Print("one and only one parameter allowed")
		os.Exit(1)
	}
	arg := flag.Arg(0)
	times,err :=strconv.Atoi(arg)
	if err != nil{
		fmt.Println("parameter must be a int number\n , exit with code 2")
		os.Exit(2)
	} else if times > 50000000{
	fmt.Println("数字太大，重新输入")
	os.Exit(3)
	}

	//得到次数 fmt.Println(times)
	avg := cacu(times)
	fmt.Print(avg)
	fmt.Println("\n")
	dur:= time.Now().Sub(startTime).Nanoseconds()
	co := dur/1000000
	d := strconv.FormatInt(co,10)

	fmt.Printf("共耗时"+d+"毫秒")

}
