package BLC

//定义一个区块链
type BlockChainQB struct {
	BlocksQB []*BlockQB
}

//创建一个区块链，包含创世区块
func CreateBlockChainWithGenesisBlockQB(data string) *BlockChainQB {
	//1.创建创世区块
	genesisBlockQB := CreateGenesisBlockQB(data)
	//2.创建区块链对象并返回
	return &BlockChainQB{[]*BlockQB{genesisBlockQB}}
}

//添加区块到区块链中
func (bc *BlockChainQB) AddBlockToBlockChainQB(data string, prevBlockHash [] byte, height int64) {
	//1.根据参数的数据，创建Block
	newBlock := NewBlockQB(data, prevBlockHash, height)
	//2.将block加入blockchain
	bc.BlocksQB = append(bc.BlocksQB, newBlock)
}
