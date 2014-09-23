package playmode

import ap "EagleEye/AudioPlayer"
import "fmt"
import "time"
import "math/rand"
import "runtime"

func Random(musiclist []string, chs []chan int) {
	fmt.Println("Current paly mode is Random.")
	player := ap.New()
	player.GetDevice()
	historylist := make([]string, 0)
	premusicindex := -1
	var nextm int
	var pos string
	var timeGap int
	var isTheSamePos string
	playNextm := false
	pors := false//获取是否处于暂停或停止状态，若是则不再获取当前播放位置，因为GetPos接口有问题，
				 //每一次调用都有可能返回已完成的string，避免在暂停时多次调用该接口导致出错而在暂停或停止时播放下一首
	
	for {
		nextm = GetNextM("random", 1, musiclist)
		premusicindex++
		historylist = append(historylist, musiclist[nextm])
		player.Stop()
		player.LoadAudio(musiclist[nextm])
		isTheSamePos = "Now is the new music"
		player.Play()
		pors = false
		fmt.Println("It is playing", musiclist[nextm])	
		playNextm = false
		for {
			select {
				case <- chs[0]:
					player.Play()
					pors = false
					fmt.Println("It is playing", musiclist[nextm])
				case <- chs[1]:
					player.Pause()
					pors = true
				case <- chs[2]:
					// player.Stop()
					// nextm = GetNextM("random", 1, musiclist)
					// premusicindex++
					// historylist = append(historylist, musiclist[nextm])
					// player.LoadAudio(musiclist[nextm])
					// isTheSamePos = "Now is the new music"
					// player.Play()
					// fmt.Println("It is playing", musiclist[nextm])		
					
					playNextm = true
				case <- chs[3]:
					if premusicindex != 0 {
						player.Stop()
						premusicindex = premusicindex -1
						player.LoadAudio(historylist[premusicindex])
						isTheSamePos = "Now is the new music"
						player.Play()
						pors = false
						fmt.Println("It is playing", historylist[premusicindex])
					} else {
						fmt.Println("It is the first music of the play history.")
					}
				case <- chs[4]:
					player.Stop()
					pors = true
					// runtime.Goexit()
				case <- chs[5]://mode order----exit range play mode here
					runtime.Goexit()
				default:
					// fmt.Println("in default.......................")
					if timeGap == 30 && pors == false{
						// fmt.Println("in default.......................")
						timeGap = 0
						pos = player.GetPos()
						// fmt.Println("Get the pos finished.")
						if isTheSamePos == pos{
							// fmt.Println("in isTheSamePos==pos,next step should be change the playNextm to be true")
							playNextm = true
							// fmt.Println("playNextm is true in default")
						}
						isTheSamePos = pos
						// fmt.Println(pos)
						// fmt.Println(len(pos))
						// fmt.Println(pos[0], "\n--------------------------------------------------------------")
						if pos[0] == '0' {
							// fmt.Println("The pos[0] is 0, the next step should be change the playNextm to be true")
							// pos = "9999999999999999999999999999"
							playNextm = true
							// fmt.Println("playNextm is true in default")
						}
					}
			}
			if playNextm == true {
				// fmt.Println("playNextm is true")
				break
			}
			time.Sleep(100 *time.Millisecond)	
			timeGap++
			
			// if timeGap == 10 {
				// fmt.Println("timeGap----", timeGap)
			// }
			if timeGap == 50 {
				// fmt.Println("clear timeGap out of default.")
				timeGap = 0
				// if isTheSamePos == pos {
							// break
				// }
			}
			
		}
	}
}

func Order(musiclist []string, chs []chan int) {
	fmt.Println("Current paly mode is order")
	player := ap.New()
	nextm := -1
	var pos string
	
	for {
		nextm = nextm + 1
		player.Stop()
		player.LoadAudio(musiclist[nextm])
		fmt.Println("It is playing", musiclist[nextm])
		player.Play()
		for {
			time.Sleep(100 *time.Millisecond)
			select {
				case <- chs[0]://play
					fmt.Println("It is playing", musiclist[nextm])
					player.Play()
				case <- chs[1]://pause
					// fmt.Println("It is playing", musiclist[nextm])
					player.Pause()
				case <- chs[2]://next
					player.Stop()
					pos = "99999999999999999999999999999"
					nextm = nextm + 1
					player.LoadAudio(musiclist[nextm])
					fmt.Println("It is playing", musiclist[nextm])
					player.Play()
				case <-chs[3]:   //pre
 					if nextm != 0 {
						player.Stop()
						pos = "99999999999999999999999999999"
						nextm = nextm - 1
						player.LoadAudio(musiclist[nextm])
						fmt.Println("It is playing", musiclist[nextm])
						player.Play()
					} else {
						fmt.Println("It is the first music of the music list.")
					}
				case <- chs[4]://stop
					player.Stop()
				case <- chs[5]://mode random-------exit order play mode here
					runtime.Goexit()
			}
			pos = player.GetPos()
			if pos[0] == '0' {
				// fmt.Println("The pos[0] is 0, break the for loop.")
				// pos = "9999999999999999999999999999"
				break
			}
			
			}
	}
}



func GetNextM(mode string, pre int, musiclist []string) int {
	var nextm int
	switch mode {
		case "order":
			if pre < len(musiclist)-1 {
				nextm = pre + 1
			} else {
				nextm = 0
			}
		case "random":
			//定义随机数生成器nexM作为下一首音乐的index
			// var n int32
			n := int32(len(musiclist))
			next := rand.NewSource(time.Now().Unix())
			nextR := rand.New(next)
			nextm = int(nextR.Int31n(n))
		default:
	}
	return nextm
}

