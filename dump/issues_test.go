package dump_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/gookit/goutil/dump"
)

func isTimeType(v interface{}) bool {
	return reflect.TypeOf(v).ConvertibleTo(reflect.TypeOf(time.Time{}))
}

// https://github.com/gookit/goutil/issues/200
func TestIssues_200(t *testing.T) {
	var dateValue = "2024-10-02T19:02:46"
	type RestorePointTimestamp time.Time
	type NormalRestore struct {
		TimeStamp time.Time
	}
	type CustomRestore struct {
		TimeStamp RestorePointTimestamp
	}

	t.Run("normal time", func(t *testing.T) {
		normal, _ := time.Parse("2006-01-02T15:04:05", dateValue)
		normalRestore := NormalRestore{
			TimeStamp: normal,
		}
		dump.Print(normalRestore)
	})

	t.Run("custom time", func(t *testing.T) {
		custom, _ := time.Parse("2006-01-02T15:04:05", dateValue)
		customRestore := CustomRestore{
			TimeStamp: RestorePointTimestamp(custom),
		}
		fmt.Println(isTimeType(customRestore.TimeStamp))
		dump.Print(customRestore)
	})
}
