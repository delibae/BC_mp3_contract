package main

// block으로 엮을 순서 mp3 데이터 / 계약서 동의 / 개인키 / 생성시간

import(
	"project.com/BC"
	"project.com/mp3_bytes"
	"fmt"
	"time"
	"strconv"
	"reflect"
	"encoding/json"
	"io/ioutil"
	"os"
)


func main(){
	NC := BC.InitBlockChain()
	
	Mbytes := mp3_bytes.MtoB("/workspace/GO_WORK/mp3_BC/SAMPLE_1.MP3")
	NC.AddBlock(Mbytes)
	
	AgreeContract := "Do you accept the contract?"
	fmt.Println(AgreeContract)
	var agree string 
	fmt.Scanln(&agree)
	ConBytes := []byte(AgreeContract + agree)
	NC.AddBlock(ConBytes)
	
	message := "PrivateKey:"
	fmt.Println(message)
	var PrivateKey string
	fmt.Scanln(&PrivateKey)
	PkBytes := []byte(message + PrivateKey)
	NC.AddBlock(PkBytes)
	
	timestamp := time.Now().Unix()
	fmt.Println(reflect.TypeOf(timestamp))
	timestring := strconv.FormatInt(timestamp,10)
	TimeBytes := []byte(timestring)
	NC.AddBlock(TimeBytes)
	
	Save := NC.Blocks
	Save1 := make([]BC.Block,len(Save))
	for i := 1; i <= len(Save); i++{
		Save1[i-1] = *Save[i-1]
	}
	// fmt.Println(Save1)
	byteArray, err := json.Marshal(Save1)
    if err != nil {
        fmt.Println(err)
    }
	er := ioutil.WriteFile("./BCM.json", byteArray, os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if er != nil {
		fmt.Println(er)
		return 
	}
}

