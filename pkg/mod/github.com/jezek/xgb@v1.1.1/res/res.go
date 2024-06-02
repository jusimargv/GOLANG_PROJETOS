// Package res is the X client API for the X-Resource extension.
package res

// This file is automatically generated from res.xml. Edit at your peril!

import (
	"github.com/jezek/xgb"

	"github.com/jezek/xgb/xproto"
)

// Init must be called before using the X-Resource extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 10, "X-Resource").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named X-Resource could be found on on the server.")
	}

	c.ExtLock.Lock()
	c.Extensions["X-Resource"] = reply.MajorOpcode
	c.ExtLock.Unlock()
	for evNum, fun := range xgb.NewExtEventFuncs["X-Resource"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["X-Resource"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	return nil
}

func init() {
	xgb.NewExtEventFuncs["X-Resource"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["X-Resource"] = make(map[int]xgb.NewErrorFun)
}

type Client struct {
	ResourceBase uint32
	ResourceMask uint32
}

// ClientRead reads a byte slice into a Client value.
func ClientRead(buf []byte, v *Client) int {
	b := 0

	v.ResourceBase = xgb.Get32(buf[b:])
	b += 4

	v.ResourceMask = xgb.Get32(buf[b:])
	b += 4

	return b
}

// ClientReadList reads a byte slice into a list of Client values.
func ClientReadList(buf []byte, dest []Client) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Client{}
		b += ClientRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a Client value to a byte slice.
func (v Client) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], v.ResourceBase)
	b += 4

	xgb.Put32(buf[b:], v.ResourceMask)
	b += 4

	return buf[:b]
}

// ClientListBytes writes a list of Client values to a byte slice.
func ClientListBytes(buf []byte, list []Client) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return xgb.Pad(b)
}

const (
	ClientIdMaskClientXID      = 1
	ClientIdMaskLocalClientPID = 2
)

type ClientIdSpec struct {
	Client uint32
	Mask   uint32
}

// ClientIdSpecRead reads a byte slice into a ClientIdSpec value.
func ClientIdSpecRead(buf []byte, v *ClientIdSpec) int {
	b := 0

	v.Client = xgb.Get32(buf[b:])
	b += 4

	v.Mask = xgb.Get32(buf[b:])
	b += 4

	return b
}

// ClientIdSpecReadList reads a byte slice into a list of ClientIdSpec values.
func ClientIdSpecReadList(buf []byte, dest []ClientIdSpec) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ClientIdSpec{}
		b += ClientIdSpecRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a ClientIdSpec value to a byte slice.
func (v ClientIdSpec) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], v.Client)
	b += 4

	xgb.Put32(buf[b:], v.Mask)
	b += 4

	return buf[:b]
}

// ClientIdSpecListBytes writes a list of ClientIdSpec values to a byte slice.
func ClientIdSpecListBytes(buf []byte, list []ClientIdSpec) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return xgb.Pad(b)
}

type ClientIdValue struct {
	Spec   ClientIdSpec
	Length uint32
	Value  []uint32 // size: xgb.Pad(((int(Length) / 4) * 4))
}

// ClientIdValueRead reads a byte slice into a ClientIdValue value.
func ClientIdValueRead(buf []byte, v *ClientIdValue) int {
	b := 0

	v.Spec = ClientIdSpec{}
	b += ClientIdSpecRead(buf[b:], &v.Spec)

	v.Length = xgb.Get32(buf[b:])
	b += 4

	v.Value = make([]uint32, (int(v.Length) / 4))
	for i := 0; i < int((int(v.Length) / 4)); i++ {
		v.Value[i] = xgb.Get32(buf[b:])
		b += 4
	}

	return b
}

// ClientIdValueReadList reads a byte slice into a list of ClientIdValue values.
func ClientIdValueReadList(buf []byte, dest []ClientIdValue) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ClientIdValue{}
		b += ClientIdValueRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a ClientIdValue value to a byte slice.
func (v ClientIdValue) Bytes() []byte {
	buf := make([]byte, (12 + xgb.Pad(((int(v.Length) / 4) * 4))))
	b := 0

	{
		structBytes := v.Spec.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	xgb.Put32(buf[b:], v.Length)
	b += 4

	for i := 0; i < int((int(v.Length) / 4)); i++ {
		xgb.Put32(buf[b:], v.Value[i])
		b += 4
	}

	return buf[:b]
}

// ClientIdValueListBytes writes a list of ClientIdValue values to a byte slice.
func ClientIdValueListBytes(buf []byte, list []ClientIdValue) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return xgb.Pad(b)
}

// ClientIdValueListSize computes the size (bytes) of a list of ClientIdValue values.
func ClientIdValueListSize(list []ClientIdValue) int {
	size := 0
	for _, item := range list {
		size += (12 + xgb.Pad(((int(item.Length) / 4) * 4)))
	}
	return size
}

type ResourceIdSpec struct {
	Resource uint32
	Type     uint32
}

// ResourceIdSpecRead reads a byte slice into a ResourceIdSpec value.
func ResourceIdSpecRead(buf []byte, v *ResourceIdSpec) int {
	b := 0

	v.Resource = xgb.Get32(buf[b:])
	b += 4

	v.Type = xgb.Get32(buf[b:])
	b += 4

	return b
}

// ResourceIdSpecReadList reads a byte slice into a list of ResourceIdSpec values.
func ResourceIdSpecReadList(buf []byte, dest []ResourceIdSpec) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ResourceIdSpec{}
		b += ResourceIdSpecRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a ResourceIdSpec value to a byte slice.
func (v ResourceIdSpec) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], v.Resource)
	b += 4

	xgb.Put32(buf[b:], v.Type)
	b += 4

	return buf[:b]
}

// ResourceIdSpecListBytes writes a list of ResourceIdSpec values to a byte slice.
func ResourceIdSpecListBytes(buf []byte, list []ResourceIdSpec) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return xgb.Pad(b)
}

type ResourceSizeSpec struct {
	Spec     ResourceIdSpec
	Bytes_   uint32
	RefCount uint32
	UseCount uint32
}

// ResourceSizeSpecRead reads a byte slice into a ResourceSizeSpec value.
func ResourceSizeSpecRead(buf []byte, v *ResourceSizeSpec) int {
	b := 0

	v.Spec = ResourceIdSpec{}
	b += ResourceIdSpecRead(buf[b:], &v.Spec)

	v.Bytes_ = xgb.Get32(buf[b:])
	b += 4

	v.RefCount = xgb.Get32(buf[b:])
	b += 4

	v.UseCount = xgb.Get32(buf[b:])
	b += 4

	return b
}

// ResourceSizeSpecReadList reads a byte slice into a list of ResourceSizeSpec values.
func ResourceSizeSpecReadList(buf []byte, dest []ResourceSizeSpec) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ResourceSizeSpec{}
		b += ResourceSizeSpecRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a ResourceSizeSpec value to a byte slice.
func (v ResourceSizeSpec) Bytes() []byte {
	buf := make([]byte, 20)
	b := 0

	{
		structBytes := v.Spec.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	xgb.Put32(buf[b:], v.Bytes_)
	b += 4

	xgb.Put32(buf[b:], v.RefCount)
	b += 4

	xgb.Put32(buf[b:], v.UseCount)
	b += 4

	return buf[:b]
}

// ResourceSizeSpecListBytes writes a list of ResourceSizeSpec values to a byte slice.
func ResourceSizeSpecListBytes(buf []byte, list []ResourceSizeSpec) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return xgb.Pad(b)
}

type ResourceSizeValue struct {
	Size               ResourceSizeSpec
	NumCrossReferences uint32
	CrossReferences    []ResourceSizeSpec // size: xgb.Pad((int(NumCrossReferences) * 20))
}

// ResourceSizeValueRead reads a byte slice into a ResourceSizeValue value.
func ResourceSizeValueRead(buf []byte, v *ResourceSizeValue) int {
	b := 0

	v.Size = ResourceSizeSpec{}
	b += ResourceSizeSpecRead(buf[b:], &v.Size)

	v.NumCrossReferences = xgb.Get32(buf[b:])
	b += 4

	v.CrossReferences = make([]ResourceSizeSpec, v.NumCrossReferences)
	b += ResourceSizeSpecReadList(buf[b:], v.CrossReferences)

	return b
}

// ResourceSizeValueReadList reads a byte slice into a list of ResourceSizeValue values.
func ResourceSizeValueReadList(buf []byte, dest []ResourceSizeValue) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ResourceSizeValue{}
		b += ResourceSizeValueRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a ResourceSizeValue value to a byte slice.
func (v ResourceSizeValue) Bytes() []byte {
	buf := make([]byte, (24 + xgb.Pad((int(v.NumCrossReferences) * 20))))
	b := 0

	{
		structBytes := v.Size.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	xgb.Put32(buf[b:], v.NumCrossReferences)
	b += 4

	b += ResourceSizeSpecListBytes(buf[b:], v.CrossReferences)

	return buf[:b]
}

// ResourceSizeValueListBytes writes a list of ResourceSizeValue values to a byte slice.
func ResourceSizeValueListBytes(buf []byte, list []ResourceSizeValue) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return xgb.Pad(b)
}

// ResourceSizeValueListSize computes the size (bytes) of a list of ResourceSizeValue values.
func ResourceSizeValueListSize(list []ResourceSizeValue) int {
	size := 0
	for _, item := range list {
		size += (24 + xgb.Pad((int(item.NumCrossReferences) * 20)))
	}
	return size
}

type Type struct {
	ResourceType xproto.Atom
	Count        uint32
}

// TypeRead reads a byte slice into a Type value.
func TypeRead(buf []byte, v *Type) int {
	b := 0

	v.ResourceType = xproto.Atom(xgb.Get32(buf[b:]))
	b += 4

	v.Count = xgb.Get32(buf[b:])
	b += 4

	return b
}

// TypeReadList reads a byte slice into a list of Type values.
func TypeReadList(buf []byte, dest []Type) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Type{}
		b += TypeRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a Type value to a byte slice.
func (v Type) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], uint32(v.ResourceType))
	b += 4

	xgb.Put32(buf[b:], v.Count)
	b += 4

	return buf[:b]
}

// TypeListBytes writes a list of Type values to a byte slice.
func TypeListBytes(buf []byte, list []Type) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return xgb.Pad(b)
}

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Card32'

// QueryClientIdsCookie is a cookie used only for QueryClientIds requests.
type QueryClientIdsCookie struct {
	*xgb.Cookie
}

// QueryClientIds sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryClientIdsCookie.Reply()
func QueryClientIds(c *xgb.Conn, NumSpecs uint32, Specs []ClientIdSpec) QueryClientIdsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryClientIds' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientIdsRequest(c, NumSpecs, Specs), cookie)
	return QueryClientIdsCookie{cookie}
}

// QueryClientIdsUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryClientIdsUnchecked(c *xgb.Conn, NumSpecs uint32, Specs []ClientIdSpec) QueryClientIdsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryClientIds' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientIdsRequest(c, NumSpecs, Specs), cookie)
	return QueryClientIdsCookie{cookie}
}

// QueryClientIdsReply represents the data returned from a QueryClientIds request.
type QueryClientIdsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumIds uint32
	// padding: 20 bytes
	Ids []ClientIdValue // size: ClientIdValueListSize(Ids)
}

// Reply blocks and returns the reply data for a QueryClientIds request.
func (cook QueryClientIdsCookie) Reply() (*QueryClientIdsReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientIdsReply(buf), nil
}

// queryClientIdsReply reads a byte slice into a QueryClientIdsReply value.
func queryClientIdsReply(buf []byte) *QueryClientIdsReply {
	v := new(QueryClientIdsReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.NumIds = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Ids = make([]ClientIdValue, v.NumIds)
	b += ClientIdValueReadList(buf[b:], v.Ids)

	return v
}

// Write request to wire for QueryClientIds
// queryClientIdsRequest writes a QueryClientIds request to a byte slice.
func queryClientIdsRequest(c *xgb.Conn, NumSpecs uint32, Specs []ClientIdSpec) []byte {
	size := xgb.Pad((8 + xgb.Pad((int(NumSpecs) * 8))))
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["X-Resource"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], NumSpecs)
	b += 4

	b += ClientIdSpecListBytes(buf[b:], Specs)

	return buf
}

// QueryClientPixmapBytesCookie is a cookie used only for QueryClientPixmapBytes requests.
type QueryClientPixmapBytesCookie struct {
	*xgb.Cookie
}

// QueryClientPixmapBytes sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryClientPixmapBytesCookie.Reply()
func QueryClientPixmapBytes(c *xgb.Conn, Xid uint32) QueryClientPixmapBytesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryClientPixmapBytes' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientPixmapBytesRequest(c, Xid), cookie)
	return QueryClientPixmapBytesCookie{cookie}
}

// QueryClientPixmapBytesUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryClientPixmapBytesUnchecked(c *xgb.Conn, Xid uint32) QueryClientPixmapBytesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryClientPixmapBytes' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientPixmapBytesRequest(c, Xid), cookie)
	return QueryClientPixmapBytesCookie{cookie}
}

// QueryClientPixmapBytesReply represents the data returned from a QueryClientPixmapBytes request.
type QueryClientPixmapBytesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Bytes_        uint32
	BytesOverflow uint32
}

// Reply blocks and returns the reply data for a QueryClientPixmapBytes request.
func (cook QueryClientPixmapBytesCookie) Reply() (*QueryClientPixmapBytesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientPixmapBytesReply(buf), nil
}

// queryClientPixmapBytesReply reads a byte slice into a QueryClientPixmapBytesReply value.
func queryClientPixmapBytesReply(buf []byte) *QueryClientPixmapBytesReply {
	v := new(QueryClientPixmapBytesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Bytes_ = xgb.Get32(buf[b:])
	b += 4

	v.BytesOverflow = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for QueryClientPixmapBytes
// queryClientPixmapBytesRequest writes a QueryClientPixmapBytes request to a byte slice.
func queryClientPixmapBytesRequest(c *xgb.Conn, Xid uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["X-Resource"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Xid)
	b += 4

	return buf
}

// QueryClientResourcesCookie is a cookie used only for QueryClientResources requests.
type QueryClientResourcesCookie struct {
	*xgb.Cookie
}

// QueryClientResources sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryClientResourcesCookie.Reply()
func QueryClientResources(c *xgb.Conn, Xid uint32) QueryClientResourcesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryClientResources' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientResourcesRequest(c, Xid), cookie)
	return QueryClientResourcesCookie{cookie}
}

// QueryClientResourcesUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryClientResourcesUnchecked(c *xgb.Conn, Xid uint32) QueryClientResourcesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryClientResources' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientResourcesRequest(c, Xid), cookie)
	return QueryClientResourcesCookie{cookie}
}

// QueryClientResourcesReply represents the data returned from a QueryClientResources request.
type QueryClientResourcesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumTypes uint32
	// padding: 20 bytes
	Types []Type // size: xgb.Pad((int(NumTypes) * 8))
}

// Reply blocks and returns the reply data for a QueryClientResources request.
func (cook QueryClientResourcesCookie) Reply() (*QueryClientResourcesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientResourcesReply(buf), nil
}

// queryClientResourcesReply reads a byte slice into a QueryClientResourcesReply value.
func queryClientResourcesReply(buf []byte) *QueryClientResourcesReply {
	v := new(QueryClientResourcesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.NumTypes = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Types = make([]Type, v.NumTypes)
	b += TypeReadList(buf[b:], v.Types)

	return v
}

// Write request to wire for QueryClientResources
// queryClientResourcesRequest writes a QueryClientResources request to a byte slice.
func queryClientResourcesRequest(c *xgb.Conn, Xid uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["X-Resource"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Xid)
	b += 4

	return buf
}

// QueryClientsCookie is a cookie used only for QueryClients requests.
type QueryClientsCookie struct {
	*xgb.Cookie
}

// QueryClients sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryClientsCookie.Reply()
func QueryClients(c *xgb.Conn) QueryClientsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryClients' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryClientsRequest(c), cookie)
	return QueryClientsCookie{cookie}
}

// QueryClientsUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryClientsUnchecked(c *xgb.Conn) QueryClientsCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryClients' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryClientsRequest(c), cookie)
	return QueryClientsCookie{cookie}
}

// QueryClientsReply represents the data returned from a QueryClients request.
type QueryClientsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumClients uint32
	// padding: 20 bytes
	Clients []Client // size: xgb.Pad((int(NumClients) * 8))
}

// Reply blocks and returns the reply data for a QueryClients request.
func (cook QueryClientsCookie) Reply() (*QueryClientsReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryClientsReply(buf), nil
}

// queryClientsReply reads a byte slice into a QueryClientsReply value.
func queryClientsReply(buf []byte) *QueryClientsReply {
	v := new(QueryClientsReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.NumClients = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Clients = make([]Client, v.NumClients)
	b += ClientReadList(buf[b:], v.Clients)

	return v
}

// Write request to wire for QueryClients
// queryClientsRequest writes a QueryClients request to a byte slice.
func queryClientsRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["X-Resource"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// QueryResourceBytesCookie is a cookie used only for QueryResourceBytes requests.
type QueryResourceBytesCookie struct {
	*xgb.Cookie
}

// QueryResourceBytes sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryResourceBytesCookie.Reply()
func QueryResourceBytes(c *xgb.Conn, Client uint32, NumSpecs uint32, Specs []ResourceIdSpec) QueryResourceBytesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryResourceBytes' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryResourceBytesRequest(c, Client, NumSpecs, Specs), cookie)
	return QueryResourceBytesCookie{cookie}
}

// QueryResourceBytesUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryResourceBytesUnchecked(c *xgb.Conn, Client uint32, NumSpecs uint32, Specs []ResourceIdSpec) QueryResourceBytesCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryResourceBytes' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryResourceBytesRequest(c, Client, NumSpecs, Specs), cookie)
	return QueryResourceBytesCookie{cookie}
}

// QueryResourceBytesReply represents the data returned from a QueryResourceBytes request.
type QueryResourceBytesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumSizes uint32
	// padding: 20 bytes
	Sizes []ResourceSizeValue // size: ResourceSizeValueListSize(Sizes)
}

// Reply blocks and returns the reply data for a QueryResourceBytes request.
func (cook QueryResourceBytesCookie) Reply() (*QueryResourceBytesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryResourceBytesReply(buf), nil
}

// queryResourceBytesReply reads a byte slice into a QueryResourceBytesReply value.
func queryResourceBytesReply(buf []byte) *QueryResourceBytesReply {
	v := new(QueryResourceBytesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.NumSizes = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Sizes = make([]ResourceSizeValue, v.NumSizes)
	b += ResourceSizeValueReadList(buf[b:], v.Sizes)

	return v
}

// Write request to wire for QueryResourceBytes
// queryResourceBytesRequest writes a QueryResourceBytes request to a byte slice.
func queryResourceBytesRequest(c *xgb.Conn, Client uint32, NumSpecs uint32, Specs []ResourceIdSpec) []byte {
	size := xgb.Pad((12 + xgb.Pad((int(NumSpecs) * 8))))
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["X-Resource"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], Client)
	b += 4

	xgb.Put32(buf[b:], NumSpecs)
	b += 4

	b += ResourceIdSpecListBytes(buf[b:], Specs)

	return buf
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn, ClientMajor byte, ClientMinor byte) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajor, ClientMinor), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn, ClientMajor byte, ClientMinor byte) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["X-Resource"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'X-Resource'. res.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, ClientMajor, ClientMinor), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajor uint16
	ServerMinor uint16
}

// Reply blocks and returns the reply data for a QueryVersion request.
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// queryVersionReply reads a byte slice into a QueryVersionReply value.
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajor = xgb.Get16(buf[b:])
	b += 2

	v.ServerMinor = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn, ClientMajor byte, ClientMinor byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["X-Resource"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = ClientMajor
	b += 1

	buf[b] = ClientMinor
	b += 1

	return buf
}
