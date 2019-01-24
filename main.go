package main

import "fmt"




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


