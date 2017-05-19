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

package rcd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"syscall"

	"github.com/OSSystems/pkg/keyvalueparser"

	shellwords "github.com/mattn/go-shellwords"
)

// RunCommands executes each executable file in path and parses stdout into a key-value map
func RunCommands(basePath string) (map[string]string, error) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	keyValueMap := map[string]string{}

	for _, file := range files {
		if file.IsDir() || file.Mode()&syscall.S_IXUSR == 0 {
			continue
		}

		output, err := execute(path.Join(basePath, file.Name()))
		if err != nil {
			return nil, err
		}

		keyValue, err := keyvalueparser.Parse(bytes.NewReader(output))
		if err != nil {
			return nil, err
		}

		for k, v := range keyValue {
			keyValueMap[k] = v
		}
	}

	return keyValueMap, nil
}

func execute(cmdLine string) ([]byte, error) {
	p := shellwords.NewParser()
	list, err := p.Parse(cmdLine)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(list[0], list[1:]...)
	ret, err := cmd.CombinedOutput()

	if exitErr, ok := err.(*exec.ExitError); ok {
		if !exitErr.Success() {
			return ret, fmt.Errorf("Error executing command '%s': %s", cmdLine, string(ret))
		}
	}

	return ret, err
}
