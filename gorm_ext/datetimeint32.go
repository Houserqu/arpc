package gorm_ext

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"gorm.io/gorm/schema"
)

// 自定义类型，用于序列化 datetime -> int64
type TimestampInt32Serializer struct{}

// Scan 实现 sql.Scanner 接口，用于从数据库读取
func (TimestampInt32Serializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	if dbValue != nil {
		var result int32
		switch v := dbValue.(type) {
		case time.Time:
			result = int32(v.Local().Unix())
		case string:
			var t time.Time
			if t, err = time.Parse("2006-01-02 15:04:05", v); err != nil {
				return
			}
			result = int32(t.Local().Unix())
		case int64:
			result = int32(v)
		case int32:
			result = v
		case int:
			result = int32(v)
		default:
			return fmt.Errorf("failed to unmarshal timestamp value: %#v", dbValue)
		}

		field.Set(ctx, dst, result)
	}

	return
}

// Value 实现 driver.Valuer 接口，用于写入数据库
func (TimestampInt32Serializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return time.Unix(int64(fieldValue.(int32)), 0), nil
}
