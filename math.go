package gltf

import "math"

// NormalizeByte normalize a float64 into a int8
func NormalizeByte(v float64) int8 {
	return int8(math.Round(float64(v) * 127))
}

// NormalizeByte denormalize a int8 into a float64
func DenormalizeByte(v int8) float64 {
	return float64(math.Max(float64(v)/127, -1))
}

// NormalizeUbyte normalize a float64 into a uint8
func NormalizeUbyte(v float64) uint8 {
	return uint8(math.Round(float64(v) * 255))
}

// DenormalizeUbyte denormalize a uint8 into a float64
func DenormalizeUbyte(v uint8) float64 {
	return float64(v) / 255
}

// NormalizeShort normalize a float64 into a int16
func NormalizeShort(v float64) int16 {
	return int16(math.Round(float64(v) * 32767))
}

// DenormalizeShort denormalize a int16 into a float64
func DenormalizeShort(v int16) float64 {
	return float64(math.Max(float64(v)/32767, -1))
}

// NormalizeUshort normalize a float64 into a uint16
func NormalizeUshort(v float64) uint16 {
	return uint16(math.Round(float64(v) * 65535))
}

// DenormalizeuShort denormalize a uint16 into a float64
func DenormalizeUshort(v uint16) float64 {
	return float64(v) / 65535
}

// NormalizeRGB transform a RGB float64 color (from 0 to 1)
// to its uint8 represtation (from 0 to 255).
func NormalizeRGB(v [3]float64) [3]uint8 {
	return [3]uint8{
		uint8(deliniarize(v[0]) * 255),
		uint8(deliniarize(v[1]) * 255),
		uint8(deliniarize(v[2]) * 255),
	}
}

// DenormalizeRGB transform a RGB uint8 color (from 0 to 255)
// to its float represtation (from 0 to 1).
func DenormalizeRGB(v [3]uint8) [3]float64 {
	return [3]float64{
		linearize(float64(v[0]) / 255),
		linearize(float64(v[1]) / 255),
		linearize(float64(v[2]) / 255),
	}
}

// NormalizeRGB transform a RGBA float64 color (from 0 to 1)
// to its uint8 represtation (from 0 to 255).
func NormalizeRGBA(v [4]float64) [4]uint8 {
	return [4]uint8{
		uint8(deliniarize(v[0]) * 255),
		uint8(deliniarize(v[1]) * 255),
		uint8(deliniarize(v[2]) * 255),
		uint8(v[3] * 255),
	}
}

// DenormalizeRGBA transform a RGBA uint8 color (from 0 to 255)
// to its float represtation (from 0 to 1).
func DenormalizeRGBA(v [4]uint8) [4]float64 {
	return [4]float64{
		linearize(float64(v[0]) / 255),
		linearize(float64(v[1]) / 255),
		linearize(float64(v[2]) / 255),
		float64(v[3]) / 255,
	}
}

// NormalizeRGB64 transform a RGB float64 color (from 0 to 1)
// to its uint16 represtation (from 0 to 65535).
func NormalizeRGB64(v [3]float64) [3]uint16 {
	return [3]uint16{
		uint16(deliniarize(v[0]) * 65535),
		uint16(deliniarize(v[1]) * 65535),
		uint16(deliniarize(v[2]) * 65535),
	}
}

// DenormalizeRGB64 transform a RGB uint16 color (from 0 to 65535)
// to its float represtation (from 0 to 1).
func DenormalizeRGB64(v [3]uint16) [3]float64 {
	return [3]float64{
		linearize(float64(v[0]) / 65535),
		linearize(float64(v[1]) / 65535),
		linearize(float64(v[2]) / 65535),
	}
}

// NormalizeRGBA64 transform a RGBA float64 color (from 0 to 1)
// to its uint16 represtation (from 0 to 65535).
func NormalizeRGBA64(v [4]float64) [4]uint16 {
	return [4]uint16{
		uint16(deliniarize(v[0]) * 65535),
		uint16(deliniarize(v[1]) * 65535),
		uint16(deliniarize(v[2]) * 65535),
		uint16(v[3] * 65535),
	}
}

// DenormalizeRGBA64 transform a RGBA uint16 color (from 0 to 65535)
// to its float represtation (from 0 to 1).
func DenormalizeRGBA64(v [4]uint16) [4]float64 {
	return [4]float64{
		linearize(float64(v[0]) / 65535),
		linearize(float64(v[1]) / 65535),
		linearize(float64(v[2]) / 65535),
		float64(v[3]) / 65535,
	}
}

func linearize(v float64) float64 {
	if v <= 0.04045 {
		return float64(v / 12.92)
	}
	return float64(math.Pow(float64(v+0.055)/1.055, 2.4))
}

func deliniarize(v float64) float64 {
	if v < 0.0031308 {
		return v * 12.92
	}
	return float64(1.055*math.Pow(float64(v), 1.0/2.4) - 0.055)
}
