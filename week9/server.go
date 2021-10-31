package week9

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"runtime/debug"
	"time"

	"github.com/pkg/errors"
)

// 字节数(大端)组转成int(有符号)
func bytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

func handleConn(conn net.Conn) {
	_ = conn.SetDeadline(time.Now().Add(time.Second))
	r := bufio.NewReader(conn)
	buffs := make([]byte, 0, 1024)
	msg := make([]byte, 0, 1024)
	pl := 0  // package length
	hl := 0  // header length
	pv := 0  // protocol version
	op := 0  // operation
	sid := 0 // sequence id
	for {
		if len(buffs) >= pl && pl > 0 {
			// 已经获取到了完整的包
			msg = buffs[pl-hl : pl]
			fmt.Printf("sid: %v, op: %v, pv: %v, msg: %v", sid, op, pv, string(msg))
			_, _ = conn.Write([]byte("ok"))

			msg = make([]byte, 0, 1024)
			if len(buffs) > pl {
				buffs = buffs[pl+1:]
			} else {
				buffs = make([]byte, 0, 1024)
			}

			pl = 0
			hl = 0
			pv = 0
			op = 0
			sid = 0
		}
		var buf [1024]byte
		n, err := r.Read(buf[:])
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Printf("Read data from %s err: %v", conn.RemoteAddr(), err)
			break
		}

		buffs = append(buffs, buf[:n]...)

		fmt.Printf("receive data from %s. \n data: %v \n", conn.RemoteAddr(), string(buffs))

		if len(buffs) >= 16 {
			pl, _ = bytesToIntS(buffs[:5])
			hl, _ = bytesToIntS(buffs[5:7])
			pv, _ = bytesToIntS(buffs[7:9])
			op, _ = bytesToIntS(buffs[9:13])
			sid, _ = bytesToIntS(buffs[13:17])
		}
	}
}

func Serve() {
	address := "127.0.0.1:9090"
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("listen %s err: %v", address, err)
		panic("listen err")
	}

	fmt.Printf("begin to listen %s", address)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("accept a connection err: %v", err)
			continue
		}

		go func(c net.Conn) {
			defer func() {
				if e := recover(); e != nil {
					fmt.Printf("handle this connection err: %v, %v", e, string(debug.Stack()))
				}
				_ = c.Close()
			}()
			handleConn(c)
		}(conn)
	}
}
