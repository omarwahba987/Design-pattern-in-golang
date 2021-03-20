package main

import "fmt"

func main() {

	channelList := make([]chan int, 0)
	//for i := 0; i < 2; i++ {
	//	channelList = append(channelList, make(chan int))
	//}

	FChannel := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		FChannel <- i // contain values from 1 to 5
	}
	//close(FChannel)


	SChannel := make(chan int, 5)
	chanel := make (chan int,5)
	chanel<-4
	//var SChannel chan int
	for j := 6; j <= 10; j++ {
		SChannel <- j // contain values from 6 to 10
	}
	//close(SChannel)
	channelList = append(channelList, FChannel)
	channelList = append(channelList, SChannel)

	//for _, x := range channelList {
	//	for f := range x {
	//		fmt.Println(f)
	//
	//	}
	//}

	//time.Sleep(5*time.Second)
	//fmt.Println(channelList)

	//for _, c := range channelList {
	//c:=channelList[1]
	//type x chan int
	/*	cl:= make([]chan int,0)
		FChannel := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			FChannel <- i // contain values from 1 to 5
		}
		SChannel := make(chan int, 5)
		for j := 6; j <= 10; j++ {
			SChannel <- j // contain values from 6 to 10
		}
		cl=append(cl,FChannel)
		cl=append(cl,SChannel)
	*/

	//fmt.Println(<-c)
	for _, c := range channelList {
	fo:
		for {
			select {
			case i := <-c:
				fmt.Println(i)
			default:
				break fo
			}

		}
	}
	//for _, c := range channelList {

	//c:=make(chan int,4)
	//c<-1
	//c<-3
	//c<-5
	//c<-7
	//close(c)
	//	for v:= range c {
	//		fmt.Println(v)
	//	}
	//}
	//}

}
