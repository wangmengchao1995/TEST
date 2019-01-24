package main

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

