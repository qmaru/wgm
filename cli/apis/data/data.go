package data

import (
	"wgm/cli/apis/common"
	"wgm/common/services/data"

	"github.com/gin-gonic/gin"
)

func ListPeerData(c *gin.Context) {
	data, err := data.PeerData()
	if err != nil {
		common.JSONHandler(c, 0, "节点详情获取失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "节点详情", data)
}
