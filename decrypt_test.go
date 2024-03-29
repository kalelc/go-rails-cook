package railscook

import (
	"testing"
)

func TestRails5CookieDecrypt(t *testing.T) {
	value := "O+yqkLpw2G5V+6vPEAilQ4D5sIsstQsm6+TM2NWf3mpWkfNXn89Za8ov5CdJKSHpzgM4+k7LEeWObDe6LHpeGz3GhSaEmAQdqnzv3ch9FJYVq+yFaxXPaAkN2BnoI6mhXaWRdLmHIemv/GNQAtQ5S42Bzlrjl6YP86LulEzruYnBOPuvj0fQLn7AB2qKHpxoNDc0e0lbs1tDY3jkXwGRjkW3BRI2CcpJuAZx6TOblIL5i6bQrWO1TNSP/Fag7uRwOVxAg84WewkTQQmzGk8xaqlnlp1y--dIB8fY1JEsUsv4sR--6tEpDOAgdax7ZQr4tTOB6g=="
	secret := "836fa3665997a860728bcb9e9a1e704d427cfc920e79d847d79c8a9a907b9e965defa4154b2b86bdec6930adbe33f21364523a6f6ce363865724549fdfc08553"
	salt := "ef9834ec009b4f01605933f35c53e29331f99a057d9a6f34c8cfcdb37179acc59230d9c3b08b4b47055c2ee8e8d5fd8fde4b8724a8be316b2543b8f3b09dfe16"

	sessionID := "05fff49bbc0f9e6e8101e4e81b947ebe"

	cookie := Cookie{Value: value, SecretKeyBase: secret, Salt: salt}
	cookie.Decrypt()

	if cookie.Content.SessionID != sessionID {
		t.Errorf("session_id in cookie doesn't match: got %v want %v", cookie.Content.SessionID, sessionID)
	}
}
