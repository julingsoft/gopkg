package xguid

import (
	"log"
	"time"

	"github.com/sony/sonyflake/v2"
)

func New() (int64, error) {
	st := sonyflake.Settings{
		TimeUnit: time.Millisecond,
	}

	s, err := sonyflake.New(st)
	if err != nil {
		return 0, err
	}

	return s.NextID()
}

func NextID() int64 {
	id, err := New()
	if err != nil {
		log.Fatalf("failed to generate id: %v", err)
	}
	return id
}
