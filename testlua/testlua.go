package main

import (
	lua "lua51/internal"
)

func main() {
	println("hello wrld")
	s := lua.Open()
	t := s.GetTop()
	println("top=", t)
	s.PushInteger(5)
	println("push 5")
	println("top=", s.GetTop())
	s.PushInteger(5)
	println("push 5")
	s.PushInteger(0)
	println("push 0")
	println("top=", s.GetTop())
	println("equal(-1,-2)=", s.Equal(-1, -2))
	println("equal(-2,-3)=", s.Equal(-2, -3))
	println("equal(1,2)=", s.Equal(1, 2))
	println("equal(2,3)=", s.Equal(2, 3))

	r := s.Loadbuffer([]byte(
		// "return 2+3", generated by '../gen_chunk.lua'
		"\x1bLuaQ\x00\x01\x04\x04\x04\x08\x00\x01\x00\x00\x00\x00\x00"+
			"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x02\x02\x03\x00\x00\x00"+
			"\x01\x00\x00\x00\x1e\x00\x00\x01\x1e\x00\x80\x00\x01\x00\x00"+
			"\x00\x03\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x03\x00"+
			"\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00"+
			"\x00\x00\x00\x00\x00\x00\x00"),
		"chunk 1")
	if r != 0 {
		panic(r)
	}

	s.Call(0, 1)
	println("call")
	println("top=", s.GetTop())
	println("equal(-1,1)=", s.Equal(-1, 1)) // expected: yes (i.e. peek()==5)
}