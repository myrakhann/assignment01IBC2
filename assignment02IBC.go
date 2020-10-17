package assignment02IBC

import (
	"fmt"
	"crypto/sha256"
)

type Block struct {
	Spender     map[string]int
	Receiver    map[string]int
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

const miningReward = 100
const rootUser = "Satoshi"

func CalculateBalance(userName string, chainHead *Block) int {
  var c *Block = chainHead
 
	for c != nil {
    _, foundS := c.Spender[userName]
    _, foundR := c.Receiver[userName]
	  if foundS {
      return c.Spender[userName]
    } else if foundR {
      return c.Receiver[userName]
    } 
		c = c.PrevPointer
	}
  return 0
  
  
}

func asSha256(o interface{}) string {
    h := sha256.New()
    h.Write([]byte(fmt.Sprintf("%v", o)))

    return fmt.Sprintf("%x", h.Sum(nil))
}


func CalculateHash(inputBlock *Block) string {
	return asSha256(inputBlock) 
}


func InsertBlock(spendingUser string, receivingUser string, miner string, amount int, chainHead *Block) *Block {

  if miner != rootUser {
    fmt.Println("Error: Only Satoshi can be a miner!\n")
    return chainHead
  }
  
  bal := CalculateBalance(spendingUser, chainHead)

  if bal < amount {
    fmt.Println("Error: This spender does not have enough balance!\n")
    return chainHead
  }

  var prevH string = ""
	if chainHead != nil {
    prevH = chainHead.CurrentHash
  }

	newBlock := Block {
		PrevPointer: chainHead,
		PrevHash: prevH,
		CurrentHash: "",
	}
	newBlock.CurrentHash = CalculateHash(&newBlock)	
  
  newBlock.Spender = make (map[string]int)
  newBlock.Receiver = make (map[string]int)

  if spendingUser != "" && receivingUser != "" {
    newBlock.Spender[spendingUser] = bal - amount
    newBlock.Receiver[receivingUser] = CalculateBalance(receivingUser, chainHead) + amount
  }
    newBlock.Receiver[miner] = CalculateBalance(miner, chainHead) + miningReward
  
	return &newBlock
}




func ListBlocks(chainHead *Block) {
	var c *Block = chainHead
  var i int = 1
	
	for c != nil {
		fmt.Println("Block " , i)
    fmt.Println("   Spender: ", c.Spender) 
    fmt.Println("   Receiver: ", c.Receiver)
		c = c.PrevPointer
    i = i+1
	}
  fmt.Println()
}

func VerifyChain(chainHead *Block) {
	var c *Block = chainHead
	var p *Block = c.PrevPointer
	var check int = 0
	
	for c != nil{
    if p == nil {
      if c.PrevHash != CalculateHash(c) {
        check = 1
        break
      }
    } else {
        if c.PrevHash != p.CurrentHash {
			  check = 1
			  break	
      }		
		}
		c = p
		p = c.PrevPointer
			
	}
	if check==1 {
		fmt.Printf("This is not verified! ")
    }
}

