package jwt

// Validate
// Confirm
// Inspect
type IJwt interface {
	InspectGoogleIdToken(accessToken string) (interface{}, error)
}