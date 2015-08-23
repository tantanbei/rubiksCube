package main

import (
	"fmt"
)

type FaceRubiksCube struct {
	items []byte
	angle int
}

type UpRubiksCube struct {
	FaceRubiksCube
}
type DownRubiksCube struct {
	FaceRubiksCube
}
type ForwardRubiksCube struct {
	FaceRubiksCube
}
type BackRubiksCube struct {
	FaceRubiksCube
}
type LeftRubiksCube struct {
	FaceRubiksCube
}
type RightRubiksCube struct {
	FaceRubiksCube
}

func TranLeftFace(l *LeftRubiksCube, u *UpRubiksCube, f *ForwardRubiksCube, b *BackRubiksCube, d *DownRubiksCube) (*LeftRubiksCube, *UpRubiksCube, *ForwardRubiksCube, *BackRubiksCube, *DownRubiksCube) {
	buf := make([]byte, 9)
	copy(buf, l.items)
	for i := 3; i < 9; i++ {
		l.items[i] = buf[i-2]
	}
	l.items[2] = buf[8]
	l.items[1] = buf[7]

	uBuf := make([]byte, 3)
	copy(uBuf, u.items[1:4])
	fBuf := make([]byte, 3)
	copy(fBuf, f.items[1:4])
	bBuf := make([]byte, 3)
	copy(bBuf, b.items[1:4])
	dBuf := make([]byte, 3)
	copy(dBuf, d.items[1:4])

	copy(d.items[1:4], fBuf)
	copy(b.items[1:4], dBuf)
	copy(u.items[1:4], bBuf)
	copy(f.items[1:4], uBuf)
	return l, u, f, b, d
}

func TranRightFace(r *RightRubiksCube, u *UpRubiksCube, f *ForwardRubiksCube, b *BackRubiksCube, d *DownRubiksCube) (*RightRubiksCube, *UpRubiksCube, *ForwardRubiksCube, *BackRubiksCube, *DownRubiksCube) {
	buf := make([]byte, 9)
	copy(buf, r.items)
	for i := 3; i < 9; i++ {
		r.items[i] = buf[i-2]
	}
	r.items[2] = buf[8]
	r.items[1] = buf[7]

	uBuf := make([]byte, 3)
	copy(uBuf, u.items[5:8])
	fBuf := make([]byte, 3)
	copy(fBuf, f.items[5:8])
	bBuf := make([]byte, 3)
	copy(bBuf, b.items[5:8])
	dBuf := make([]byte, 3)
	copy(dBuf, d.items[5:8])

	copy(d.items[5:8], bBuf)
	copy(b.items[5:8], uBuf)
	copy(u.items[5:8], fBuf)
	copy(f.items[5:8], dBuf)
	return r, u, f, b, d
}

func TranUpFace(u *UpRubiksCube, r *RightRubiksCube, f *ForwardRubiksCube, b *BackRubiksCube, l *LeftRubiksCube) (*UpRubiksCube, *RightRubiksCube, *ForwardRubiksCube, *BackRubiksCube, *LeftRubiksCube) {
	buf := make([]byte, 9)
	copy(buf, u.items)
	for i := 3; i < 9; i++ {
		u.items[i] = buf[i-2]
	}
	u.items[2] = buf[8]
	u.items[1] = buf[7]

	rBuf := make([]byte, 3)
	copy(rBuf, r.items[3:6])
	fBuf := make([]byte, 3)
	copy(fBuf, f.items[3:6])
	bBuf := make([]byte, 3)
	copy(bBuf, b.items[3:6])
	lBuf := make([]byte, 3)
	copy(lBuf, l.items[3:6])

	copy(r.items[3:6], bBuf)
	copy(f.items[3:6], lBuf)
	copy(b.items[3:6], rBuf)
	copy(l.items[3:6], fBuf)
	return u, r, f, b, l
}

func TranDownFace(d *DownRubiksCube, r *RightRubiksCube, f *ForwardRubiksCube, b *BackRubiksCube, l *LeftRubiksCube) (*DownRubiksCube, *RightRubiksCube, *ForwardRubiksCube, *BackRubiksCube, *LeftRubiksCube) {
	buf := make([]byte, 9)
	copy(buf, d.items)
	for i := 3; i < 9; i++ {
		d.items[i] = buf[i-2]
	}
	d.items[2] = buf[8]
	d.items[1] = buf[7]

	rBuf := make([]byte, 3)
	copy(rBuf, r.items[3:6])
	fBuf := make([]byte, 3)
	copy(fBuf, f.items[3:6])
	bBuf := make([]byte, 3)
	copy(bBuf, b.items[3:6])
	lBuf := make([]byte, 3)
	copy(lBuf, l.items[3:6])

	copy(r.items[3:6], bBuf)
	copy(b.items[3:6], lBuf)
	copy(f.items[3:6], rBuf)
	copy(l.items[3:6], fBuf)
	return d, r, f, b, l
}

func main() {
	U := &UpRubiksCube{FaceRubiksCube{items: []byte{'w', 'w', 'w', 'w', 'w', 'w', 'w', 'w', 'w'}, angle: 0}}
	D := &DownRubiksCube{FaceRubiksCube{items: []byte{'y', 'y', 'y', 'y', 'y', 'y', 'y', 'y', 'y'}, angle: 0}}
	F := &ForwardRubiksCube{FaceRubiksCube{items: []byte{'r', 'r', 'r', 'r', 'r', 'r', 'r', 'r', 'r'}, angle: 0}}
	B := &BackRubiksCube{FaceRubiksCube{items: []byte{'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o', 'o'}, angle: 0}}
	L := &LeftRubiksCube{FaceRubiksCube{items: []byte{'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g', 'g'}, angle: 0}}
	R := &RightRubiksCube{FaceRubiksCube{items: []byte{'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, angle: 0}}
	//fmt.Println(R, L, U, F, B, D)
	var input string
	fmt.Println("Import the command:")
	fmt.Scanf("%s", &input)
	for _, c := range input {
		switch c {
		case 'l':
			L, U, F, B, D = TranLeftFace(L, U, F, B, D)
		case 'r':
			R, U, F, B, D = TranRightFace(R, U, F, B, D)
		case 'u':
			U, R, F, B, L = TranUpFace(U, R, F, B, L)
		case 'd':
			D, R, F, B, L = TranDownFace(D, R, F, B, L)
		}
	}
	fmt.Println(R, L, U, F, B, D)
}
