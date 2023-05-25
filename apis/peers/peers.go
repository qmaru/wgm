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
		common.JSONHandler(c, 0, "Peer List failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Peer List", data)
}

func AddPeer(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	var peerData models.Peers
	err := c.ShouldBindJSON(&peerData)
	if err != nil {
		common.JSONHandler(c, 0, "Peer data error", []any{})
		return
	}

	if peerData.UserID < 0 {
		common.JSONHandler(c, 0, "Peer id error", []any{})
		return
	}

	err = peers.PeerAdd(&peerData)
	if err != nil {
		common.JSONHandler(c, 0, "Add Peer failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Add Peer succeed", []any{})
}

func UpdatePeer(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	peerID := c.Param("peerID")
	peerIDN, err := common.IDtoInt(peerID)
	if err != nil {
		common.JSONHandler(c, 0, "Peer id error", []any{})
		return
	}

	var peerData models.Peers
	err = c.ShouldBindJSON(&peerData)
	if err != nil {
		common.JSONHandler(c, 0, "Peer data error", []any{})
		return
	}

	if peerData.UserID < 0 {
		common.JSONHandler(c, 0, "Peer id error", []any{})
		return
	}

	peerData.CommonModel.ID = peerIDN
	err = peers.PeerUpdate(&peerData)
	if err != nil {
		common.JSONHandler(c, 0, "Update Peer failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Update Peer succeed", []any{})
}

func DeletePeer(c *gin.Context) {
	c.Header("Content-Typ", "application/json")

	peerID := c.Param("peerID")
	peerIDN, err := common.IDtoInt(peerID)
	if err != nil {
		common.JSONHandler(c, 0, "Peer id error", []any{})
		return
	}

	err = peers.PeerDelete(peerIDN)
	if err != nil {
		common.JSONHandler(c, 0, "Delete Peer failed: "+err.Error(), []any{})
		return
	}
	common.JSONHandler(c, 1, "Delete Peer succeed", []any{})
}
