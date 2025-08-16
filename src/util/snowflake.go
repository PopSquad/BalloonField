package util

import (
	"sync"
	"time"
)

func SnowflakeGenerator(machineID int64) func() int64 {
	var (
		mutex        sync.Mutex
		lastStamp    int64
		sequence     int64
		epoch        int64 = 1609459200000
		machineBits  int64 = 10
		seqBits      int64 = 12
		maxMachine         = -1 ^ (-1 << machineBits)
		maxSeq             = -1 ^ (-1 << seqBits)
		timeShift          = machineBits + seqBits
		machineShift       = seqBits
	)

	if machineID < 0 || machineID > int64(maxMachine) {
		panic("machineID out of range")
	}

	return func() int64 {
		mutex.Lock()
		defer mutex.Unlock()

		now := time.Now().UnixNano() / 1e6
		if now == lastStamp {
			sequence = (sequence + 1) & int64(maxSeq)
			if sequence == 0 {
				for now <= lastStamp {
					now = time.Now().UnixNano() / 1e6
				}
			}
		} else {
			sequence = 0
		}
		lastStamp = now

		id := ((now - epoch) << timeShift) | (machineID << machineShift) | sequence
		return id
	}
}

var GenID = SnowflakeGenerator(1)
