package cpucaches

// FooV1 shows hows the arrangement of fields in a struct without care could lead to more memory being consumed and
// thus impacting performance. As see below the field b1 occupies 1 byte, i occupies 8 bytes and the b2 field will occupy
// 1 byte. In total, this leads to an allocation of 24 bytes. This could be improved with proper data alignment on the struct
type FooV1 struct {
	b1 byte
	i  int64
	b2 byte
}

// FooV2 shows how careful data alignment in descending order provides better compaction of the fields and better utilization of CPU memory
// In this case the i field will occupy 8 bytes, while b1 and b2 both occupy 1 byte each. However, the locality of b1 and b2 leads to a more
// efficient management of memory as the struct now occupies 16 bytes in total as compared to FooV1 which occupied 24 bytes. This improves CPU caching
// and also potentially improves performance.
type FooV2 struct {
	i  int64
	b1 byte
	b2 byte
}

// sumFooV1 sums all the i fields from each FooV1 struct in the slice provided
func sumFooV1(foos []FooV1) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}

// sumFooV2 sums all the i fields from each FooV2 struct in the slice provided
func sumFooV2(foos []FooV2) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}
