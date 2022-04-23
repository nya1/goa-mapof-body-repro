
goa version 3.7.2

[Design](./design/design.go)

The field name `anyobject` with Attribute `MapOf(String, Any)` should describe
a map of string to any unknown value.


For HTTP requests, the field `anyobject` can be sent like the following (i.e. any json object):

```jsonc
{
  "anyobject": {
    "testkey": 10,
    "test1": "testval",
    "myarr": [
      {
        "nestedk": {
          "nestedinner": [
            {
              "leafk": "testleaf"
            }
          ]
        }
      }
    ]
  }
}
```

The service will correctly receive the mapping, but the [OpenAPI generated](./gen/http/openapi3.yaml#L53)
it's not describing the correct specification for the `anyobject` field, the following is created:

```yaml
    MultiplyRequestBody:
      type: object
      properties:
        anyobject:
          type: object
          description: pass any json object
          example:
            ob:
              ok: 1
              testk: testl
            test: 10
          additionalProperties: # <----
            type: string
            example: Eos qui exercitationem sed non.
            format: binary # <----
```

According to [OpenAPI 3.0 specification the binary string format should be used to describe a binary data (for files)](https://swagger.io/docs/specification/data-models/data-types/#format)



## Expected

`anyobject` body field should be described by OpenAPI as a [Free-Form Object](https://swagger.io/docs/specification/data-models/data-types/#free-form) without the string type and binary format (without an example value)


```yaml
    MultiplyRequestBody:
      type: object
      properties:
        anyobject:
          type: object
          description: pass any json object
          example:
            ob:
              ok: 1
              testk: testl
            test: 10
```
