package metadata

import (
	"errors"
	"strconv"
)

type MetadataInt struct {
	Value      uint64
	IsUnsigned bool
}

func (m *MetadataInt) UnmarshalCBOR(bytes []byte) error {
	panic("implement me")
}

func (m *MetadataInt) AsMap() (MetadataMap, error) {
	return MetadataMap{}, errors.New("not a map")
}
func (m *MetadataInt) AsList() (MetadataList, error) {
	return MetadataList{}, errors.New("not a list")
}
func (m *MetadataInt) AsBytes() (MetadataBytes, error) {
	return MetadataBytes{}, errors.New("not bytes")
}
func (m *MetadataInt) AsInt() (MetadataInt, error) {
	return *m, nil
}
func (m *MetadataInt) AsText() (MetadataText, error) {
	return "", errors.New("not text")
}

func (m *MetadataInt) String() string {
	if m.IsUnsigned {
		return strconv.FormatUint(m.Value, 10)
	} else {
		return strconv.FormatInt(int64(m.Value), 10)
	}
}

func NewMetadataInt(value int64) *MetadataInt {
	return &MetadataInt{
		IsUnsigned: value > 0,
		Value:      uint64(value),
	}
}

func NewMetadataUInt(value uint64) *MetadataInt {
	return &MetadataInt{
		IsUnsigned: true,
		Value:      value,
	}
}