# go-common > ext

This package contains extensions to the standard library to solve common challenges. It is intended to be imported at the root of your package. As such, it is very limited in scope.

```go
import (
	. "github.com/mattfenwick/go-common/ext"
)

func strPtr(s string) *string {
  // Note: Don't reference the package for increased readability.
	return Ptr(s)
}
```
