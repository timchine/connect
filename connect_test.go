package connect

import (
	"testing"
)

func TestNewExternalProcedure(t *testing.T) {
	NewExternalProcedure(NewMysqlConfig("127.0.0.1", "3307", "root", "123456", "sweet"), NewRedisConfig("127.0.0.1", "6379", ""))
}