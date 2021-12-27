package main

import (
	"AOC/pkg/utils"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/icza/bitio"
)

type packet struct {
	Version    byte
	TypeID     byte
	Data       []byte
	Subpackets []packet
}

func (p packet) String() string {
	e, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err.Error()
	}
	return (string(e))
}

func (p packet) versionSum() uint64 {
	result := uint64(p.Version)
	if len(p.Subpackets) > 0 {
		for _, sub := range p.Subpackets {
			result += sub.versionSum()
		}
	}

	return result
}

func (p packet) value() int {
	if p.TypeID == 4 {
		v := 0
		for i := 0; i < len(p.Data); i++ {
			v += int(p.Data[i]) << (4 * (len(p.Data) - 1 - i))
		}
		return v
	}

	if p.TypeID == 0 { /* sum */
		result := 0
		for _, sub := range p.Subpackets {
			result += sub.value()
		}
		return result
	}
	if p.TypeID == 1 { /* product */
		result := 1
		for _, sub := range p.Subpackets {
			result *= sub.value()
		}
		return result
	}
	if p.TypeID == 2 { /* minimum */
		result := math.MaxInt
		for _, sub := range p.Subpackets {
			v := sub.value()
			if v < result {
				result = v
			}
		}
		return result
	}
	if p.TypeID == 3 { /* maximum */
		result := 0
		for _, sub := range p.Subpackets {
			v := sub.value()
			if v > result {
				result = v
			}
		}
		return result
	}
	if p.TypeID == 5 { /* greater than */
		if p.Subpackets[0].value() > p.Subpackets[1].value() {
			return 1
		} else {
			return 0
		}
	}
	if p.TypeID == 6 { /* less than */
		if p.Subpackets[0].value() < p.Subpackets[1].value() {
			return 1
		} else {
			return 0
		}
	}
	if p.TypeID == 7 { /* equal */
		if p.Subpackets[0].value() == p.Subpackets[1].value() {
			return 1
		} else {
			return 0
		}
	}

	panic("not good")
}

func parse(r *bitio.Reader, readBits *int) *packet {
	version, _ := r.ReadBits(3)
	*readBits += 3

	typeID, _ := r.ReadBits(3)
	*readBits += 3

	if typeID == 4 { // literal packet
		data := []byte{}
		for {
			hasMore, _ := r.ReadBits(1)
			value, _ := r.ReadBits(4)
			*readBits += 5

			data = append(data, byte(value))
			if hasMore == 0 {
				break
			}
		}

		return &packet{Version: byte(version), TypeID: byte(typeID), Data: data}
	} else { // operator packet
		lengthTypeID, _ := r.ReadBits(1)
		*readBits++

		if lengthTypeID == 0 {
			totalLengthInBits, _ := r.ReadBits(15)
			*readBits += 15

			subpackets := []packet{}
			for totalLengthInBits > 0 {
				old := *readBits
				subpacket1 := parse(r, readBits)
				if subpacket1 != nil {
					totalLengthInBits -= uint64(*readBits - old)
					subpackets = append(subpackets, *subpacket1)
				}
			}

			return &packet{Version: byte(version), TypeID: byte(typeID), Data: []byte{}, Subpackets: subpackets}
		} else {
			totalNumberOfSubpackets, _ := r.ReadBits(11)
			*readBits += 11

			subpackets := []packet{}
			for i := 0; i < int(totalNumberOfSubpackets); i++ {
				subpacket := parse(r, readBits)
				subpackets = append(subpackets, *subpacket)
			}
			return &packet{Version: byte(version), TypeID: byte(typeID), Data: []byte{}, Subpackets: subpackets}
		}
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])
	b, _ := hex.DecodeString(lines[0])
	r := bitio.NewReader(bytes.NewBuffer(b)) // length type ID 0 that contains two sub-packets

	readBits := 0
	p := parse(r, &readBits)

	fmt.Println("Read", readBits, "out of", len(b)*8, "bits.")
	fmt.Println("Part 1 Answer:", p.versionSum())
	fmt.Println("Part 2 Answer:", p.value())
}
