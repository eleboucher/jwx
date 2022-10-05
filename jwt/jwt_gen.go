// Generated by "sketch" utility. DO NOT EDIT
package jwt

import (
	"bytes"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/lestrrat-go/blackmagic"
	"github.com/lestrrat-go/jwx/v2/internal/json"
	"github.com/lestrrat-go/jwx/v2/jwt/internal/types"
)

type stdToken struct {
	mu         sync.RWMutex
	audience   types.Audience
	expiration *types.NumericDate
	issuedAt   *types.NumericDate
	issuer     *string
	jwtID      *string
	notBefore  *types.NumericDate
	subject    *string
	dc         DecodeCtx
	extra      map[string]interface{}
}

// These constants are used when the JSON field name is used.
// Their use is not strictly required, but certain linters
// complain about repeated constants, and therefore internally
// this used throughout
const (
	AudienceKey   = "aud"
	ExpirationKey = "exp"
	IssuedAtKey   = "iat"
	IssuerKey     = "iss"
	JwtIDKey      = "jti"
	NotBeforeKey  = "nbf"
	SubjectKey    = "sub"
)

// Get retrieves the value associated with a key
func (v *stdToken) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *stdToken) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case AudienceKey:
		if val := v.audience; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case ExpirationKey:
		if val := v.expiration; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case IssuedAtKey:
		if val := v.issuedAt; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case IssuerKey:
		if val := v.issuer; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case JwtIDKey:
		if val := v.jwtID; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case NotBeforeKey:
		if val := v.notBefore; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.GetValue())
		}
	case SubjectKey:
		if val := v.subject; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	default:
		if v.extra != nil {
			val, ok := v.extra[key]
			if ok {
				return blackmagic.AssignIfCompatible(dst, val)
			}
		}
	}
	return fmt.Errorf(`no such key %q`, key)
}

// Set sets the value of the specified field. The name must be a JSON
// field name, not the Go name
func (v *stdToken) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case AudienceKey:
		var object types.Audience
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.audience = object
	case ExpirationKey:
		var object types.NumericDate
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.expiration = &object
	case IssuedAtKey:
		var object types.NumericDate
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.issuedAt = &object
	case IssuerKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field iss, got %T`, value)
		}
		v.issuer = &converted
	case JwtIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field jti, got %T`, value)
		}
		v.jwtID = &converted
	case NotBeforeKey:
		var object types.NumericDate
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.notBefore = &object
	case SubjectKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field sub, got %T`, value)
		}
		v.subject = &converted
	default:
		if v.extra == nil {
			v.extra = make(map[string]interface{})
		}
		v.extra[key] = value
	}
	return nil
}

// Has returns true if the field specified by the argument has been populated.
// The field name must be the JSON field name, not the Go-structure's field name.
func (v *stdToken) Has(name string) bool {
	switch name {
	case AudienceKey:
		return v.audience != nil
	case ExpirationKey:
		return v.expiration != nil
	case IssuedAtKey:
		return v.issuedAt != nil
	case IssuerKey:
		return v.issuer != nil
	case JwtIDKey:
		return v.jwtID != nil
	case NotBeforeKey:
		return v.notBefore != nil
	case SubjectKey:
		return v.subject != nil
	default:
		if v.extra != nil {
			if _, ok := v.extra[name]; ok {
				return true
			}
		}
		return false
	}
}

// Keys returns a slice of string comprising of JSON field names whose values
// are present in the object.
func (v *stdToken) Keys() []string {
	keys := make([]string, 0, 8)
	if v.audience != nil {
		keys = append(keys, AudienceKey)
	}
	if v.expiration != nil {
		keys = append(keys, ExpirationKey)
	}
	if v.issuedAt != nil {
		keys = append(keys, IssuedAtKey)
	}
	if v.issuer != nil {
		keys = append(keys, IssuerKey)
	}
	if v.jwtID != nil {
		keys = append(keys, JwtIDKey)
	}
	if v.notBefore != nil {
		keys = append(keys, NotBeforeKey)
	}
	if v.subject != nil {
		keys = append(keys, SubjectKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasAudience returns true if the field `aud` has been populated
func (v *stdToken) HasAudience() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.audience != nil
}

// HasExpiration returns true if the field `exp` has been populated
func (v *stdToken) HasExpiration() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.expiration != nil
}

// HasIssuedAt returns true if the field `iat` has been populated
func (v *stdToken) HasIssuedAt() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.issuedAt != nil
}

// HasIssuer returns true if the field `iss` has been populated
func (v *stdToken) HasIssuer() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.issuer != nil
}

// HasJwtID returns true if the field `jti` has been populated
func (v *stdToken) HasJwtID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.jwtID != nil
}

// HasNotBefore returns true if the field `nbf` has been populated
func (v *stdToken) HasNotBefore() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.notBefore != nil
}

// HasSubject returns true if the field `sub` has been populated
func (v *stdToken) HasSubject() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.subject != nil
}

// Audience represents the `aud` field as described in https://tools.ietf.org/html/rfc7519#section-4.1.3
func (v *stdToken) Audience() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.audience; val != nil {
		return val.GetValue()
	}
	return nil
}

// Expiration represents the `exp` field as described in https://tools.ietf.org/html/rfc7519#section-4.1.4
func (v *stdToken) Expiration() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.expiration; val != nil {
		return val.GetValue()
	}
	return time.Time{}
}

// IssuedAt represents the `iat` field as described in https://tools.ietf.org/html/rfc7519#section-4.1.6
func (v *stdToken) IssuedAt() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.issuedAt; val != nil {
		return val.GetValue()
	}
	return time.Time{}
}

// Issuer represents the `iss` field as described in https://tools.ietf.org/html/rfc7519#section-4.1.1
func (v *stdToken) Issuer() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.issuer; val != nil {
		return *val
	}
	return ""
}

// JwtID represents the `jti` field as described in https://tools.ietf.org/html/rfc7519#section-4.1.7
func (v *stdToken) JwtID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.jwtID; val != nil {
		return *val
	}
	return ""
}

// NotBefore represents the `nbf` field as described in https://tools.ietf.org/html/rfc7519#section-4.1.5
func (v *stdToken) NotBefore() time.Time {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.notBefore; val != nil {
		return val.GetValue()
	}
	return time.Time{}
}

// Subject represents the `sub` field as described in https://tools.ietf.org/html/rfc7519#section-4.1.2
func (v *stdToken) Subject() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.subject; val != nil {
		return *val
	}
	return ""
}

// Remove removes the value associated with a key
func (v *stdToken) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case AudienceKey:
		v.audience = nil
	case ExpirationKey:
		v.expiration = nil
	case IssuedAtKey:
		v.issuedAt = nil
	case IssuerKey:
		v.issuer = nil
	case JwtIDKey:
		v.jwtID = nil
	case NotBeforeKey:
		v.notBefore = nil
	case SubjectKey:
		v.subject = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *stdToken) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	extra := make(map[string]interface{})
	for key, val := range v.extra {
		extra[key] = val
	}
	return blackmagic.AssignIfCompatible(dst, &stdToken{
		audience:   v.audience,
		expiration: v.expiration,
		issuedAt:   v.issuedAt,
		issuer:     v.issuer,
		jwtID:      v.jwtID,
		notBefore:  v.notBefore,
		subject:    v.subject,
		dc:         v.dc,
		extra:      extra,
	})
}

// MarshalJSON serializes stdToken into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *stdToken) MarshalJSON() ([]byte, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	buf.WriteByte('{')
	for i, k := range v.Keys() {
		var val interface{}
		if err := v.getNoLock(k, &val, true); err != nil {
			return nil, fmt.Errorf(`failed to retrieve value for field %q: %w`, k, err)
		}

		if i > 0 {
			buf.WriteByte(',')
		}
		if err := enc.Encode(k); err != nil {
			return nil, fmt.Errorf(`failed to encode map key name: %w`, err)
		}
		buf.WriteByte(':')
		if err := enc.Encode(val); err != nil {
			return nil, fmt.Errorf(`failed to encode map value for %q: %w`, k, err)
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON deserializes a piece of JSON data into stdToken.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *stdToken) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.audience = nil
	v.expiration = nil
	v.issuedAt = nil
	v.issuer = nil
	v.jwtID = nil
	v.notBefore = nil
	v.subject = nil

	dec := json.NewDecoder(bytes.NewReader(data))
	var extra map[string]interface{}

LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf(`error reading JSON token: %w`, err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			if tok == '}' { // end of object
				break LOOP
			}
			// we should only get into this clause at the very beginning, and just once
			if tok != '{' {
				return fmt.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string:
			switch tok {
			case AudienceKey:
				var val types.Audience
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, AudienceKey, err)
				}
				v.audience = val
			case ExpirationKey:
				var val types.NumericDate
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, ExpirationKey, err)
				}
				v.expiration = &val
			case IssuedAtKey:
				var val types.NumericDate
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, IssuedAtKey, err)
				}
				v.issuedAt = &val
			case IssuerKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, IssuerKey, err)
				}
				v.issuer = &val
			case JwtIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, JwtIDKey, err)
				}
				v.jwtID = &val
			case NotBeforeKey:
				var val types.NumericDate
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, NotBeforeKey, err)
				}
				v.notBefore = &val
			case SubjectKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, SubjectKey, err)
				}
				v.subject = &val
			default:
				var val interface{}
				if err := v.decodeExtraField(tok, dec, &val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, tok, err)
				}
				if extra == nil {
					extra = make(map[string]interface{})
				}
				extra[tok] = val
			}
		}
	}

	if extra != nil {
		v.extra = extra
	}
	return nil
}

type Builder struct {
	mu     sync.Mutex
	err    error
	once   sync.Once
	object *stdToken
}

// NewBuilder creates a new Builder instance.
// Builder is safe to be used uninitialized as well.
func NewBuilder() *Builder {
	return &Builder{}
}
func (b *Builder) initialize() {
	b.err = nil
	b.object = &stdToken{}
}
func (b *Builder) Audience(in []string) *Builder {
	return b.Claim(AudienceKey, in)
}
func (b *Builder) Expiration(in time.Time) *Builder {
	return b.Claim(ExpirationKey, in)
}
func (b *Builder) IssuedAt(in time.Time) *Builder {
	return b.Claim(IssuedAtKey, in)
}
func (b *Builder) Issuer(in string) *Builder {
	return b.Claim(IssuerKey, in)
}
func (b *Builder) JwtID(in string) *Builder {
	return b.Claim(JwtIDKey, in)
}
func (b *Builder) NotBefore(in time.Time) *Builder {
	return b.Claim(NotBeforeKey, in)
}
func (b *Builder) Subject(in string) *Builder {
	return b.Claim(SubjectKey, in)
}

// Claim sets the value of any field. The name should be the JSON field name.
// Type check will only be performed for pre-defined types
func (b *Builder) Claim(name string, value interface{}) *Builder {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return b
	}

	if err := b.object.Set(name, value); err != nil {
		b.err = err
	}
	return b
}
func (b *Builder) Build() (Token, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.once.Do(b.initialize)
	if b.err != nil {
		return nil, b.err
	}
	obj := b.object
	b.once = sync.Once{}
	b.once.Do(b.initialize)
	return obj, nil
}
func (b *Builder) MustBuild() Token {
	object, err := b.Build()
	if err != nil {
		panic(err)
	}
	return object
}

// New creates an empty token
func New() Token {
	return &stdToken{}
}

func (v *stdToken) DecodeCtx() DecodeCtx {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.dc
}

func (v *stdToken) SetDecodeCtx(dc DecodeCtx) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.dc = dc
}

func (v *stdToken) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
	if dc := v.dc; dc != nil {
		if localReg := dc.Registry(); localReg != nil {
			decoded, err := localReg.Decode(dec, name)
			if err == nil {
				if err := blackmagic.AssignIfCompatible(dst, decoded); err != nil {
					return fmt.Errorf(`failed to assign decoded value for %q: %w`, name, err)
				}
				return nil
			}
		}
	}

	decoded, err := registry.Decode(dec, name)
	if err == nil {
		if err := blackmagic.AssignIfCompatible(dst, decoded); err != nil {
			return fmt.Errorf(`failed to assign decoded value for %q: %w`, name, err)
		}
		return nil
	}

	return fmt.Errorf(`failed to decode field %q: %w`, name, err)
}
