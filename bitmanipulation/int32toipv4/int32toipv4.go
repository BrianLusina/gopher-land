package int32toipv4

import "fmt"

func Int32ToIpv4(ip int32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, (ip>>16)&0xff, (ip>>8)&0xff, ip&0xff)
}
