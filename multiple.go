// Copyright (c) 2016, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package errs

import (
	"fmt"
	"strings"
)

type Collector struct {
	Message string
	Errors  []error
}

func NewCollector(msg string) *Collector {
	return &Collector{
		Message: msg,
		Errors:  make([]error, 0),
	}
}

func (c *Collector) Add(err error) {
	if err != nil {
		c.Errors = append(c.Errors, err)
	}
}

func (c *Collector) Error() *Multiple {
	if len(c.Errors) > 0 {
		return &Multiple{c.Message, c.Errors}
	}
	return nil
}

type Multiple struct {
	Message string
	Errors  []error
}

func (e *Multiple) Error() string {
	xs := make([]string, len(e.Errors))
	for i, e := range e.Errors {
		xs[i] = e.Error()
	}
	return fmt.Sprintf("%s: %s", e.Message, strings.Join(xs, "; "))
}
