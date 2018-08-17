package BLC

import (
	"time"

)

//step1:定一个一Block
type BlockQB struct {
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

//step2：提供一个函数用于创建一个区块
func NewBlockQB(data string, prevBlockHash [] byte, height int64) *BlockQB {
	//创建区块
	block := &BlockQB{height, prevBlockHash, []byte(data), time.Now().Unix(), nil,0}
	//设置hash
	//block.SetHash()
	pow:=NewProofOfWorkQB(block)
	hash,nonce:=pow.Run()
	block.HashQB = hash
	block.NonceQB = nonce

	return block
}



//step4:创建一个创世区块
func CreateGenesisBlockQB(data string) *BlockQB{

	return NewBlockQB(data,make([]byte,32,32),0)
}
