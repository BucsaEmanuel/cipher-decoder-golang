- [Hyperskill Go - Cipher Decoder project](#hyperskill-go-cipher-decoder-project)
    - [About the project](#about-the-project)
        - [Status](#status)
        - [See also](#see-also)
    - [Getting started](#getting-started)

# Hyperskill Go - Cipher Decoder project

## About the project

This is the implementation of the Go Cipher Decoder project on Hyperskill.

### Status

The project is posted as it currently is, having passed the last step of the project ( 3 / 3 ).

### See also

* [Hyperskill - Cipher Decoder (Go)](https://hyperskill.org/projects/236)

## Getting started

Either
`go build main.go`
then `./main`

or
`go run main.go`

## Examples
### Example 1: success (Bob has chosen b=7 here, resulting in s = 565 )
```
> g is 245 and p is 999
OK
> A is 232
B is 236
Pbee rhn ftkkr fx?  // Will you marry me?
> Rxta, hdtr!  // Yeah, okay!
Zkxtm!  // Great!
```

### Example 2: better luck next time (Bob has chosen b=7 here, resulting in s=565 )

```
> g is 245 and p is 999
OK
> A is 232
B is 236
Pbee rhn ftkkr fx?  // Will you marry me?
> Exm'l ux ykbxgwl.  // Let's be friends.
Patm t ibmr!  // What a pity!
```