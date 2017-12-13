package board


const (
	GigabitEthernet="GigabitEthernet"
	TenGigabitEthernet="Ten-GigabitEthernet"
	VlanInterface="Vlan-interface"
)

type BoardInterface struct {
	BoardName string
	IfType string
	IfName string
	Endpoint *BoardInterface
}
