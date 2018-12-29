/*******************************************************************************
The MIT License (MIT)

Copyright (c) 2019 Hajime Nakagami

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*******************************************************************************/

package firebirdsql

import (
	"github.com/shopspring/decimal"
)

func dpdToInt(dpd uint16) int {
    // Convert DPD encodined value to int (0-999)
    // dpd: DPD encoded value. 10bit unsigned int

	return 0
}

func calcSignificand(prefix int64, dpdBits uint16, numBits int) int64 {
    // prefix: High bits integer value
    // dpdBits: dpd encoded bits
    // numBits: bit length of dpd_bits
    // https://en.wikipedia.org/wiki/Decimal128_floating-point_format#Densely_packed_decimal_significand_field

	return 0
}

func decimal128ToSignDigitsExponent(b []byte) (sign int, digits int, exponent int){
	// https://en.wikipedia.org/wiki/Decimal128_floating-point_format
	return 0, 0, 0
}

func decimalFixedToDecimal(b []byte, scale int) decimal.Decimal {
	return decimal.Zero
}

func decimal128ToDecimal(b []byte) decimal.Decimal {
    // https://en.wikipedia.org/wiki/Decimal64_floating-point_format
	return decimal.Zero
}
