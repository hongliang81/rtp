package codecs

// PCMAPayloader payloads G711A packets
type PCMAPayloader struct{}

// Payload fragments an G711A packet across one or more byte arrays
func (p *PCMAPayloader) Payload(mtu int, payload []byte) [][]byte {
	var out [][]byte
	if payload == nil || mtu <= 0 {
		return out
	}

	for len(payload) > mtu {
		o := make([]byte, mtu)
		copy(o, payload[:mtu])
		payload = payload[mtu:]
		out = append(out, o)
	}
	o := make([]byte, len(payload))
	copy(o, payload)
	return append(out, o)
}
