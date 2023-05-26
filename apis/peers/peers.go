package peers

import (
	"wgm/apis/common"
	"wgm/dbs/models"
	"wgm/services/peers"

	"github.com/gin-gonic/gin"
)

func ListPeers(c *gin.Context) {
	data, err := peers.PeerList()
	if err != nil {
		common.JSONHandler(c, 0, "节点列表获取失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "节点列表", data)
}

func AddPeer(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	var peerData models.Peers
	err := c.ShouldBindJSON(&peerData)
	if err != nil {
		common.JSONHandler(c, 0, "节点数据错误", []any{})
		return
	}

	if peerData.UserID < 0 {
		common.JSONHandler(c, 0, "节点ID错误", []any{})
		return
	}

	err = peers.PeerAdd(&peerData)
	if err != nil {
		common.JSONHandler(c, 0, "添加节点失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "添加节点成功", []any{})
}

func UpdatePeer(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	peerID := c.Param("peerID")
	peerIDN, err := common.IDtoInt(peerID)
	if err != nil {
		common.JSONHandler(c, 0, "节点ID错误", []any{})
		return
	}

	var peerData models.Peers
	err = c.ShouldBindJSON(&peerData)
	if err != nil {
		common.JSONHandler(c, 0, "节点数据错误", []any{})
		return
	}

	if peerData.UserID < 0 {
		common.JSONHandler(c, 0, "节点ID错误", []any{})
		return
	}

	peerData.CommonModel.ID = peerIDN
	err = peers.PeerUpdate(&peerData)
	if err != nil {
		common.JSONHandler(c, 0, "更新节点失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "更新节点成功", []any{})
}

func DeletePeer(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	peerID := c.Param("peerID")
	peerIDN, err := common.IDtoInt(peerID)
	if err != nil {
		common.JSONHandler(c, 0, "节点ID错误", []any{})
		return
	}

	err = peers.PeerDelete(peerIDN)
	if err != nil {
		common.JSONHandler(c, 0, "删除节点失败: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "删除节点成功", []any{})
}
