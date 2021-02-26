//
// Copyright (c) 2021 Cisco Systems, Inc and its affiliates
// All Rights reserved
//
package msxswagger

import (
	_ "github.com/CiscoDevNet/go-msx-swagger/auth"
	"github.com/rakyll/statik/fs"
	"net/http"
)

func getauthfs() (http.FileSystem, error) {
	f, err := fs.NewWithNamespace("auth")
	return f, err
}
