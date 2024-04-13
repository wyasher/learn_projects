package socks5

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
)

// 地址类型
const (
	ipv4Address uint8 = 0x01 // ipv4类型
	fqdnAddress uint8 = 0x03 // FQDN代表Fully Qualified Domain Name 完全域名类型
	ipv6Address uint8 = 0x04 // ipv6类型
)

var (
	unrecognizedAddrType = errors.New("unrecognized address type")
)

type AddressSpec interface {
	fmt.Stringer
	Decode(r io.Reader) error
	Type() uint8
	Encode() []byte
	// Len 地址部分协议长度
	Len() int
	Address() string
}
type BaseAddrSpec struct {
	Port int    // 端口号
	IP   net.IP // ip地址 v4
}
type IPV4AddrSpec struct {
	BaseAddrSpec
}

func (i *IPV4AddrSpec) Type() uint8 {
	return ipv4Address
}

func (i *IPV4AddrSpec) Decode(r io.Reader) error {
	addr := make([]byte, 4)
	if _, err := io.ReadAtLeast(r, addr, len(addr)); err != nil {
		return err
	}
	i.IP = addr
	return decodePort(r, &i.BaseAddrSpec)
}

func (i *IPV4AddrSpec) Encode() []byte {
	bytes := make([]byte, i.Len())
	bytes[0] = ipv4Address
	copy(bytes[1:], i.IP)
	encodePort(bytes, i.Port, 5)
	return bytes
}
func (i *IPV4AddrSpec) Len() int {
	// 1 addressType 4 ipv6 2 port
	return 7
}
func (i *IPV4AddrSpec) String() string {
	return fmt.Sprintf("%s:%d", i.IP.String(), i.Port)
}

func (i *IPV4AddrSpec) Address() string {
	return net.JoinHostPort(i.IP.String(), strconv.Itoa(i.Port))
}

type IPV6AddrSpec struct {
	BaseAddrSpec
	IP net.IP // ip地址 v6
}

func (i *IPV6AddrSpec) Type() uint8 {
	return ipv6Address
}

func (i *IPV6AddrSpec) String() string {
	return fmt.Sprintf("%s:%d", i.IP.String(), i.Port)
}

func (i *IPV6AddrSpec) Address() string {
	return net.JoinHostPort(i.IP.String(), strconv.Itoa(i.Port))
}

func (i *IPV6AddrSpec) Decode(r io.Reader) error {
	addr := make([]byte, 16)
	if _, err := io.ReadAtLeast(r, addr, len(addr)); err != nil {
		return err
	}
	i.IP = addr
	return decodePort(r, &i.BaseAddrSpec)
}

func (i *IPV6AddrSpec) Encode() []byte {
	bytes := make([]byte, i.Len())
	bytes[0] = ipv6Address
	copy(bytes[1:], i.IP)
	encodePort(bytes, i.Port, 17)
	return bytes
}
func (i *IPV6AddrSpec) Len() int {
	// 1 addressType 16 ipv6 2 port
	return 19
}

type FQDNAddrSpec struct {
	BaseAddrSpec
	FQDN string
}

func (f *FQDNAddrSpec) Type() uint8 {
	return fqdnAddress
}

func (f *FQDNAddrSpec) String() string {
	return fmt.Sprintf("%s (%s):%d", f.FQDN, f.IP, f.Port)
}

func (f *FQDNAddrSpec) Address() string {
	return net.JoinHostPort(f.FQDN, strconv.Itoa(f.Port))
}

// Decode 域名类型时，第一个字节为域名长度 然后才是域名数据
func (f *FQDNAddrSpec) Decode(r io.Reader) error {
	addrType := []byte{0}
	if _, err := r.Read(addrType); err != nil {
		return err
	}
	addrLen := int(addrType[0])
	fqdn := make([]byte, addrLen)
	if _, err := io.ReadAtLeast(r, fqdn, addrLen); err != nil {
		return err
	}
	f.FQDN = string(fqdn)
	return decodePort(r, &f.BaseAddrSpec)
}

func (f *FQDNAddrSpec) Encode() []byte {
	bytes := make([]byte, f.Len())
	bytes[0] = fqdnAddress
	copy(bytes[1:], f.FQDN)
	encodePort(bytes, f.Port, f.Len()-3)
	return bytes
}
func (f *FQDNAddrSpec) Len() int {
	// 1 addressType 1 urlLen  n url 2 port
	return 1 + 1 + len(f.FQDN) + 2
}

// encode 端口信息
func encodePort(bytes []byte, port, index int) {
	// port 高位
	bytes[index] = byte(port >> 8)
	index++
	// port 低位
	bytes[index] = byte(port & 0xFF)
}

// ConvertTcpAddr 将TcpAddr转换为AddressSpec
func ConvertTcpAddr(addr *net.TCPAddr) AddressSpec {
	var addrSpec AddressSpec
	var baseAddressSpec = BaseAddrSpec{
		Port: addr.Port,
		IP:   addr.IP,
	}
	if len(addr.IP) == net.IPv4len {
		addrSpec = &IPV4AddrSpec{BaseAddrSpec: baseAddressSpec}
	} else {
		addrSpec = &IPV6AddrSpec{BaseAddrSpec: baseAddressSpec}
	}
	return addrSpec
}

// 读取port
func decodePort(r io.Reader, address *BaseAddrSpec) error {
	// 读取port
	port := []byte{0, 0}
	if _, err := io.ReadAtLeast(r, port, len(port)); err != nil {
		return err
	}
	address.Port = (int(port[0]) << 8) | int(port[1])
	return nil
}

// NewAddressSpec 获取地址信息
func NewAddressSpec(r io.Reader) (AddressSpec, error) {
	var address AddressSpec
	var err error
	// 读取地址类型
	addrType := []byte{0}
	if _, err := r.Read(addrType); err != nil {
		return nil, err
	}
	// 读取地址
	switch addrType[0] {
	case ipv4Address:
		address = &IPV4AddrSpec{}
	case ipv6Address:
		address = &IPV6AddrSpec{}
	case fqdnAddress:
		address = &FQDNAddrSpec{}
	default:
		// 无法识别 不支持的地址类型
		return nil, unrecognizedAddrType
	}
	err = address.Decode(r)
	return address, err
}
