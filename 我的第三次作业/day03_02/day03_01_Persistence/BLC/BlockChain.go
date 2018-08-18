package BLC

import (
	"github.com/boltdb/bolt"
	"os"
	"fmt"
	"log"
	"math/big"
	"time"
)

//定义一个区块链
type BlockChainQB struct {
	//Blocks []*Block
	DB  *bolt.DB //对应的数据库对象
	Tip [] byte  //存储区块中最后一个块的hash值
}

//创建一个区块链，包含创世区块
func CreateBlockChainWithGenesisBlockQB(data string) *BlockChainQB {

	/*
	1.判断数据库如果存在，
	2.数据库不存在，创建创世区块，并存入到数据库中
	 */
	if dbExists() {
		fmt.Println("数据库已经存在。。。")
		//打开数据库
		db, err := bolt.Open(DBNameQB, 0600, nil)
		if err != nil {
			log.Panic(err)
		}

		var blockchain *BlockChainQB

		err = db.View(func(tx *bolt.Tx) error {
			//打开bucket，读取l对应的最新的hash
			b := tx.Bucket([]byte(BlockBucketNameQB))
			if b != nil {
				//读取最新hash
				hash := b.Get([]byte("l"))
				blockchain = &BlockChainQB{db, hash}
			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		return blockchain
	}

	//数据库不存在
	fmt.Println("数据库不存在。。")
	/*
	1.创建创世区块
	2.存入到数据库中
	 */
	genesisBlock := CreateGenesisBlockQB(data)
	db, err := bolt.Open(DBNameQB, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		//创世区块序列化后，存入到数据库中
		b, err := tx.CreateBucketIfNotExists([]byte(BlockBucketNameQB))
		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			err = b.Put(genesisBlock.HashQB, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			b.Put([]byte("l"), genesisBlock.HashQB)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return &BlockChainQB{db, genesisBlock.HashQB}
}

//添加区块到区块链中
func (QB *BlockChainQB) AddBlockToBlockChainQB(data string) {
	//1.根据参数的数据，创建Block
	//newBlock := NewBlock(data, prevBlockHash, height)
	//2.将block加入blockchain
	//bc.Blocks = append(bc.Blocks, newBlock)
	/*
	1.操作bc对象，获取DB
	2.创建新的区块
	3.序列化后存入到数据库中
	 */
	err := QB.DB.Update(func(tx *bolt.Tx) error {
		//打开bucket
		b := tx.Bucket([]byte(BlockBucketNameQB))
		if b != nil {
			//获取bc的Tip就是最新hash，从数据库中读取最后一个block：hash，height
			blockByets := b.Get(QB.Tip)
			lastBlock := DeserializeBlock(blockByets) //数据库中的最后一个区块
			//创建新的区块
			newBlock := NewBlock(data, lastBlock.HashQB, lastBlock.HeightQB+1)
			//序列化后存入到数据库中
			err := b.Put(newBlock.HashQB, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			//更新：bc的tip，以及数据库中l的值
			b.Put([]byte("l"), newBlock.HashQB)
			QB.Tip = newBlock.HashQB

		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

//提供一个方法，用于判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(DBNameQB); os.IsNotExist(err) {
		return false
	}
	return true
}

//新增方法，用于遍历数据库，打印所有的区块
func (QB *BlockChainQB) PrintChains() {
	/*
	.bc.DB.View(),
		根据hash，获取block的数据
		反序列化
		打印输出


	 */

	//获取迭代器
	it := QB.Iterator()
	for {
		//step1：根据currenthash获取对应的区块
		block := it.Next()
		fmt.Printf("第%d个区块的信息：\n", block.HeightQB+1)
		fmt.Printf("\t高度：%d\n", block.HeightQB)
		fmt.Printf("\t上一个区块Hash：%x\n", block.PrevBlockHashQB)
		fmt.Printf("\t自己的Hash：%x\n", block.HashQB)
		fmt.Printf("\t数据：%s\n", block.DataQB)
		fmt.Printf("\t随机数：%d\n", block.NonceQB)
		//fmt.Printf("\t时间：%d\n", block.TimeStamp)
		fmt.Printf("\t时间：%s\n", time.Unix(block.TimeStampQB, 0).Format("2006-01-02 15:04:05")) // 时间戳-->time-->Format("")

		//step2：判断block的prevBlcokhash为0,表示该block是创世取块，将诶数循环
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHashQB)
		if big.NewInt(0).Cmp(hashInt) == 0 {
			/*
			x.Cmp(y)
				-1 x < y
				0 x = y
				1 x > y
			 */
			break
		}

	}
}

//获取blockchainiterator的对象
func (QB *BlockChainQB) Iterator() *BlockChainIterator {
	return &BlockChainIterator{QB.DB, QB.Tip}
}
