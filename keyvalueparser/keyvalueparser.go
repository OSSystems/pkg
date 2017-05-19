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
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Parse parses source r into a map of key and values
func Parse(r io.Reader) (map[string]string, error) {
	b := bufio.NewReader(r)

	keyvalue := make(map[string]string)

	l := 0

	for {
		line, _, err := b.ReadLine()
		if err == io.EOF {
			break
		}

		l++

		parts := strings.SplitN(string(line), "=", 2)
		if len(parts) < 2 {
			return nil, fmt.Errorf("'=' expected on line %d", l)
		}

		keyvalue[parts[0]] = parts[1]
	}

	return keyvalue, nil
}
