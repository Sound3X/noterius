package pills

import (
	"github.com/Nyarum/noterius/interfaces"
	"github.com/Nyarum/noterius/library/network"
	"github.com/Nyarum/noterius/pills/incoming"
	"github.com/Nyarum/noterius/pills/outcoming"
)

type Pill struct {
	incomingCrumbs  map[int]interfaces.PillDecoder
	outcomingCrumbs map[int]interfaces.PillEncoder

	opcode int
}

func NewPill() *Pill {
	return &Pill{
		incomingCrumbs: map[int]interfaces.PillDecoder{
			431: &incoming.CrumbAuth{},
		},
		outcomingCrumbs: map[int]interfaces.PillEncoder{
			940: &outcoming.CrumbDate{},
		},
	}
}

func (p *Pill) SetOpcode(opcode int) *Pill {
	p.opcode = opcode

	return p
}

func (p *Pill) GetIncomingCrumb() interfaces.PillDecoder {
	return p.incomingCrumbs[p.opcode]
}

func (p *Pill) GetOutcomingCrumb() interfaces.PillEncoder {
	return p.outcomingCrumbs[p.opcode]
}

func (p *Pill) Encrypt(pe interfaces.PillEncoder) []byte {
	netes := network.NewParser([]byte{})

	data := pe.Process().PostHandler(netes)
	netes.Reset()

	header := Header{Len: uint16(len(data) + 8), UniqueId: 128, Opcode: uint16(p.opcode)}

	netes.SetEndian(network.LittleEndian).WriteUint16(header.Len)
	netes.SetEndian(network.BigEndian).WriteUint32(header.UniqueId)
	netes.SetEndian(network.LittleEndian).WriteUint16(header.Opcode)
	netes.WriteBytes([]byte(data))

	return netes.Bytes()
}

func (p *Pill) Decrypt(buf []byte) int {
	var (
		header Header          = Header{}
		netes  *network.Parser = network.NewParser(buf)
	)

	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Len)
	netes.SetEndian(network.BigEndian).ReadUint32(&header.UniqueId)
	netes.SetEndian(network.LittleEndian).ReadUint16(&header.Opcode)

	return p.SetOpcode(int(header.Opcode)).GetIncomingCrumb().PreHandler(netes).Process()
}
