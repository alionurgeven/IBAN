# IBAN

IBAN is a very lightweight IBAN validator.

Simply import the library as follows:

```go
import ( "github.com/alionurgeven/IBAN/pkg/ibanvalidator" )
```

And use it as shown below:

```go
IBAN := "AE070331234567890123456"

isValid := ibanvalidator.Validate(IBAN)
```
