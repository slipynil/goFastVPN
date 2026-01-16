package amneziawg

// Close closes the amneziawg client.
func (s WireGuard) Close() error {
	return s.client.Close()
}
