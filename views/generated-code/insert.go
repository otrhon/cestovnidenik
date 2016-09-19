package tmpl

import (
	"bytes"
)

func Insert() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("<!DOCTYPE html>\n<html>\n<head>\n  <meta charset=\"UTF-8\">\n  <script src=\"http://cdn.tinymce.com/4/tinymce.min.js\"></script>\n  <script>\n \n  tinymce.init({ \n    selector:'textarea',\n    language_url: '/content/javascript/tinyMce-cs_CZ.js',    \n    menubar: false,\n    height: '500px',\n    resize: false,\n    plugins: [ \"table \", \"code\", \"autolink\", \"link\", \"save\"], \n    width: '100%',\n    block_formats: 'Paragraph=p;Header 1=h1;Header 2=h2;Header 3=h3',\n    toolbar: 'save | undo redo | formatselect bold italic underline strikethrough | link flickrImage | alignleft aligncenter alignright | bullist numlist table | code',\n    setup: function (editor) {\n    \teditor.addButton('flickrImage', {\n        \ttext: 'Obrázek',\n\t\t\ticon: false,\n\t\t\tonclick: function () {\n\t\t\teditor.insertContent('<img style=\"display: block; padding: 5px; margin: 0 auto\" src=\"http://farm9.staticflickr.com/8792/28867667215_4507842293_m.jpg\">');\n\t\t\t},\n\t\t});\n\t\t\t\n    \teditor.on('init', function() \n\t\t{\n\t\t\tthis.getDoc().body.style.fontSize = '14px';\n\t\t\tthis.getDoc().body.style.fontFamily = 'Verdana';\n    \t});\n  },\n  });</script>\n</head>\n\n<body>\n\t<form method=\"POST\" action=\"/save\">\n\t<div style=\"margin: 0 auto; width: 80%\">\n\t\t<input type=\"text\" name=\"heading\" />\n\t\t<textarea name=\"content\">Easy (and free!) You should check out our premium features.</textarea>\n\t</div>\n\t</form>\n</body>\n\n</html>")

	return _buffer.String()
}
