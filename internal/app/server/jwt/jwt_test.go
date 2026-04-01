package jwt

import "testing"

func TestGenerateAndParseJwtToken(t *testing.T) {
	token := GenerateJwtToken(&JwtCustomClaims{UserID: 1, Username: "user"})
	t.Log(token)
	if claims, err := ParseJwtToken(token); err != nil {
		t.Log(err)
	} else {
		t.Logf("user_id: %v, user_name: %v, issuer: %v", claims.UserID, claims.Username, claims.Issuer)
	}
}
