// Generated by "sketch" utility. DO NOT EDIT
package jwk

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/blackmagic"
	"github.com/lestrrat-go/byteslice"
	"github.com/lestrrat-go/jwx/v2/cert"
	"github.com/lestrrat-go/jwx/v2/internal/json"
	"github.com/lestrrat-go/jwx/v2/jwa"
)

type RSAPublicKey interface {
	Key

	// FromRaw initializes the key internals from a Go native key type of *rsa.PublicKey
	FromRaw(*rsa.PublicKey) error
	Algorithm() jwa.KeyAlgorithm
	E() []byte
	N() []byte
	DecodeCtx() DecodeCtx
}

func newRSAPublicKey() *rsaPublicKey {
	return &rsaPublicKey{}
}

func (v *rsaPublicKey) DecodeCtx() DecodeCtx {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.dc
}

func (v *rsaPublicKey) SetDecodeCtx(dc DecodeCtx) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.dc = dc
}

func (v *rsaPublicKey) decodeExtraField(name string, dec *json.Decoder, dst interface{}) error {
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

type rsaPublicKey struct {
	mu                     sync.RWMutex
	algorithm              jwa.KeyAlgorithm
	keyID                  *string
	keyOps                 *KeyOperationList
	keyUsage               *KeyUsageType
	x509CertChain          *cert.Chain
	x509CertThumbprint     *string
	x509CertThumbprintS256 *string
	x509URL                *string
	e                      *byteslice.Buffer
	n                      *byteslice.Buffer
	dc                     DecodeCtx
	extra                  map[string]interface{}
}

// Get retrieves the value associated with a key
func (v *rsaPublicKey) Get(key string, dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.getNoLock(key, dst, false)
}

// getNoLock is a utility method that is called from Get, MarshalJSON, etc, but
// it can be used from user-supplied code. Unlike Get, it avoids locking for
// each call, so the user needs to explicitly lock the object before using,
// but otherwise should be faster than sing Get directly
func (v *rsaPublicKey) getNoLock(key string, dst interface{}, raw bool) error {
	switch key {
	case AlgorithmKey:
		if val := v.algorithm; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case KeyIDKey:
		if val := v.keyID; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case KeyOpsKey:
		if val := v.keyOps; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case KeyTypeKey:
		return blackmagic.AssignIfCompatible(dst, v.KeyType())
	case KeyUsageKey:
		if val := v.keyUsage; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.Get())
		}
	case X509CertChainKey:
		if val := v.x509CertChain; val != nil {
			return blackmagic.AssignIfCompatible(dst, val)
		}
	case X509CertThumbprintKey:
		if val := v.x509CertThumbprint; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case X509CertThumbprintS256Key:
		if val := v.x509CertThumbprintS256; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case X509URLKey:
		if val := v.x509URL; val != nil {
			return blackmagic.AssignIfCompatible(dst, *val)
		}
	case RSAEKey:
		if val := v.e; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.Bytes())
		}
	case RSANKey:
		if val := v.n; val != nil {
			if raw {
				return blackmagic.AssignIfCompatible(dst, val)
			}
			return blackmagic.AssignIfCompatible(dst, val.Bytes())
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
func (v *rsaPublicKey) Set(key string, value interface{}) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	switch key {
	case AlgorithmKey:
		object, err := jwa.KeyAlgorithmFrom(value)
		if err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.algorithm = object
	case KeyIDKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field kid, got %T`, value)
		}
		v.keyID = &converted
	case KeyOpsKey:
		converted, ok := value.(KeyOperationList)
		if !ok {
			return fmt.Errorf(`expected value of type KeyOperationList for field key_ops, got %T`, value)
		}
		v.keyOps = &converted
	case KeyTypeKey:
		// constant value is no-op on Set
		return nil
	case KeyUsageKey:
		var object KeyUsageType
		if err := object.Accept(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.keyUsage = &object
	case X509CertChainKey:
		converted, ok := value.(*cert.Chain)
		if !ok {
			return fmt.Errorf(`expected value of type *cert.Chain for field x5c, got %T`, value)
		}
		v.x509CertChain = converted
	case X509CertThumbprintKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field x5t, got %T`, value)
		}
		v.x509CertThumbprint = &converted
	case X509CertThumbprintS256Key:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field x5t#S256, got %T`, value)
		}
		v.x509CertThumbprintS256 = &converted
	case X509URLKey:
		converted, ok := value.(string)
		if !ok {
			return fmt.Errorf(`expected value of type string for field x5u, got %T`, value)
		}
		v.x509URL = &converted
	case RSAEKey:
		var object byteslice.Buffer
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.e = &object
	case RSANKey:
		var object byteslice.Buffer
		if err := object.AcceptValue(value); err != nil {
			return fmt.Errorf(`failed to accept value: %w`, err)
		}
		v.n = &object
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
func (v *rsaPublicKey) Has(name string) bool {
	switch name {
	case AlgorithmKey:
		return v.algorithm != nil
	case KeyIDKey:
		return v.keyID != nil
	case KeyOpsKey:
		return v.keyOps != nil
	case KeyTypeKey:
		return true
	case KeyUsageKey:
		return v.keyUsage != nil
	case X509CertChainKey:
		return v.x509CertChain != nil
	case X509CertThumbprintKey:
		return v.x509CertThumbprint != nil
	case X509CertThumbprintS256Key:
		return v.x509CertThumbprintS256 != nil
	case X509URLKey:
		return v.x509URL != nil
	case RSAEKey:
		return v.e != nil
	case RSANKey:
		return v.n != nil
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
func (v *rsaPublicKey) Keys() []string {
	keys := make([]string, 0, 12)
	if v.algorithm != nil {
		keys = append(keys, AlgorithmKey)
	}
	if v.keyID != nil {
		keys = append(keys, KeyIDKey)
	}
	if v.keyOps != nil {
		keys = append(keys, KeyOpsKey)
	}
	keys = append(keys, KeyTypeKey)
	if v.keyUsage != nil {
		keys = append(keys, KeyUsageKey)
	}
	if v.x509CertChain != nil {
		keys = append(keys, X509CertChainKey)
	}
	if v.x509CertThumbprint != nil {
		keys = append(keys, X509CertThumbprintKey)
	}
	if v.x509CertThumbprintS256 != nil {
		keys = append(keys, X509CertThumbprintS256Key)
	}
	if v.x509URL != nil {
		keys = append(keys, X509URLKey)
	}
	if v.e != nil {
		keys = append(keys, RSAEKey)
	}
	if v.n != nil {
		keys = append(keys, RSANKey)
	}

	if len(v.extra) > 0 {
		for k := range v.extra {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

// HasAlgorithm returns true if the field `alg` has been populated
func (v *rsaPublicKey) HasAlgorithm() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.algorithm != nil
}

// HasKeyID returns true if the field `kid` has been populated
func (v *rsaPublicKey) HasKeyID() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.keyID != nil
}

// HasKeyOps returns true if the field `key_ops` has been populated
func (v *rsaPublicKey) HasKeyOps() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.keyOps != nil
}

// HasKeyType returns true if the field `kty` has been populated
func (v *rsaPublicKey) HasKeyType() bool {
	return true
}

// HasKeyUsage returns true if the field `use` has been populated
func (v *rsaPublicKey) HasKeyUsage() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.keyUsage != nil
}

// HasX509CertChain returns true if the field `x5c` has been populated
func (v *rsaPublicKey) HasX509CertChain() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.x509CertChain != nil
}

// HasX509CertThumbprint returns true if the field `x5t` has been populated
func (v *rsaPublicKey) HasX509CertThumbprint() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.x509CertThumbprint != nil
}

// HasX509CertThumbprintS256 returns true if the field `x5t#S256` has been populated
func (v *rsaPublicKey) HasX509CertThumbprintS256() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.x509CertThumbprintS256 != nil
}

// HasX509URL returns true if the field `x5u` has been populated
func (v *rsaPublicKey) HasX509URL() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.x509URL != nil
}

// HasE returns true if the field `e` has been populated
func (v *rsaPublicKey) HasE() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.e != nil
}

// HasN returns true if the field `n` has been populated
func (v *rsaPublicKey) HasN() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return v.n != nil
}

func (v *rsaPublicKey) Algorithm() jwa.KeyAlgorithm {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.algorithm; val != nil {
		return val
	}
	return jwa.UnknownKeyAlgorithm("")
}

func (v *rsaPublicKey) KeyID() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.keyID; val != nil {
		return *val
	}
	return ""
}

func (v *rsaPublicKey) KeyOps() KeyOperationList {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.keyOps; val != nil {
		return *val
	}
	return nil
}

func (v *rsaPublicKey) KeyType() jwa.KeyType {
	return jwa.RSA
}

func (v *rsaPublicKey) KeyUsage() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.keyUsage; val != nil {
		return val.Get()
	}
	return ""
}

func (v *rsaPublicKey) X509CertChain() *cert.Chain {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.x509CertChain; val != nil {
		return val
	}
	return nil
}

func (v *rsaPublicKey) X509CertThumbprint() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.x509CertThumbprint; val != nil {
		return *val
	}
	return ""
}

func (v *rsaPublicKey) X509CertThumbprintS256() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.x509CertThumbprintS256; val != nil {
		return *val
	}
	return ""
}

func (v *rsaPublicKey) X509URL() string {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.x509URL; val != nil {
		return *val
	}
	return ""
}

func (v *rsaPublicKey) E() []byte {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.e; val != nil {
		return val.Bytes()
	}
	return []byte(nil)
}

func (v *rsaPublicKey) N() []byte {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if val := v.n; val != nil {
		return val.Bytes()
	}
	return []byte(nil)
}

// Remove removes the value associated with a key
func (v *rsaPublicKey) Remove(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	switch key {
	case AlgorithmKey:
		v.algorithm = nil
	case KeyIDKey:
		v.keyID = nil
	case KeyOpsKey:
		v.keyOps = nil
	case KeyTypeKey:
		// no-op
	case KeyUsageKey:
		v.keyUsage = nil
	case X509CertChainKey:
		v.x509CertChain = nil
	case X509CertThumbprintKey:
		v.x509CertThumbprint = nil
	case X509CertThumbprintS256Key:
		v.x509CertThumbprintS256 = nil
	case X509URLKey:
		v.x509URL = nil
	case RSAEKey:
		v.e = nil
	case RSANKey:
		v.n = nil
	default:
		delete(v.extra, key)
	}

	return nil
}

func (v *rsaPublicKey) Clone(dst interface{}) error {
	v.mu.RLock()
	defer v.mu.RUnlock()

	var extra map[string]interface{}
	if len(v.extra) > 0 {
		extra = make(map[string]interface{})
		for key, val := range v.extra {
			extra[key] = val
		}
	}
	return blackmagic.AssignIfCompatible(dst, &rsaPublicKey{
		algorithm:              v.algorithm,
		keyID:                  v.keyID,
		keyOps:                 v.keyOps,
		keyUsage:               v.keyUsage,
		x509CertChain:          v.x509CertChain,
		x509CertThumbprint:     v.x509CertThumbprint,
		x509CertThumbprintS256: v.x509CertThumbprintS256,
		x509URL:                v.x509URL,
		e:                      v.e,
		n:                      v.n,
		dc:                     v.dc,
		extra:                  extra,
	})
}

// MarshalJSON serializes rsaPublicKey into JSON.
// All pre-declared fields are included as long as a value is
// assigned to them, as well as all extra fields. All of these
// fields are sorted in alphabetical order.
func (v *rsaPublicKey) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON deserializes a piece of JSON data into rsaPublicKey.
//
// Pre-defined fields must be deserializable via "encoding/json" to their
// respective Go types, otherwise an error is returned.
//
// Extra fields are stored in a special "extra" storage, which can only
// be accessed via `Get()` and `Set()` methods.
func (v *rsaPublicKey) UnmarshalJSON(data []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.algorithm = nil
	v.keyID = nil
	v.keyOps = nil
	v.keyUsage = nil
	v.x509CertChain = nil
	v.x509CertThumbprint = nil
	v.x509CertThumbprintS256 = nil
	v.x509URL = nil
	v.e = nil
	v.n = nil

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
			case AlgorithmKey:
				var acceptValue interface{}
				if err := dec.Decode(&acceptValue); err != nil {
					return fmt.Errorf(`failed to decode vlaue for %q: %w`, AlgorithmKey, err)
				}
				var val jwa.KeyAlgorithm
				val, err = jwa.KeyAlgorithmFrom(acceptValue)
				if err != nil {
					return fmt.Errorf(`failed to accept value for %q: %w`, AlgorithmKey, err)
				}
				v.algorithm = val
			case KeyIDKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, KeyIDKey, err)
				}
				v.keyID = &val
			case KeyOpsKey:
				var val KeyOperationList
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, KeyOpsKey, err)
				}
				v.keyOps = &val
			case KeyTypeKey:
				var acceptValue interface{}
				if err := dec.Decode(&acceptValue); err != nil {
					return fmt.Errorf(`failed to decode vlaue for %q: %w`, KeyTypeKey, err)
				}
				var val jwa.KeyType
				err = val.Accept(acceptValue)
				if err != nil {
					return fmt.Errorf(`failed to accept value for %q: %w`, KeyTypeKey, err)
				}
				if val != jwa.RSA {
					return fmt.Errorf(`field %q must be jwa.RSA (got %#v)`, tok, val)
				}
			case KeyUsageKey:
				var acceptValue interface{}
				if err := dec.Decode(&acceptValue); err != nil {
					return fmt.Errorf(`failed to decode vlaue for %q: %w`, KeyUsageKey, err)
				}
				var val KeyUsageType
				err = val.Accept(acceptValue)
				if err != nil {
					return fmt.Errorf(`failed to accept value for %q: %w`, KeyUsageKey, err)
				}
				v.keyUsage = &val
			case X509CertChainKey:
				var val cert.Chain
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, X509CertChainKey, err)
				}
				v.x509CertChain = &val
			case X509CertThumbprintKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, X509CertThumbprintKey, err)
				}
				v.x509CertThumbprint = &val
			case X509CertThumbprintS256Key:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, X509CertThumbprintS256Key, err)
				}
				v.x509CertThumbprintS256 = &val
			case X509URLKey:
				var val string
				if err := dec.Decode(&val); err != nil {
					return fmt.Errorf(`failed to decode value for %q: %w`, X509URLKey, err)
				}
				v.x509URL = &val
			case RSAEKey:
				var acceptValue interface{}
				if err := dec.Decode(&acceptValue); err != nil {
					return fmt.Errorf(`failed to decode vlaue for %q: %w`, RSAEKey, err)
				}
				var val byteslice.Buffer
				err = val.AcceptValue(acceptValue)
				if err != nil {
					return fmt.Errorf(`failed to accept value for %q: %w`, RSAEKey, err)
				}
				v.e = &val
			case RSANKey:
				var acceptValue interface{}
				if err := dec.Decode(&acceptValue); err != nil {
					return fmt.Errorf(`failed to decode vlaue for %q: %w`, RSANKey, err)
				}
				var val byteslice.Buffer
				err = val.AcceptValue(acceptValue)
				if err != nil {
					return fmt.Errorf(`failed to accept value for %q: %w`, RSANKey, err)
				}
				v.n = &val
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
	if v.e == nil {
		return fmt.Errorf(`required field e is missing for object rsaPublicKey`)
	}
	if v.n == nil {
		return fmt.Errorf(`required field n is missing for object rsaPublicKey`)
	}

	if extra != nil {
		v.extra = extra
	}
	return nil
}
