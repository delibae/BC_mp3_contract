package main
 
import ("io/ioutil"
		"os"
	   )
		
 
func main() {
    //파일 읽기
    bytes, err := ioutil.ReadFile("/workspace/GO_WORK/SAMPLE_1.MP3")
    if err != nil {
        panic(err)
    }
	
	er := ioutil.WriteFile("/workspace/GO_WORK/New.mp3",bytes,os.FileMode(644))
	if er != nil {
        panic(er)
    }
	
}

	