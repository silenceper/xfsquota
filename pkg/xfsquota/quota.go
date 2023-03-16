package xfsquota

import (
	"strconv"

	"github.com/docker/go-units"
	"github.com/silenceper/xfsquota/pkg/projectquota"
)

// XfsQuota is the struct of xfs quota
type XfsQuota struct {
	*projectquota.ProjectQuota
}

// NewXfsQuota creates a new XfsQuota
func NewXfsQuota() *XfsQuota {
	return &XfsQuota{
		ProjectQuota: projectquota.NewProjectQuota(),
	}
}

// GetQuota returns the quota for the given path
func (q *XfsQuota) GetQuota(path string) (*projectquota.DiskQuotaSize, error) {
	return q.ProjectQuota.GetQuota(path)
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
	return q.ProjectQuota.SetQuota(path, &projectquota.DiskQuotaSize{
		Quota:  uint64(size),
		Inodes: inodes,
	})
}

// CleanQuota clears the quota for the given path
func (q *XfsQuota) CleanQuota(path string) error {
	return q.ProjectQuota.ClearQuota(path)
}
