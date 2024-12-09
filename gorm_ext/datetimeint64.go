package gorm_ext

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"gorm.io/gorm/schema"
)

// 自定义类型，用于序列化 datetime -> int64
type TimestampInt64Serializer struct{}

// Scan 实现 sql.Scanner 接口，用于从数据库读取
func (TimestampInt64Serializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	if dbValue != nil {
		var result int64
		switch v := dbValue.(type) {
		case time.Time:
			result = v.Local().UnixMilli()
		case string:
			var t time.Time
			if t, err = time.Parse("2006-01-02 15:04:05", v); err != nil {
				return
			}
			result = t.Local().UnixMilli()
		case int64:
			result = v
		case int32:
			result = int64(v)
		case int:
			result = int64(v)
		default:
			return fmt.Errorf("failed to unmarshal timestamp value: %#v", dbValue)
		}

		if result < 0 {
			result = 0
		}

		field.Set(ctx, dst, result)
	}

	return
}

// Value 实现 driver.Valuer 接口，用于写入数据库
func (TimestampInt64Serializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return time.Unix(fieldValue.(int64), 0), nil // 将时间戳转为 time.Time 类型
}
