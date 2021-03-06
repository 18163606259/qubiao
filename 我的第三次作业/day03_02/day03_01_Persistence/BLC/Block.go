package BLC

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	//字段属性
	//1.高度：区块在区块链中的编号，第一个区块页叫创世区块，为0
	HeightQB int64
	//2.上一个区块的Hash值
	PrevBlockHashQB []byte
	//3.数据：data，交易数据
	DataQB []byte
	//4.时间戳
	TimeStampQB int64
	//5.自己的hash
	HashQB []byte

	//6.Nonce
	NonceQB int64
}

func NewBlock(data string, prevBlockHash [] byte, height int64) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil,0}
	//设置hash
	//block.SetHash()
	pow:=NewProofOfWorkQB(block)
	hash,nonce:=pow.Run()
	block.HashQB = hash
	block.NonceQB = nonce

	return block
}

/*-------------------------------------------------------------------------------------------------------------------------*/

func CreateGenesisBlockQB(data string) *Block{

	return NewBlock(data,make([]byte,32,32),0)
}

//定义block的方法，用于序列化该block对象，获取[]byte
func (block *Block) Serialize()[]byte{
	//1.创建一个buff
	var buf bytes.Buffer

	//2.创建一个编码器
	encoder:=gob.NewEncoder(&buf)

	//3.编码
	err:=encoder.Encode(block)
	if err != nil{
		log.Panic(err)
	}

	return buf.Bytes()
}

//定义一个函数，用于将[]byte，转为block对象，反序列化
func DeserializeBlock(blockBytes [] byte) *Block{
	var block Block
	//1.先创建一个reader
	reader:=bytes.NewReader(blockBytes)
	//2.创建解码器
	decoder:=gob.NewDecoder(reader)
	//3.解码
	err:=decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}
