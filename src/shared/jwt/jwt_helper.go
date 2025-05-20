package jwt

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gookit/goutil"
)

type JwtHelper struct {
	Claims       jwt.MapClaims
	RealmRoles   []string
	AccountRoles []string
	Scopes       []string
}

func NewJwtHelper(claims jwt.MapClaims) *JwtHelper {
	return &JwtHelper{
		Claims:       claims,
		RealmRoles:   ParseRealmRoles(claims),
		AccountRoles: ParseAccountRoles(claims),
		Scopes:       ParseScopes(claims),
	}
}

func (j *JwtHelper) GetUserId() (string, error) {
	return j.Claims.GetSubject()
}

func (j *JwtHelper) IsUserInRealmRole(role string) bool {
	return goutil.Contains(j.RealmRoles, role)
}

func (j *JwtHelper) TokenHasScope(scope string) bool {
	return goutil.Contains(j.Scopes, scope)
}

func ParseRealmRoles(claims jwt.MapClaims) []string {
	var realmRoles []string = make([]string, 0)

	if claim, ok := claims["realm_access"]; ok {
		if roles, ok := claim.(map[string]interface{})["roles"]; ok {
			for _, role := range roles.([]interface{}) {
				realmRoles = append(realmRoles, role.(string))
			}
		}
	}
	return realmRoles
}

func ParseAccountRoles(claims jwt.MapClaims) []string {
	var accountRoles []string = make([]string, 0)
	if claim, ok := claims["ressource_access"]; ok {
		if acc, ok := claim.(map[string]interface{})["account"]; ok {
			if role, ok := acc.(map[string]interface{})["roles"]; ok {
				for _, r := range role.([]interface{}) {
					accountRoles = append(accountRoles, r.(string))
				}
			}
		}
	}
	return accountRoles
}

func ParseScopes(claims jwt.MapClaims) []string {
	scopeStr, err := ParseString(claims, "scopes")
	if err != nil {
		return make([]string, 0)
	}
	scopes := strings.Split(scopeStr, " ")
	return scopes
}

func ParseString(claims jwt.MapClaims, key string) (string, error) {

	raw, ok := claims[key]
	if !ok {
		return "", nil
	}
	iss, ok := raw.(string)
	return iss, nil
}
