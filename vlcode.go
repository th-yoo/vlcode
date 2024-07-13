// Package vlcode implements an SMF styled Variable Length Code (VLC) Encoder and Decoder.
package vlcode

import "math/bits"

const Version = "1.0.0"

// Encode encodes a uint into a Variable Length Code (VLC) byte array.
// It follows the SMF styled VLC encoding scheme.
func Encode(v uint) []byte {
    // Determine the number of bits needed to represent v
    n_bits := bits.Len(v)
    if n_bits == 0 {
        return []byte{0x00}
    }

    // Calculate the number of bytes required for the encoded VLC
    n_bytes := (n_bits + 6) / 7

    // Create a byte slice to hold the encoded VLC
    bytes := make([]byte, n_bytes)
    bytes[n_bytes-1] = byte(v & 0x7F)

    // Encode v into bytes using SMF styled VLC encoding
    for i := n_bytes - 2; i >= 0; i-- {
        shift := (n_bytes - 1 - i) * 7
        bytes[i] = byte((v >> shift) & 0x7F) | 0x80
    }

    return bytes
}

// Decode decodes a Variable Length Code (VLC) byte array into a uint.
// It follows the SMF styled VLC decoding scheme.
// Returns the decoded value and the number of bytes read from b.
func Decode(b []byte) (uint, uint) {
    sz := uint(len(b))
    if sz == 0 {
        return 0, 0
    }

    n_left := sz - 1
    value := uint(b[0] & 0x7F)

    if b[0]&0x80 != 0 {
        // Decode bytes using SMF styled VLC decoding
        for i := uint(1); i < sz; i++ {
            if n_left == 0 {
                return 0, 0
            }
            value <<= 7
            value |= uint(b[i] & 0x7F)
            if b[i]&0x80 == 0 {
                break
            }
            n_left--
        }
    }

    return value, (sz - n_left)
}
