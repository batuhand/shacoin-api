package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"batuhand.com/api/models"
)

func AppendTempFile(data []byte) {

	//Looking for if the file exists
	file, err := os.OpenFile("blocks/temp.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		//If file don't exists, create one
		err := ioutil.WriteFile("blocks/temp.json", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := file.WriteString(string(data)); err != nil {
			log.Fatal(err)
		}

	}
	defer file.Close()

}

func WriteSealedBlock(block models.SealedBlock) {
	blockData, _ := json.Marshal(block)
	blockID := GetLastBlockID() + 1
	err := ioutil.WriteFile("blocks/block-"+strconv.Itoa(blockID)+".json", blockData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckTempFileSize() {
	fi, err := os.Stat("blocks/temp.json")
	if err != nil {
		log.Fatal(err)
	}
	// get the size
	size := fi.Size()
	fmt.Println(size)
	if size > 1 {
		WriteSealedBlock(SealBlock())
	}
}

func SealBlock() models.SealedBlock {
	tempData := ReadFile("blocks/temp.json")
	transaction := models.Transaction{}
	json.Unmarshal(tempData, &transaction)
	fmt.Println("SA---------------")
	blockID := GetLastBlockID() + 1
	prevHash := GetLastBlockHash()
	block := models.Block{
		BlockID:  blockID,
		PrevHash: prevHash,
		Data:     transaction,
	}
	hash := asSha256(block)
	if err := os.Truncate("blocks/temp.json", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
	sealedBlock := models.SealedBlock{
		BlockID:   blockID,
		PrevHash:  prevHash,
		Data:      transaction,
		BlockHash: hash,
	}
	return sealedBlock
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func ReadFile(fileName string) []byte {
	//Print the contents of the file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	return data
}

func GetLastBlockID() int {
	biggestBlockID := 0
	files, err := ioutil.ReadDir("./blocks/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Name() == "temp.json" {
			continue
		} else {
			blockID, _ := strconv.Atoi(string([]rune(strings.Split(f.Name(), "-")[1])[0]))
			fmt.Println(blockID)
			if blockID > biggestBlockID {
				biggestBlockID = blockID
			}

		}
	}
	return biggestBlockID
}

// Returns the currency of the given wallet
func GetWalletAmount(walletId string) float64 {
	var currency float64
	files, err := ioutil.ReadDir("./blocks/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Name() == "temp.json" {
			continue
		} else {
			blockID, _ := strconv.Atoi(string([]rune(strings.Split(f.Name(), "-")[1])[0]))
			lastBlockData := ReadFile("blocks/block-" + strconv.Itoa(blockID) + ".json")
			block := models.SealedBlock{}
			err := json.Unmarshal(lastBlockData, &block)
			if err != nil {
				log.Fatal(err)

			} else if block.Data.RecieverWalletAdress == walletId {
				currency = currency + block.Data.Amount
			} else if block.Data.SenderWalletAdress == walletId {
				currency = currency - block.Data.Amount
			}
		}
	}
	fmt.Println(currency)
	return currency
}

func GetLastBlockHash() string {
	lastBlockID := GetLastBlockID()
	if lastBlockID > 0 {
		blockFileName := "block-" + strconv.Itoa(lastBlockID) + ".json"
		lastBlockData := ReadFile("blocks/" + blockFileName)
		block := models.SealedBlock{}
		err := json.Unmarshal(lastBlockData, &block)
		if err != nil {
			log.Fatal(err)
			return ""
		} else {
			fmt.Println(block.BlockHash)
			return block.BlockHash
		}
	} else {
		return ""
	}

}
