// Code generated by "dbusutil-gen em -type Validator"; DO NOT EDIT.

package main

import (
	"pkg.deepin.io/lib/dbusutil"
)

func (v *Validator) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "ValidateHostname",
			Fn:      v.ValidateHostname,
			InArgs:  []string{"hostname"},
			OutArgs: []string{"result"},
		},
		{
			Name:    "ValidateHostnameTemp",
			Fn:      v.ValidateHostnameTemp,
			InArgs:  []string{"hostname"},
			OutArgs: []string{"result"},
		},
		{
			Name:    "ValidateUsername",
			Fn:      v.ValidateUsername,
			InArgs:  []string{"username"},
			OutArgs: []string{"code", "result"},
		},
	}
}