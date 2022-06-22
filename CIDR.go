package netparse

import "net/netip"

func ParseFromCIDR(pre netip.Prefix) []netip.Addr {
	start := pre.Addr()
	var lists []netip.Addr
	for {
		if !pre.Contains(start.Next()) {
			break
		}
		lists = append(lists, start)
		start = start.Next()

	}
	return lists
}