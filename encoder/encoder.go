package encoder

type Encoder interface {
	Decode([]byte, interface{}) error
	Encode(interface{}) ([]byte, error)
	String(interface{}) (string, error)
	Type() string
}
