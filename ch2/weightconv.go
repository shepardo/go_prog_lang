// Package weight performs Pounds and Kilograms conversions

package weightconv

import "fmt"

type Kilograms float64
type Pounds float64


func (p Pounds) String() string { return fmt.Sprintf("%glb", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%gkg", k) }

// KToP converts Kilograms to Pounds
func KToP(k Kilograms) Pounds { return Pounds(k * 2.20462) }

// PToK converts Pounds to Kilograms
func PToK(p Pounds) Kilograms { return Kilograms(p * 0.9999988107) }
