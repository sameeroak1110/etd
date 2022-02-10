/* *****************************************************************************
Copyright (c) 2022, sameeroak1110 (sameeroak1110@gmail.com)
All rights reserved.
BSD 3-Clause License.

Package     : github.com/sameeroak1110/etd
Filename    : github.com/sameeroak1110/etd/etd.go
File-type   : golang source code file

Compiler/Runtime: go version go1.17 linux/amd64

Version History
Version     : 1.0
Author      : sameer oak (sameeroak1110@gmail.com)
Description :
- Execute-trap pattern. Catches panic and defines recover.
***************************************************************************** */
package etd

import (
	"runtime/debug"

	"github.com/sameeroak1110/logger"
)


func Shoot(exception Exception) {
	panic(exception)
}


func (etd ExecTrapDefault) Execute() {
	defer func() {
		if r := recover(); r != nil {
			logger.Log(pkgname, logger.WARNING, "Recovered from panic: %v", r)
			logger.Log(pkgname, logger.WARNING, "Dumping stack:\n==============")
			debug.PrintStack()
			etd.Trap(r)
			if etd.Trap != nil {
				etd.Trap(r)
			}
		}

		if etd.Default != nil {
			etd.Default()
		}
	}()

	etd.Exec()
}
