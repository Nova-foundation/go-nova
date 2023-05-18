package gossip

import (
	"github.com/Nova-foundation/go-nova/nova"
)

func (s *Store) AddUpgradeHeight(h nova.UpgradeHeight) {
	orig := s.GetUpgradeHeights()
	// allocate new memory to avoid race condition in cache
	cp := make([]nova.UpgradeHeight, 0, len(orig)+1)
	cp = append(append(cp, orig...), h)

	s.rlp.Set(s.table.UpgradeHeights, []byte{}, cp)
	s.cache.UpgradeHeights.Store(cp)
}

func (s *Store) GetUpgradeHeights() []nova.UpgradeHeight {
	if v := s.cache.UpgradeHeights.Load(); v != nil {
		return v.([]nova.UpgradeHeight)
	}
	hh, ok := s.rlp.Get(s.table.UpgradeHeights, []byte{}, &[]nova.UpgradeHeight{}).(*[]nova.UpgradeHeight)
	if !ok {
		return []nova.UpgradeHeight{}
	}
	s.cache.UpgradeHeights.Store(*hh)
	return *hh
}
