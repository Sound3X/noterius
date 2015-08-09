package network

import (
	"testing"
)

func TestReadToStruct(t *testing.T) {
	var TestStruct struct {
		Id    int32
		Level uint32
		HP    uint32
	}

	netes := NewParser([]byte{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04})
	if err := netes.Read(&TestStruct.Id).Error; err != nil {
		t.Errorf("%v", "Error read from bytes to Id field")
	}

	if err := netes.Read(&TestStruct.Level).Error; err != nil {
		t.Errorf("%v", "Error read from bytes to Level field")
	}

	if err := netes.Read(&TestStruct.HP).Error; err != nil {
		t.Errorf("%v", "Error read from byte to HP field")
	}
}

func TestWriteFromStruct(t *testing.T) {
	var TestStruct = struct {
		Id    int32
		Level uint32
		HP    uint32
	}{
		2, 3, 4,
	}

	netes := NewParser([]byte{})
	if err := netes.Write(TestStruct.Id).Error; err != nil {
		t.Errorf("%v", "Error write to buffer from Id field")
	}

	if err := netes.Write(TestStruct.Level).Error; err != nil {
		t.Errorf("%v", "Error write to buffer from Level field")
	}

	if err := netes.Write(TestStruct.HP).Error; err != nil {
		t.Errorf("%v", "Error write to buffer from HP field")
	}
}

func BenchmarkReadToStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var TestStruct struct {
			Id    int32
			Level uint32
			HP    uint32
			Name  string
		}

		netes := NewParser([]byte{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x07, 0x4e, 0x79, 0x61, 0x72, 0x75, 0x6d, 0x00})
		netes.Read(&TestStruct.Id).Read(&TestStruct.Level).Read(&TestStruct.HP).Read(&TestStruct.Name)
	}
}

func BenchmarkWriteFromStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var TestStruct = struct {
			Id    int32
			Level int32
			HP    int32
			Name  string
		}{
			2, 3, 4, "Nyarum",
		}

		netes := NewParser([]byte{})
		netes.Write(TestStruct.Id).Write(TestStruct.Level).Write(TestStruct.HP).Write(TestStruct.Name)
	}
}
