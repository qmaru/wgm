package data

import (
	"fmt"

	"wgm/common/services/peers"
	"wgm/common/services/routes"
)

func PeerData() (map[string]any, error) {
	data := make(map[string]any)

	routeData, err := routes.RouteList()
	if err != nil {
		return nil, fmt.Errorf("route data error")
	}

	peerData, err := peers.PeerList()
	if err != nil {
		return nil, fmt.Errorf("peer data error")
	}

	data["routes"] = routeData
	data["peers"] = peerData
	return data, nil
}
