package embed

import (
	"io"
	"strings"
)

const fontawesomealljs string = `/*!
* Font Awesome Free 5.10.1 by @fontawesome - https://fontawesome.com
* License - https://fontawesome.com/license/free (Icons: CC BY 4.0, Fonts: SIL OFL 1.1, Code: MIT License)
*/

//FontAwesomeJS FontAwesomeJS
func FontAwesomeJS() io.Reader {
	return strings.NewReader(fontawesomealljs)
}