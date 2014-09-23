package main

import "mTinyPlayer/manager"
import "mTinyPlayer/manager/getlist"
import "mTinyPlayer/manager/playmode"
import "fmt"
import "runtime"
// import ap "EagleEye/AudioPlayer"
// import "time"

func main() {
	mpath := manager.GetMusicPath()
	filelist := getlist.GetFileList(mpath)
	musiclist := getlist.GetMusicList(filelist)
	
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
	}
	runtime.GOMAXPROCS(10)
	go playmode.Random(musiclist, chs)
	
	var s string
	for {
		// fmt.Println("Supported command: \nplay/pause/next/pre/stop/exit")
		fmt.Scan(&s)
		switch s {
			case "exit":
				go wc(chs[5])
				runtime.Goexit()
			case "play":
				go wc(chs[0])
			case "pause":
				go wc(chs[1])
			case "next":
				go wc(chs[2])
			case "pre":
				go wc(chs[3])
			case "stop":
				go wc(chs[4])
			// case "order":
				// go wc(chs[5])
				// go playmode.Order(musiclist, chs)
			// case "range":
				// go wc(chs[5])
				// go playmode.Random(musiclist, chs)
			default :
				fmt.Println("Unsupported command!")
		}
	}
}

func wc(ch chan int) {
	fmt.Println("Sending command....")
	ch <- 98798798
	fmt.Println("The command was sended.")
	runtime.Goexit()
}