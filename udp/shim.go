package udp

import "encoding/binary"

func paddingHeaderSign(data []byte) []byte {
	var padded []byte
	header := []byte("NSUCV1")
	length := uint16(len(data))
	byteLen := make([]byte, 2)
	binary.LittleEndian.PutUint16(byteLen, length)
	padded = append(padded, header...)
	padded = append(padded, byteLen...)
	padded = append(padded, data...)
	return padded
}
