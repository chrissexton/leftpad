package leftpad

import (
	"github.com/robertkrimen/otto"
)

const leftpad string = `
function leftpad (str, len, ch) {
  str = String(str);

  var i = -1;

  if (!ch && ch !== 0) ch = ' ';

  len = len - str.length;

  while (++i < len) {
    str = ch + str;
  }

  return str;
}

var result = leftpad(str, len, ch);
`

// LeftPad will return a version of str padded from the left with char until a total length.
func LeftPad(str string, length int, char string) string {
	// create javascript vm
	vm := otto.New()

	// set paramaters for left-pad function
	vm.Set("str", str)
	vm.Set("len", length)
	vm.Set("ch", char)

	// execute function in vm
	vm.Run(leftpad)

	// get and return value
	val, _ := vm.Get("result")
	return val.String()
}
