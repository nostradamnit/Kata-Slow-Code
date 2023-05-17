/*
Copyright (c) 2023 Murex

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package slow_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pascal_triangle_acceptance_test(t *testing.T) {
	// /!\ Should only pass at the end:
	// - Keep skipped
	// - Write and pass intermediate tests
	// - Un-skip at the end when should be passing
	t.Skip("test currently disabled")

	assert.Equal(t, []int{1}, pascalTriangle(0))
	assert.Equal(t, []int{1, 1}, pascalTriangle(1))
	assert.Equal(t, []int{1, 2, 1}, pascalTriangle(2))
	assert.Equal(t, []int{1, 3, 3, 1}, pascalTriangle(3))
	assert.Equal(t, []int{1, 4, 6, 4, 1}, pascalTriangle(4))
	assert.Equal(t, []int{1, 5, 10, 10, 5, 1}, pascalTriangle(5))
	assert.Equal(t, []int{1, 6, 15, 20, 15, 6, 1}, pascalTriangle(6))
	assert.Equal(t, []int{1, 7, 21, 35, 35, 21, 7, 1}, pascalTriangle(7))

	//...

	assert.Equal(t, []int{1, 15, 105, 455, 1365, 3003, 5005, 6435, 6435, 5005, 3003, 1365, 455, 105, 15, 1}, pascalTriangle(15))
}
