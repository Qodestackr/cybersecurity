package rpc

import (
	"bytes"
	"fmt"
	"gopkg.in/vmihailenco/msgpack.v2"
	"net/http"
)

// Metasploit contains the information pertaining to the
// Metasploit RPC server
type Metasploit struct {
	Host  string
	User  string
	Pass  string
	Token string
}

// SessionListRequest contains all the information for invoking
// the ```sessions.list``` API method
type SessionListRequest struct {
	// _msgpack explicitly tells the msgpack encoder/decoder
	// to encode/decode data as an array as opposed to a map
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	AuthToken string
}

// Sessions stores the response returned by invoking the
// ```sessions.list``` API method
type SessionListResponse struct {
	Type        string `msgpack:"omitempty"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_pack"`
	Exploit     string `msgpack:"via_exploit"`
	Payload     strin  `msgpack:"via_payload"`
	Desc        string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	TargetHost  string `msgpack:"target_host"`
	Username    string `msgpack:"username"`
	Uuid        string `msgpack:"uuid"`
	ExploitUuid string `msgpack:"exploit_uuid"`
	// Ignore routes
}

// Login contains all the information for logging user into
// the RPC server
type Login struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

// LoginResponse will store the response values for both
// successful and unsuccessful logins
type LoginResponse struct {
	Result       string
	Token        string
	Error        bool
	ErrorClass   string
	ErrorMessage string
}

// Logout contains all the information for logging a user out
// of the RPC server
type Logout struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

// LogoutResponse stores the responses returned when logging out
type LogoutResponse struct {
	Result string
}

// New creates a new instance of type ``Metasploit`` and
// attempts to login to the RPC server by invoking the
// ```Login``` receiver function
func New(host, user, pass string) (*Metasploit, error) {
	m := &Metasploit{
		Host: host,
		User: user,
		Pass: pass,
	}

	if err := m.Login(); err != nil {
		return nil, err
	}

	return m, nil
}

// send sends requests to the RPC server
func (m *Metasploit) send(request interface{}, response interface{}) error {
	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Enode(request)
	url := fmt.Sprintf("http://%s/api", m.Host)

	resp, err := http.Post(url, "binary/message-pack", buf)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if err := msgpack.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}

// Login allows user to login to their Metasploit RPC server
func (m *Metasploit) Login() error {
	login := &Login{
		Method:   "auth.login",
		Username: m.Username,
		Password: m.Password,
	}

	var resp LoginResponse
	if err := msf.send(login, &resp); err != nil {
		return err
	}

	msf.Token = resp.Token
	return nil
}

// Logout allows a user to logout of their Metasploit RPC server
func (m *Metasploit) Logout() error {
	logout := &Logout{
		Method:      "auth.logout",
		Token:       m.Token,
		LogoutToken: m.Token,
	}

	var resp LogoutResponse
	if err := m.send(logout, &resp); err != nil {
		return err
	}

	m.Token = ""
	return nil
}

// SessionList returns all the active Metasploit sessions
func (m *Metasploit) SessionList() (map[uint32]SessionListRequest, error) {
	req := &SessionListRequest{
		Method: "session.list",
		Token:  m.Token,
	}
	resp := make(map[uint32]SessionListRequest)
	if err := m.send(req, &resp); err != nil {
		return nil, err
	}

	for id, session := range resp {
		session.Id = id
		resp[id] = session
	}

	return resp, nil
}
