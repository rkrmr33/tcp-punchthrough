package helpers

import "syscall"

var sockOpts = [...]int{
	syscall.SO_REUSEADDR,
	syscall.TCP_KEEPALIVE,
}
