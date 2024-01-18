package node

import "testing"

func TestXxx(t *testing.T) {
	client := newNodeConfigClient("https://uat.api.universalmacro.com/core", "1747545071349137408", "ABgiO0qO36O6FSsUlS9UxLVbviNopKoWhWaV9Ms8bfyqLNzlXvwIWPGDPUMZahtJ")
	client.GetConfig()
}
