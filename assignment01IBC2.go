package assignment01IBC2

import (
	"fmt"
	"crypto/sha256"
)


type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func asSha256(o interface{}) string {
    h := sha256.New()
    h.Write([]byte(fmt.Sprintf("%v", o)))

    return fmt.Sprintf("%x", h.Sum(nil))
}


func CalculateHash(inputBlock *Block) string {
	r := asSha256(inputBlock)

	return r
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {
  var prevH string = ""
	if chainHead != nil {
    prevH = chainHead.currentHash
  }
	newBlock := Block {
		transactions: transactionsToInsert,
		prevPointer: chainHead,
		prevHash: prevH,
		currentHash: "",
	}
	newBlock.currentHash = CalculateHash(&newBlock)
	return &newBlock

}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
  var c *Block = chainHead;
	var check int = 0;

	for c != nil {
	  for i:= 0; i < len(c.transactions ); i++ {
		  if c.transactions[i] == oldTrans {
			  c.transactions[i]=newTrans
        check = 1
			  break
		  }
    }
    if (check==1) {
        break
    }
		c = c.prevPointer
	}

}

func ListBlocks(chainHead *Block) {
	var c *Block = chainHead
  var i int = 1

	for c != nil {
		fmt.Println("Block " , i, c.transactions)
		c = c.prevPointer
    i = i+1
	}
  fmt.Println()
}

func VerifyChain(chainHead *Block) {
	var c *Block = chainHead
	var p *Block = c.prevPointer
	var check int = 0

	for c != nil{
    if p == nil {
      if c.prevHash != CalculateHash(c) {
        check = 1
        break
      }
    } else {
        if c.prevHash != p.currentHash {
			  check = 1
			  break
      }
		}
		c = p
		p = c.prevPointer

	}
	if check==1 {
			fmt.Printf("This is not a valid Blockchain!")
		} else {
			fmt.Println("This is a valid Blockchain!")
		}
}
