package utils
import "net"
type User struct {
	Name string
	Conn net.Conn
}