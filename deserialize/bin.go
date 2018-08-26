package deserialize

import (
	"encoding/binary"
	"reflect"

	"github.com/shamaton/msgpack/def"
)

func (d *deserializer) isCodeBin(v byte) bool {
	switch v {
	case def.Bin8, def.Bin16, def.Bin32:
		return true
	}
	return false
}

func (d *deserializer) asBin(offset int, k reflect.Kind) ([]byte, int, error) {
	code, offset := d.readSize1(offset)

	switch code {
	case def.Bin8:
		l, offset := d.readSize1(offset)
		o := offset + int(uint8(l))
		return d.data[offset:o], o, nil
	case def.Bin16:
		bs, offset := d.readSize2(offset)
		o := offset + int(binary.BigEndian.Uint16(bs))
		return d.data[offset:o], o, nil
	case def.Bin32:
		bs, offset := d.readSize4(offset)
		o := offset + int(binary.BigEndian.Uint32(bs))
		return d.data[offset:o], o, nil
	}

	return emptyBytes, 0, d.errorTemplate(code, k)
}
