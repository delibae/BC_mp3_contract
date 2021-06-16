package main
 
import ("io/ioutil"
		"fmt"
	   )
		
 
func main() {
    //파일 읽기
    bytes, err := ioutil.ReadFile("/workspace/GO_WORK/SAMPLE_1.MP3")
    if err != nil {
        panic(err)
    }
	fmt.Println(bytes)
}

	