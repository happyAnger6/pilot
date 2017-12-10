package board

import "encoding/json"

const (
	boardtypempu = "mpu"
	boardtypelpu = "lpu"
)

type Board struct {
	ProjName string
	BoardType string
	ChassisNumber int64
	SlotNumber int64
	CpuNumber int64
	GInterfaceNum int64
	TGInterfaceNum int64
	Image string
	RunNode string
}

// NewFromJSON creates an Image configuration from json.
func NewFromJSON(src []byte) (*Board, error) {
	board := &Board{}

	if err := json.Unmarshal(src, board); err != nil {
		return nil, err
	}

	return board, nil
}
