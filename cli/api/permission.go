package api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/ipfs-force-community/sophon-auth/core"
)

func PermissionedAPI(api LocalAPI) LocalAPI {
	var out LocalAPIStruct
	auth.PermissionedProxy(core.PermArr, []auth.Permission{}, api, &out.Internal)
	return &out
}
