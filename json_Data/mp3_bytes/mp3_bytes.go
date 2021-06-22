package mp3_bytes
 
import ("io/ioutil"
		"os"
	   )
		
 
func MtoB(address string) []byte{
    //파일 읽기
    bytes, err := ioutil.ReadFile(address)
    if err != nil {
        panic(err)
    }
	return bytes
}

func BtoM(bytes []byte, address string) {
	er := ioutil.WriteFile(address,bytes,os.FileMode(644))
	if er != nil {
        panic(er)
    }
}

