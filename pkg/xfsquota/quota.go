package xfsquota

import (
	"github.com/containers/storage/drivers/quota"
	"github.com/containers/storage/pkg/directory"
	"github.com/docker/go-units"
	"strconv"
)

type XfsQuota struct {
	*quota.Control
}

// NewXfsQuota creates a new XfsQuota
func NewXfsQuota() *XfsQuota {
	return &XfsQuota{}
}

func (q *XfsQuota) Init(basePath string) error {
	control, err := quota.NewControl(basePath)
	if err != nil {
		panic(err)
	}
	q.Control = control
	return nil
}

type QuotaStatus struct {
	quota.Quota
	directory.DiskUsage
}

// GetQuota returns the quota for the given path
func (q *XfsQuota) GetQuota(path string) (*QuotaStatus, error) {
	quotaRes := quota.Quota{}
	err := q.Control.GetQuota(path, &quotaRes)
	if err != nil {
		return nil, err
	}
	diskUsageRes := directory.DiskUsage{}
	err = q.Control.GetDiskUsage(path, &diskUsageRes)
	if err != nil {
		return nil, err
	}
	return &QuotaStatus{
		Quota:     quotaRes,
		DiskUsage: diskUsageRes,
	}, nil
}

// SetQuota sets the quota for the given path
func (q *XfsQuota) SetQuota(path string, sizeVal, inodeVal string) error {
	size, err := units.RAMInBytes(sizeVal)
	if err != nil {
		return err
	}
	inodes, err := strconv.ParseUint(inodeVal, 10, 64)
	if err != nil {
		return err
	}
	return q.Control.SetQuota(path, quota.Quota{
		Size:   uint64(size),
		Inodes: inodes,
	})
}
