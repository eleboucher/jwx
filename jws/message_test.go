package jws_test

import (
	"testing"

	"github.com/lestrrat-go/jwx/internal/json"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		const src = `{
  "payload": "eyJpc3MiOiJqb2UiLA0KICJleHAiOjEzMDA4MTkzODAsDQogImh0dHA6Ly9leGFtcGxlLmNvbS9pc19yb290Ijp0cnVlfQ",
  "signatures": [
    {
      "header": {
        "kid": "2010-12-29"
      },
      "protected": "eyJhbGciOiJSUzI1NiJ9",
      "signature": "cC4hiUPoj9Eetdgtv3hF80EGrhuB__dzERat0XF9g2VtQgr9PJbu3XOiZj5RZmh7AAuHIm4Bh-0Qc_lF5YKt_O8W2Fp5jujGbds9uJdbF9CUAr7t1dnZcAcQjbKBYNX4BAynRFdiuB--f_nZLgrnbyTyWzO75vRK5h6xBArLIARNPvkSjtQBMHlb1L07Qe7K0GarZRmB_eSN9383LcOLn6_dO--xi12jzDwusC-eOkHWEsqtFZESc6BfI7noOPqvhJ1phCnvWh6IeYI2w9QOYEUipUTI8np6LbgGY9Fs98rqVt5AXLIhWkWywlVmtVrBp0igcN_IoypGlUPQGe77Rw"
    },
    {
      "header": {
        "kid": "e9bc097a-ce51-4036-9562-d2ade882db0d"
      },
      "protected": "eyJhbGciOiJFUzI1NiJ9",
      "signature": "DtEhU3ljbEg8L38VWAfUAqOyKAM6-Xx-F4GawxaepmXFCgfTjDxw5djxLa8ISlSApmWQxfKTUJqPP3-Kg6NU1Q"
    }
  ]
}`

		var m jws.Message
		if !assert.NoError(t, json.Unmarshal([]byte(src), &m), `json.Unmarshal should succeed`) {
			return
		}

		buf, err := json.MarshalIndent(m, "", "  ")
		if !assert.NoError(t, err, `json.Marshal should succeed`) {
			return
		}

		if !assert.Equal(t, src, string(buf), `roundtrip should match`) {
			return
		}
	})
}
