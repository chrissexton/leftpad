LeftPad
=======

LeftPad is a wrapper porting the ever useful left-pad package from Javascript to Go.

  1. uses existing, production ready, code from javascript
  2. 100% not trolling

    go get -u github.com/jamescun/leftpad


Usage
-----

```go
str := leftpad.LeftPad("foo", 6, " ")
```

