package arpc

import (
	"log"
	"time"

	"github.com/sony/sonyflake"
)

// 雪花ID生成器
var Sf *sonyflake.Sonyflake
var SfSettings = sonyflake.Settings{
	StartTime: time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC), // 开始时间
}

func InitSnowID() {
	Sf = sonyflake.NewSonyflake(SfSettings)
}

// 生成雪花 ID
func NewSnowID() uint64 {
	id, err := Sf.NextID()
	if err != nil {
		log.Fatalf("Failed to generate new Snow ID: %v", err)
	}
	return id
}

// 雪花 ID 解析时间
func SnowIDTime(id uint64) time.Time {
	decomposed := sonyflake.Decompose(id)
	// 雪花id中的时间是以 10ms 为单位的距离开始时间的时长
	return SfSettings.StartTime.Add(time.Duration(decomposed["time"]*10) * time.Millisecond)
}
