package main

import (
	"crypto/sha256"
	"fmt"
)

/*
1.定义结构（区块头的字段比正常的少）
  1.前区块哈希
  2.当前区块哈希
  3.数据
2.创建区块
3.生成哈希
4.引入区块链
5.添加区块
6.重构代码


 */

type Block struct {
	PrevBlockHash  []byte  //前区块哈希
	Hash  []byte  //当前区块哈希
	Data  []byte  //数据，目前使用字节流

}
//创建区块，对block的每一个字段填充即可
func NewBlock(data string,prevBlockHash []byte) *Block {
	block:=Block{
		PrevBlockHash:prevBlockHash,
		Data:[]byte(data),
		Hash:[]byte{},

	}
	block.SetHash()

	return &block
}
//为了生成区块哈希，我们自己实现一个简单的函数，来计算哈希，无随机值，没有难度
func (block *Block)SetHash()  {
	var data []byte
	data=append(data,block.PrevBlockHash...)
	data=append(data,block.Data...)
	hash:=sha256.Sum256(data)
	block.Hash=hash[:]

}

//使用block数组模拟创建区块链
type BlockChain struct {
	Blocks []*Block
}

//实现创建区块链的方法
func NewBlockChain() *BlockChain {
	//在创建的时候添加区块，创世块
	genesisblock:=NewBlock("lueluelue",[]byte{0x0000000000000000})
	bc:=BlockChain{Blocks:[]*Block{genesisblock}}
	return &bc

}
//添加区块
func (bc *BlockChain)AddBlock(data string)  {
	//创建区块

	//bc.Blocks最后一个区块的Hash就是当前新区块的PrevHash
	lastBlock:=bc.Blocks[len(bc.Blocks)-1]
	prevHash:=lastBlock.Hash
	block:=NewBlock(data,prevHash)
	//添加到Block数组中
	bc.Blocks=append(bc.Blocks, block)
}



func main()  {
	fmt.Printf("hello world\n")

	//block:=NewBlock("lueluelue",[]byte{0x0000000000000000})
	bc:=NewBlockChain()
	bc.AddBlock("你说啥")

	for _,block:=range bc.Blocks{
		fmt.Printf("-------------------------------------------\n")
		fmt.Printf("PrevBlockHash:%x\n",block.PrevBlockHash)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Printf("Data:%s\n",block.Data)

	}


}


