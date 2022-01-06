package models

type Block struct{
	BlockID int
	PrevHash string
	Data Transaction
}