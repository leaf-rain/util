package tool

/*
	雪花算法组成部分(64bit):

	0(1位，且始终为0)|时间戳(41位)|工作机器id(10位)|序列号(12位)
*/

import (
	"math/rand"
	"time"
)

var (
	// Epoch 2021年3月3号0:00时刻的毫秒级时间戳
	Epoch         int64 = 1614700800000
	machineID     int64 // 机器ID
	sn            int64 // 序列号
	lastTimeStamp int64 // 记录上次的时间戳(毫秒级)
)

func init() {
	lastTimeStamp = time.Now().UnixNano()/1e6 - Epoch
	machineID = rand.Int63n(1000000000)
}

// SetMachineID is 设置机器ID
func SetMachineID(mid int64) {
	machineID = mid << 12
}

// GenUUID 生成 UUID
func GenUUID() uint64 {
	// 单位为毫秒
	curTimeStamp := time.Now().UnixNano()/1e6 - Epoch
	if curTimeStamp == lastTimeStamp {
		sn++
		//序列号为12位， 2^12 = 4096个
		if sn > 4095 {
			//序列号超出，则重置序列号。这也意味着每毫秒最多能生成4096个id值
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano()/1e6 - Epoch
			lastTimeStamp = curTimeStamp //顺便更新下上次的时间戳
			sn = 0
		}
		//与运算 对应位全为1时，则为1.否则为0
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22

		//或运算 对应位全为0时，则为0。否则为1
		id := rightBinValue | machineID | sn
		return uint64(id)
	} else if curTimeStamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = curTimeStamp
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		return uint64(rightBinValue | machineID | sn)
	}
	return 0
}
