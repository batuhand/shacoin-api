package models

type SealedBlock struct {
	BlockID   int         `json:"block_id"`
	PrevHash  string      `json:"prev_hash"`
	Data      Transaction `json:"data"`
	BlockHash string      `json:"block_hash"`
}
