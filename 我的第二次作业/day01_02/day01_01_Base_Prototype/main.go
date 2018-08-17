package main

import (
	"day01_02/day01_01_Base_Prototype/BLC"
	"fmt"
)

func main() {

	//bytes:=make([]byte,3,3)
	//fmt.Println(bytes)
	//1.测试区块
/*	blockQB:=BLC.NewBlockQB("I am  a blockQB",make([]byte,32,32),0)
	fmt.Println(blockQB)
	fmt.Println(blockQB.TimeStampQB)
	fmt.Println(blockQB.PrevBlockHashQB)
	fmt.Println(blockQB.HashQB)
*/
	//2.测试创世区块
	genesisBlockQB:=BLC.CreateGenesisBlockQB("Genesis BlockQB")
	fmt.Println(genesisBlockQB)

	//3.创建一个区块链
	blockChainQB := BLC.CreateBlockChainWithGenesisBlockQB("Genesis Block")
	/*	fmt.Println(blockChain)
		fmt.Println(blockChain.BlocksQB) // blockchain 的字段数组
		fmt.Println(blockChain.BlocksQB[0]) //创世区块
	*/
	//4.测试添加区块
	fmt.Println(len(blockChainQB.BlocksQB)) //区块链中，有一个区块：创世区块
	blockChainQB.AddBlockToBlockChainQB("send 100RMB to qubiao", blockChainQB.BlocksQB[len(blockChainQB.BlocksQB)-1].HashQB, blockChainQB.BlocksQB[len(blockChainQB.BlocksQB)-1].HeightQB+1)
	blockChainQB.AddBlockToBlockChainQB("send 50RMB to laoda", blockChainQB.BlocksQB[len(blockChainQB.BlocksQB)-1].HashQB, blockChainQB.BlocksQB[len(blockChainQB.BlocksQB)-1].HeightQB+1)
	blockChainQB.AddBlockToBlockChainQB("send 3RMB to laoer", blockChainQB.BlocksQB[len(blockChainQB.BlocksQB)-1].HashQB, blockChainQB.BlocksQB[len(blockChainQB.BlocksQB)-1].HeightQB+1)
	blockChainQB.AddBlockToBlockChainQB("send 1000RMB to mm", blockChainQB.BlocksQB[len(blockChainQB.BlocksQB)-1].HashQB, blockChainQB.BlocksQB[len(blockChainQB.BlocksQB)-1].HeightQB+1)
	fmt.Println(blockChainQB)
	fmt.Println(blockChainQB.BlocksQB[1].HeightQB)
	fmt.Println(blockChainQB.BlocksQB[1].PrevBlockHashQB)
	fmt.Println(blockChainQB.BlocksQB[0].HashQB)

	fmt.Println(len(blockChainQB.BlocksQB))

}
