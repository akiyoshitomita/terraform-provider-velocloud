package vcoclient

type FirewallStatefulFirewallSettings struct {
	EstablishedTcpFlowTimeout    int `json:"establishedTcpFlowTimeout,omitempty"`
	NonEstablishedTcpFlowTimeout int `json:"nonEstablishedTcpFlowTimeout,omitempty"`
	UdpFlowTimeout               int `json:"udpFlowTimeout,omitempty"`
	OtherFlowTimeout             int `json:"otherFlowTimeout,omitempty"`
}
