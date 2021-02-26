//
// Copyright (c) 2021 Cisco Systems, Inc and its affiliates
// All Rights reserved
//
package msxswagger

import (
	_ "github.com/CiscoDevNet/go-msx-swagger/noauth"
	"github.com/rakyll/statik/fs"
	"net/http"
)

func getnoauthfs() (http.FileSystem, error) {
	return fs.NewWithNamespace("noauth")
}
