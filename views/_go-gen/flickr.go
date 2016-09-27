package tmpl

import (
	"bytes"
)

func Flickr() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("<html>\n<header>\n    <meta charset=\"UTF-8\">\n    <link rel=\"stylesheet\" href=\"/content/css/flickr.css\">\n    <script src=\"/content/js/flickr.js\" type=\"text/javascript\"></script>\n    <script src=\"//cdn.tinymce.com/4/tinymce.min.js\"></script>\n    <script src=\"/content/js/tinyMceInit.js\"></script>\n</header>\n\n<body>\n    <div style=\"width: 60%; float: left\">\n        <form method=\"POST\" action=\"/save\">\n            <div style=\"margin: 0 auto; width: 80%\">\n                <input type=\"text\" name=\"heading\" />\n                <textarea name=\"content\">Easy (and free!) You should check out our premium features.</textarea>\n            </div>\n        </form>\n    </div>\n    <div style=\"width: 40%; float: right\">\n        <div id=\"flickrAlbums\"></div>\n        <div id=\"flickrPhotos\"></div>\n    </div>\n</body>\n\n</html>")

	return _buffer.String()
}
