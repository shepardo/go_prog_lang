// Package lengthconv performs Feet and meter length conversions

package lengthconv

import "fmt"

type Feet float64
type Meters float64


func (f Feet) String() string { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string { return fmt.Sprintf("%gm", m) }

// MToF converts Meters to Feet
func MToF(m Meters) Feet { return Feet(m * 0.3048) }

// FToM converts Feet to Meters
func FToM(f Feet) Meters { return Meters(f /  3.281) }

