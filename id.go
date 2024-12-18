package arpc

import (
	"log"

	"github.com/sony/sonyflake"
)

// 雪花ID生成器
var Sf *sonyflake.Sonyflake

func InitSnowID() {
	Sf = sonyflake.NewSonyflake(sonyflake.Settings{})
}

// 生成雪花 ID
func NewSnowID() uint64 {
	id, err := Sf.NextID()
	if err != nil {
		log.Fatalf("Failed to generate new Snow ID: %v", err)
	}
	return id
}
