package tmpl

import (
	"bytes"
)

func Flickr() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("<html>\n<header>\n    <link rel=\"stylesheet\" href=\"/content/css/flickr.css\">\n    <script src=\"/content/js/flickr.js\" type=\"text/javascript\"></script>\n</header>\n<body>\n    <div id=\"flickr\"></div>\n</body>\n</html>")

	return _buffer.String()
}
