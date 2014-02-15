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

import "net"

type RtmpSocket struct {
	conn *net.TCPConn
	RecvBytes uint64
	SendBytes uint64
}

func NewRtmpSocket(conn *net.TCPConn) (r *RtmpSocket) {
	r = new(RtmpSocket)
	r.conn = conn
	return
}

func (r *RtmpSocket) Read(b []byte) (n int, err error) {
	n, err = r.conn.Read(b)

	if n > 0 {
		r.RecvBytes += uint64(n)
	}

	return
}

func (r *RtmpSocket) Write(b []byte) (n int, err error) {
	n, err = r.conn.Write(b)

	if n > 0 {
		r.SendBytes += uint64(n)
	}

	return
}
