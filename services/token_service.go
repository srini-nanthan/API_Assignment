package services

var revokedTokens = make(map[string]bool)

func RevokeToken(token string) {
    revokedTokens[token] = true
}

func IsTokenRevoked(token string) bool {
    return revokedTokens[token]
}
