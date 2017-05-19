/*
 * PKG - A collection of go utility packages
 * Copyright (C) 2017
 * O.S. Systems Sofware LTDA: contato@ossystems.com.br
 *
 * SPDX-License-Identifier:     MIT
 *
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to
 * the following conditions:
 *
 * The above copyright notice and this permission notice shall be
 * included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 * LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 * OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package keyvalueparser

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	expected := map[string]string{
		"key":         "value",
		"another_key": "value",
	}

	keyvalue, err := Parse(bytes.NewReader([]byte("key=value\nanother_key=value")))

	assert.NoError(t, err)
	assert.Equal(t, expected, keyvalue)
}

func TestParseInvalidKey(t *testing.T) {
	keyvalue, err := Parse(bytes.NewReader([]byte("key=value\nanother_key=value\ninvalid_key")))

	assert.EqualError(t, err, "'=' expected on line 3")
	assert.Nil(t, keyvalue)
}

func TestParseEmptyKey(t *testing.T) {
	expected := map[string]string{
		"key":       "value",
		"empty_key": "",
	}

	keyvalue, err := Parse(bytes.NewReader([]byte("key=value\nempty_key=")))

	assert.NoError(t, err)
	assert.Equal(t, expected, keyvalue)
}
