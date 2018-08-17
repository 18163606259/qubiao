package BLC

import (
	"math/big"
	"crypto/sha256"
	"bytes"
	"fmt"
)

//工作量证明：
//step1：构建Pow类型

const TargetBit = 16 //目标哈希的0个个数,16,20,24,28

type ProofOfWorkQB struct {
	BlockQB  *BlockQB   //要验证的block
	TargetQB *big.Int //目标hash  0000 0000 0000 0000 1000 0000 0000 0000 0000 0000 ... 256bit
}

//step2，提供一个pow对象
func NewProofOfWorkQB(block *BlockQB) *ProofOfWorkQB {
	//1.创建pow对象
	pow := &ProofOfWorkQB{}
	//2.设置属性值
	pow.BlockQB = block
	targetQB := big.NewInt(1)           //target的类型：bigInt,数值1        // 目标hash，初始值为1
	targetQB.Lsh(targetQB, 256-TargetBit) //左移256-16
	pow.TargetQB = targetQB
	/*
	hash：256bit
	16进制：32个
	4个0--->16个0
	0000 0000 0000 0000 1000 0000 0000 0000    256
	 */

	/*
	0000 0001
	0010 0000

	8-2
	256-16
	 */
	return pow

}

//step3：设计一个函数，得到有效hash，nonce
func (pow *ProofOfWorkQB) Run() ([]byte, int64) {
	/*
	1.将block的字段属性，拼接成一个数组
	2.定义一个nonce的值：初始值为1，
	3.产生hash--->和目标hash比较，
	 */
	//A: 定义一个nonce随机数
	var nonce int64 = 1

	var hash [32]byte
	for {

		//B： 获取拼接后的字节数组
		dataBytes := pow.prepareDataQB(nonce)

		//C:产生hash
		hash = sha256.Sum256(dataBytes) //[32]byte
		fmt.Printf("\r%d,%x", nonce, hash)

		//D:比较hash和目标hash
		/*
		hash：[]byte --->big.Int
		目标hash：pow.target-->big.Int

			func (x *Int) Cmp(y *Int) (r int)
				Cmp compares x and y and returns:

					-1 if x <  y
					0 if x == y
					+1 if x >  y
	 */

		hashInt := new(big.Int)
		hashInt.SetBytes(hash[:])

		if pow.TargetQB.Cmp(hashInt) == 1 {
			break
		}

		nonce++

	}

	return hash[:], nonce

}

/*
根据block的字段属性，以及传来的nonce值，拼接成一个字节数组
 */
func (pow *ProofOfWorkQB) prepareDataQB(nonce int64) []byte {
	data := bytes.Join([][]byte{
		pow.BlockQB.PrevBlockHashQB,
		pow.BlockQB.DataQB,
		IntToHex(pow.BlockQB.TimeStampQB),
		IntToHex(pow.BlockQB.HeightQB),
		IntToHex(nonce)}, []byte{})
	return data
}

//提供一个方法：
func (pow *ProofOfWorkQB) IsValidQB() bool {
	hashInt := new(big.Int)
	hashInt.SetBytes(pow.BlockQB.HashQB)
	return pow.TargetQB.Cmp(hashInt) == 1
}
