# assertor package

Assertor is a tiny helper to facilitate parameters validation.
If at least one assertion failed an ErrValidate error is returned by Validate().
Returned error contains (wraped) all reasons why validation failed.

## Usage example

```go
    input1 := "123"
chatEnabled := "true"
chatServer := "chat.myserver.com:12345"

v := New()

v.Assert(len(input1) == 3, "string must be 3 characters long, not %d", len(input1))

if v.Assert(chatEnabled == "true", "") {
v.Assert(chatServer != "", "chat server address is missing")
}

if err := v.Validate(); err != nil {
return err
}

```
