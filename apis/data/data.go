package data

import (
	"wgm/apis/common"
	"wgm/services/data"

	"github.com/gin-gonic/gin"
)

func ListPeerData(c *gin.Context) {
	data, err := data.PeerData()
	if err != nil {
		common.JSONHandler(c, 0, "List Peer data failed", []any{})
		return
	}
	common.JSONHandler(c, 1, "List Peer data succeed", data)
}
