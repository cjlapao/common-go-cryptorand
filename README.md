# common-go-cryptorand

![Code Coverage](./badges/coverage.svg)
[![CI Release](https://github.com/cjlapao/common-go-cryptorand/actions/workflows/ci.yml/badge.svg)](https://github.com/cjlapao/common-go-cryptorand/actions/workflows/ci.yml)
[![PR Builder](https://github.com/cjlapao/common-go-cryptorand/actions/workflows/pr.yml/badge.svg)](https://github.com/cjlapao/common-go-cryptorand/actions/workflows/pr.yml)
[![CodeQL](https://github.com/cjlapao/common-go-cryptorand/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/cjlapao/common-go-cryptorand/actions/workflows/github-code-scanning/codeql)
[![Lint](https://github.com/cjlapao/common-go-cryptorand/actions/workflows/linter.yml/badge.svg)](https://github.com/cjlapao/common-go-cryptorand/actions/workflows/linter.yml)

Library to generate crypto random strings or numbers that are more random than
the usual math/rand this will be good to use in security giving an more
unpredictable randomness.

While it uses the standard math/rand it will use the crypto/rand to generate the
seed for the math/rand this will make it as the same as using crypto/rand but
with the easy of use of math/rand

If you want to use the crypto/rand directly you can use the following functions:

```go
GetAlphaNumericRandomString()
GetAlphaRandomString()
GetNumericRandomString()
GetRandomString()
```

Those will directly use the crypto/rand to generate the random string while the
`CryptoRand` will use the math/rand with the crypto/rand as seed.

## Usage
