package projectquota

// DiskQuotaSize group disk quota size
type DiskQuotaSize struct {
	Quota      uint64 `json:"quota"`
	Inodes     uint64 `json:"inodes"`
	QuotaUsed  uint64 `json:"-"`
	InodesUsed uint64 `json:"-"`
}
