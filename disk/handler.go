package disk

import (
	"context"
	"github.com/xgd16/x-object-storage/types"
)

func New(drive types.ObjectStorage) (types.ObjectStorage, error) {
	return drive.Init(context.Background())
}
