package tunnel

import (
	"encoding/binary"
	"errors"
	"github.com/black40x/tunl-core/commands"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"sync"
	"time"
)

const prefixSize = 4

var ErrorConnectionClosed = errors.New("connection closed")

type CommandCallback func(cmd *commands.Transfer)
type DisconnectCallback func()
type ErrorCallback func(err error)

type TunlConn struct {
	ID             string
	Conn           net.Conn
	onDisconnected DisconnectCallback
	onCommand      CommandCallback
	onError        ErrorCallback
	ExpireAt       time.Time
	ConnectedAt    time.Time
	mu             sync.Mutex
}

func NewTunlConn(conn net.Conn) *TunlConn {
	return &TunlConn{
		Conn:        conn,
		ConnectedAt: time.Now(),
	}
}

func (t *TunlConn) SetOnDisconnected(c DisconnectCallback) {
	t.onDisconnected = c
}

func (t *TunlConn) SetOnCommand(c CommandCallback) {
	t.onCommand = c
}

func (t *TunlConn) SetOnError(c ErrorCallback) {
	t.onError = c
}

func (t *TunlConn) handleError(err error) {
	if t.onError != nil {
		t.onError(err)
	}
}

func (t *TunlConn) handleCommand(trans *commands.Transfer) {
	if t.onCommand != nil {
		t.onCommand(trans)
	}
}

func (t *TunlConn) handleDisconnected() {
	if t.onDisconnected != nil {
		t.onDisconnected()
	}
}

func (t *TunlConn) HandleExpire() {
	for {
		if time.Now().Unix() >= t.ExpireAt.Unix() {
			t.Send(&commands.SessionExpired{})
			time.Sleep(time.Second)
			t.Close()
			break
		}
		time.Sleep(time.Second)
	}
}

func (t *TunlConn) HandleConnection() {
	defer t.Conn.Close()
	for {
		trans := &commands.Transfer{}
		data, err := t.Read()
		if err != nil {
			if err == io.EOF {
				t.handleDisconnected()
				break
			}
			continue
		}
		err = proto.Unmarshal(data, trans)
		if err != nil {
			t.handleError(err)
		} else {
			t.handleCommand(trans)
		}
	}
}

func (t *TunlConn) Read() ([]byte, error) {
	if t.Conn == nil {
		return nil, ErrorConnectionClosed
	}

	prefix := make([]byte, prefixSize)
	_, err := io.ReadFull(t.Conn, prefix)

	if err != nil {
		return nil, err
	}

	totalDataLength := binary.BigEndian.Uint32(prefix[:])

	data := make([]byte, totalDataLength-prefixSize)

	_, err = io.ReadFull(t.Conn, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *TunlConn) makePacket(data []byte) []byte {
	buf := make([]byte, prefixSize+len(data))
	binary.BigEndian.PutUint32(buf[:prefixSize], uint32(prefixSize+len(data)))
	copy(buf[prefixSize:], data[:])

	return buf
}

func (t *TunlConn) Write(data []byte) (n int, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Conn == nil {
		return 0, errors.New("connection closed")
	}

	return t.Conn.Write(t.makePacket(data))
}

func (t *TunlConn) Send(m proto.Message) (n int, err error) {
	trans := &commands.Transfer{}
	switch m.(type) {
	case *commands.Connect:
		trans.Command = &commands.Transfer_Connect{
			Connect: m.(*commands.Connect),
		}
	case *commands.HttpRequest:
		trans.Command = &commands.Transfer_HttpRequest{
			HttpRequest: m.(*commands.HttpRequest),
		}
	case *commands.HttpResponse:
		trans.Command = &commands.Transfer_HttpResponse{
			HttpResponse: m.(*commands.HttpResponse),
		}
	case *commands.BodyChunk:
		trans.Command = &commands.Transfer_BodyChunk{
			BodyChunk: m.(*commands.BodyChunk),
		}
	case *commands.SessionExpired:
		trans.Command = &commands.Transfer_SessionExpired{
			SessionExpired: m.(*commands.SessionExpired),
		}
	case *commands.ServerFull:
		trans.Command = &commands.Transfer_ServerFull{
			ServerFull: m.(*commands.ServerFull),
		}
	}

	buf, err := proto.Marshal(trans)
	if err != nil {
		return 0, err
	}

	return t.Write(buf)
}

func (t *TunlConn) SetExpireAt(e time.Time) {
	t.ExpireAt = e
}

func (t *TunlConn) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.Conn.Close()
}
