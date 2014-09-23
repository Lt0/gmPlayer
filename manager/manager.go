package manager

import "os"
import "fmt"
import "path/filepath"
import "EagleEye/Win32Api"
import "io"
import "encoding/json"

func GetMusicPath() string {
	npath := Win32Api.GetStartupDirectory()
	npath = filepath.Join(npath, "pathfile.txt")
	
	f, err1 := os.OpenFile(npath, os.O_CREATE, 0)
	
	if err1 != nil {
		fmt.Println("openfile failed, error:", err1)
	}
	
	tmp2 := make([]byte, 1024)
	var musicpath string
	num, err2 := f.Read(tmp2)
	tmp := make([]byte, num)
	for k, _ := range(tmp) {
		tmp[k] = tmp2[k]
	}
	

	if err2 == io.EOF || err2 != nil || filepath.IsAbs(string(tmp)) != true {
		if filepath.IsAbs(string(tmp)) != true && err2 != io.EOF{
			f, _ = os.Create(npath)

		}
		fmt.Println("the file is nil or the path is incorrect. error:", err2)
		fmt.Println("input the musicpath:")
		var path string
		fmt.Scanln(&path)
		_, err3 := f.WriteString(path)
		if err3 != nil {
			fmt.Println("write failed, error:", err3)
		} else {
			musicpath = path
			fmt.Println("Write succeed.")
		}
	} else if err2 == nil {
		musicpath = string(tmp)
	} 
	return musicpath
}

type Setting struct {
	musicpath string
	playmode string
}



func GetPlayMode() string {
	path := Win32Api.GetStartupDirectory()
	path = filepath.Join(path, "playmode.txt")
	var pmode string
	var mSetting Setting
	f, err := os.OpenFile(path, os.O_CREATE, 0)
	if err != nil {
		fmt.Println("Open playmode.txt failed.Error:", err)
	}
	setinfo := make([]byte, 1024)
	n1, err1 := f.Read(setinfo)
	if n1 == 0 {
		fmt.Println("The settingfile.txt is nil.\nPlease input the playmode(order or random):")
		var nplaymode string
		fmt.Scan(&nplaymode)
		su := false
		for {
			switch nplaymode {
				case "order":
					mSetting.playmode = "order"
					b, err3 := json.Marshal(mSetting)
					if err3 != nil {
						fmt.Println("Marshal mSetting failed.")
					} else {
						fmt.Println("Marshal mSetting succeed.")
					}
					_, err4 := f.Write(b)
					if err4 != nil {
						fmt.Println("Write b to settingfile.txt failed. error:", err)
					} else {
						fmt.Println("Write b to settingfile.txt succeed")
					}
					pmode = nplaymode
					su = true
				case "random":
					mSetting.playmode = "random"
					fmt.Println("The mSetting.playmode is:", mSetting.playmode)
					b, err3 := json.Marshal(mSetting)
					if err3 != nil {
						fmt.Println("Marshal mSetting failed.")
					} else {
						fmt.Println("Marshal mSetting succeed.")
					}
					_, err4 := f.Write(b)
					if err4 != nil {
						fmt.Println("Write b to settingfile.txt failed. error:", err)
					} else {
						fmt.Println("Write b to settingfile.txt succeed")
					}
					fmt.Println("The mSetting.playmode is:", mSetting.playmode)
					fmt.Println("The b is :", string(b))
					pmode = nplaymode
					su = true
				default :
					fmt.Println("Unexpected command.")
					fmt.Scan(&nplaymode)
			}
			if su == true {
				break
			}
		}
	} else if err1 != nil {
		fmt.Println("Read settingfile.txt failed. Error:", err1)
	} else {
		err2 := json.Unmarshal(setinfo, mSetting)
		if err2 != nil {
			fmt.Println("Unmarshal the settingfile.txt failed.", err2)
		} else {
			fmt.Println("Unmarshal the settingfile.txt succeed.")
		}
		switch mSetting.playmode {
			case "order":
				fmt.Println("The playmode is order.")
				pmode = mSetting.playmode
			case "random":
				fmt.Println("The playmode is random.")
				pmode = mSetting.playmode
			default :
				fmt.Println("The playmode is unexpected.")
				var nplaymode string
				fmt.Scan(&nplaymode)
				for {
					switch nplaymode {
						case "order":
							mSetting.playmode = nplaymode
							b, err3 := json.Marshal(mSetting)
							if err3 != nil {
								fmt.Println("Marshal mSetting failed.")
							} else {
								fmt.Println("Marshal mSetting succeed.")
							}
							_, err4 := f.Write(b)
							if err4 != nil {
								fmt.Println("Write b to settingfile.txt failed. error:", err)
							} else {
								fmt.Println("Write b to settingfile.txt succeed")
							}
							pmode = nplaymode
							break
						case "random":
							mSetting.playmode = nplaymode
							b, err3 := json.Marshal(mSetting)
							if err3 != nil {
								fmt.Println("Marshal mSetting failed.")
							} else {
								fmt.Println("Marshal mSetting succeed.")
							}
							_, err4 := f.Write(b)
							if err4 != nil {
								fmt.Println("Write b to settingfile.txt failed. error:", err)
							} else {
								fmt.Println("Write b to settingfile.txt succeed")
							}
							pmode = nplaymode
							break
						default :
							fmt.Println("Unexpected command.")
							fmt.Scan(&nplaymode)
					}
				}//for
		}
	}//else
	return pmode
}


	
			