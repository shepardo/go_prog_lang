package tempconv0

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32)* 5 / 9) }

// KToC converts a Kelvin temperature to Celcius.
func KToC(k Kelvin) Celsius { return Celsius(k - Kelvin(float64(AbsoluteZeroC))) }

// CToK converst a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(KToC(k)*9/5 + 32) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(Celsius((f - 32)* 5 / 9)) }	