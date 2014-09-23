package getlist

import "fmt"
import "os"
import "path/filepath"
import "path"
import "mTinyPlayer/manager"

func GetFileList(path string) (list []string) {
	fmt.Println("Running GetFileList....")
	mlist := make([]string, 0)
	// fmt.Println(len(mlist))
	// su := false
	for {
		err := filepath.Walk(path, func(tmppath string, f os.FileInfo, err error) error {
			if f == nil {return err}
			if f.IsDir() {return nil}
			mlist = append(mlist, tmppath)
			// fmt.Println("Get the file:", tmppath)
			// fmt.Println( "Append ：", mlist[(len(mlist)-1)], "to mlist succeed.")
			return nil
		})
		
		if err != nil {
			fmt.Println("GetFileList failed.\n Error:", err)
			fmt.Println("--------------------------------------------------------\n Input the corrent music path:")

			path = manager.GetMusicPath()
			
		} else {
			break
		}
	}
	
	return mlist
}


func GetMusicList(filelist []string) (musiclist []string) {
	fmt.Println("Running GetMusicList....")
	mlist := make([]string, 0)
	for _, v := range(filelist) {
		if IsMusic(v) {
			mlist = append(mlist, v)
		}
	}
	for k, v := range(mlist) {
		fmt.Println(k, "-", v)
	}
	fmt.Println("There are", len(mlist), "music in the directory.")
	return mlist
}


func IsMusic(name string) bool {
	filetype := path.Ext(name)
	b := false
	switch filetype {
		case ".mp3":
			b = true
		case ".wav":
			b = true
		default : 
	}
	return b	
}

