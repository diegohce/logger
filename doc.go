// Copyright (c) 2019, Diego Cena.
// Use of this source code is governed by an AGPL-3.0
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/diegohce/logger

/*
Package logger is a simple logging mechanism I use (so far)
on my projects.

Example:

	import (
		"github.com/diegohce/logger"
	)

	func main() {
		log := logger.New("demo::")

		log.Info().Println("Hello, World!")
	}
*/
package logger
