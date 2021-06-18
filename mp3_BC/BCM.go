package main

// block으로 엮을 순서 mp3 데이터 / 계약서 동의 / 개인키 / 생성시간

import(
	"project.com/BC"
	"project.com/mp3_bytes"
	"fmt"
	"time"
	"strconv"
	"reflect"
)


func main(){
	NC := BC.InitBlockChain()
	
	Mbytes := mp3_bytes.MtoB("/workspace/GO_WORK/mp3_BC/SAMPLE_1.MP3")
	NC.AddBlock(Mbytes)
	
	AgreeContract := "Do you accept the contract?"
	ConBytes := []byte(AgreeContract)
	NC.AddBlock(ConBytes)
	
	PrivateKey := "This is PrivateKey"
	PkBytes := []byte(PrivateKey)
	NC.AddBlock(PkBytes)
	
	timestamp := time.Now().Unix()
	fmt.Println(reflect.TypeOf(timestamp))
	timestring := strconv.FormatInt(timestamp,10)
	TimeBytes := []byte(timestring)
	NC.AddBlock(TimeBytes)
	fmt.Println(NC.Blocks)
	
}

