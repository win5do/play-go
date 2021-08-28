package mockhttp

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"strconv"
)

var RN = []byte("\r\n")

func ReadConn(conn net.Conn) (bytes.Buffer, error) {
	var buf bytes.Buffer
	var bodyLen int

	rd := bufio.NewReader(conn)

	for {
		line, err := rd.ReadBytes('\n')
		if err != nil {
			return buf, err
		}

		buf.Write(line)
		if len(line) <= 2 {
			break
		}

		i := bytes.IndexByte(line, ':')
		if i < 0 {
			continue
		}
		key := line[:i]
		val := bytes.TrimSpace(line[i+1 : len(line)-2])
		log.Printf("header key: %s, val: %s", key, val)
		if bytes.Equal(key, []byte("Content-Length")) {
			x, err := strconv.Atoi(string(val))
			if err != nil {
				return buf, err
			}
			bodyLen = x
		}
	}

	header := buf.String()
	body := make([]byte, bodyLen)
	_, err := rd.Read(body)
	if err != nil {
		return buf, err
	}
	buf.Write(body)

	log.Printf("Header: \n%s", header)
	log.Printf("Body：\n%s", string(body))

	return buf, nil
}
