package room

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/PopSquad/BalloonField/src/util"
)

const (
	CMD00       byte = 0x00 // 登录服务
	PRM00_02    byte = 0x02 // 登录请求
	PRM00_64    byte = 0x64 // 登录响应
	DAT00_64_00 byte = 0x00 // 登陆响应数据 - 登陆成功
	DAT00_64_01 byte = 0x01 // 登陆响应数据 - 用户名错误
	DAT00_64_02 byte = 0x02 // 登陆响应数据 - 密码错误
	DAT00_64_03 byte = 0x03 // 登陆响应数据 - 服务器不存在
	DAT00_64_04 byte = 0x04 // 登陆响应数据 - 请重新登录广场服务器
	DAT00_64_05 byte = 0x05 // 登陆响应数据 - 请重新登录广场服务器(是的他跟上一个是同样用途)
	DAT00_64_06 byte = 0x06 // 登陆响应数据 - 服务器满
	DAT00_64_07 byte = 0x07 // 登陆响应数据 - 客户端版本低
	DAT00_64_08 byte = 0x08 // 登陆响应数据 - 客户端版本高
	DAT00_64_09 byte = 0x09 // 登陆响应数据 - 服务器内部错误
	DAT00_64_0A byte = 0x0A // 登陆响应数据 - 重复的登录尝试
	DAT00_64_0B byte = 0x0B // 登陆响应数据 - 登录出错

	CMD01    byte = 0x01 // 房间列表服务？
	PRM01_64 byte = 0x64
	PRM01_65 byte = 0x65
	PRM01_66 byte = 0x66
	PRM01_67 byte = 0x67
	PRM01_68 byte = 0x68 // 加入房间？
	PRM01_69 byte = 0x69
	PRM01_6A byte = 0x6A
	PRM01_6B byte = 0x6B

	CMD02    byte = 0x02 // 房间内服务？
	PRM02_64 byte = 0x64
	PRM02_66 byte = 0x66
	PRM02_67 byte = 0x67
	PRM02_68 byte = 0x68
	PRM02_69 byte = 0x69
	PRM02_6A byte = 0x6A

	CMD04    byte = 0x04 // 未知的default分支
	PRM04_64 byte = 0x64

	CMD06    byte = 0x06
	PRM06_64 byte = 0x64
	PRM06_65 byte = 0x65
	PRM06_66 byte = 0x66
	PRM06_67 byte = 0x67
	PRM06_68 byte = 0x68 // 商城购买
	PRM06_69 byte = 0x69
	PRM06_6A byte = 0x6A
	PRM06_6B byte = 0x6B
)

type LoginRespPacket struct {
	ResultCode uint8 // 登录结果代码
	Unk1       uint8 // 未知字段（可能是错误码高位）
	Unk2       uint32
	Unk3       [4]byte
	Unk4       [17]byte
	MsgLen     uint16 // 附加消息长度
	Msg        string // 附加消息
}

func BuildLoginRespPacket(pkt LoginRespPacket) []byte {
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, pkt.ResultCode)
	binary.Write(buf, binary.LittleEndian, pkt.Unk1)
	binary.Write(buf, binary.LittleEndian, pkt.Unk2)
	binary.Write(buf, binary.LittleEndian, pkt.Unk3)
	binary.Write(buf, binary.LittleEndian, pkt.Unk4)

	msgBytes, err := util.Str2GBK(pkt.Msg)
	if err != nil {
		log.Printf("UTF-8 to GBK conversion failed: %v", err)
		msgBytes = []byte{}
	}

	if len(msgBytes) > 2040 {
		msgBytes = msgBytes[:2040]
	}
	pkt.MsgLen = uint16(len(msgBytes))

	binary.Write(buf, binary.LittleEndian, pkt.MsgLen)
	buf.Write(msgBytes)

	return buf.Bytes()
}
