package service

import (
	cond "github.com/vela-security/vela-cond"
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/export"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

func lookupL(L *lua.LState) int {
	cnd := cond.CheckMany(L)
	su := New()
	su.collect(cnd)
	L.Push(su)
	return 1
}

func indexL(L *lua.LState, key string) lua.LValue {
	cnd := cond.New("name = " + key)
	su := New()
	su.collect(cnd)
	return su
}

func WithEnv(env assert.Environment) {
	xEnv = env
	xEnv.Set("service", export.New("vela.service.export", export.WithFunc(lookupL), export.WithIndex(indexL)))
}
