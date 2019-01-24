package main

import (
	"crypto/sha256"
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