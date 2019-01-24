package main

import "fmt"


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

	return &block
}



func main()  {
	fmt.Printf("hello world\n")

	block:=NewBlock("lueluelue",[]byte{0x0000000000000000})
	fmt.Printf("PrevBlockHash:%x\n",block.PrevBlockHash)

	fmt.Printf("Hash:%x\n",block.Hash)
	fmt.Printf("Data:%s\n",block.Data)


}


