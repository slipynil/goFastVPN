package amneziawg

// Close closes the amneziawg client.
func (a *awg) Close() error {
	return a.client.Close()
}
