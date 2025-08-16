package plaza

import (
	"bytes"
	"encoding/binary"

	"github.com/PopSquad/BalloonField/src/util"
)

const (
	CMD96    byte = 0x96 // 登录服务
	PRM96_01 byte = 0x01 // 登录请求
	PRM96_64 byte = 0x64 // 登录响应

	CMD97    byte = 0x97 // 网关服务
	PRM97_64 byte = 0x64 // 服务器节点列表
	PRM97_65 byte = 0x65 // 公告?
	PRM97_66 byte = 0x66
	PRM97_67 byte = 0x67 // Echo
	PRM97_00 byte = 0x00 // Echo Response
)

type SvrNodeInfo struct {
	Byte0          uint8
	Word1          uint16
	ChannelName    [25]byte
	Dword1C        uint32
	Word20         uint16
	Word22         uint16
	RoomServerIP   uint32
	RoomServerPort uint16
	Dword2A        uint32
	Dword2E        uint32
	Dword32        uint32
	UShort36       uint16
	Word38         uint16
	DWord42        uint32
	RegionName     [25]byte
}

func BuildParam97_64Data(cmdOrSessionID uint32, records []SvrNodeInfo) []byte {
	buf := new(bytes.Buffer)

	_ = binary.Write(buf, binary.LittleEndian, cmdOrSessionID)                     // 0x00-0x03 ukn
	_ = binary.Write(buf, binary.LittleEndian, uint16(binary.Size(SvrNodeInfo{}))) // 0x04-0x05 recordSize
	_ = buf.WriteByte(uint8(len(records)))                                         // 0x06 recordCount

	for _, rec := range records {
		_ = binary.Write(buf, binary.LittleEndian, rec)
	}

	return buf.Bytes()
}

func MockRecord() SvrNodeInfo {
	localhost, _ := util.IPv4ToUint32("127.0.0.1")
	regionName, _ := util.Str2GBK("区域1")
	channelName, _ := util.Str2GBK("频道1")
	var r SvrNodeInfo
	r.Byte0 = 1
	r.Word1 = 2
	copy(r.ChannelName[:], channelName)
	r.Dword1C = 3
	r.Word20 = 4
	r.Word22 = 5
	r.RoomServerIP = localhost
	r.RoomServerPort = 9000
	r.Dword2A = 6
	r.Dword2E = 7
	r.Dword32 = 8
	r.UShort36 = 9
	r.Word38 = 10
	r.DWord42 = 11
	copy(r.RegionName[:], regionName) // TODO: 前两个byte会被截断
	return r
}
