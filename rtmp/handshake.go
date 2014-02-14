// The MIT License (MIT)
//
// Copyright (c) 2014 winlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package rtmp

import (
	"net"
	"fmt"
)

func SimpleHandshake(conn *net.TCPConn) (err error) {
	c0c1 := make([]byte, 1537)
	// TODO: FIXME: read in block mode.
	nsize, err := conn.Read(c0c1)
	if err != nil {
		return
	}
	fmt.Println("read c0c1, size=", nsize)

	s0s1s2 := make([]byte, 3073)
	copy(s0s1s2[0:1537], c0c1)
	// TODO: FIXME: write in block mode.
	nsize, err = conn.Write(s0s1s2)
	if err != nil {
		return
	}
	fmt.Println("write s0s1s2, size=", nsize)

	c2 := make([]byte, 1536)
	nsize, err = conn.Read(c2)
	if err != nil {
		return
	}
	fmt.Println("read c2, size=", nsize)

	return
}
