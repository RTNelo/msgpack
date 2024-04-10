package common

import (
	"reflect"
	"unsafe"
)

type typeface struct {
	metaptr uintptr
	data    unsafe.Pointer
}

func Type2rtypeptr(t reflect.Type) unsafe.Pointer {
	return (*typeface)(unsafe.Pointer(&t)).data
}
