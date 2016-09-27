package tmpl

import (
	"bytes"
)

func Insert() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("<!DOCTYPE html>\n<html>\n<head>\n  <meta charset=\"UTF-8\">\n  <script src=\"//cdn.tinymce.com/4/tinymce.min.js\"></script>\n  <script src=\"/content/js/tinyMceInit.js\"></script>\n</head>\n\n<body>\n\t<form method=\"POST\" action=\"/save\">\n\t<div style=\"margin: 0 auto; width: 80%\">\n\t\t<input type=\"text\" name=\"heading\" />\n\t\t<textarea name=\"content\">Easy (and free!) You should check out our premium features.</textarea>\n\t</div>\n\t</form>\n</body>\n\n</html>")

	return _buffer.String()
}
