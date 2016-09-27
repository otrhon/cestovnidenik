this.tinymce.init({
  selector: 'textarea',
  language_url: '/content/js/tinyMce-cs_CZ.js',
  menubar: false,
  resize: false,
  plugins: ["table ", "code", "autolink", "link", "save"],
  width: '100%',
  save_enablewhendirty: false,
  height: toScreenHeight(120),
  block_formats: 'Paragraph=p;Header 1=h1;Header 2=h2;Header 3=h3',
  toolbar: 'save | undo redo | formatselect bold italic underline strikethrough | link flickrImage | alignleft aligncenter alignright | bullist numlist table | code',
  setup: function (editor) {
    editor.addButton('flickrImage', {
      text: 'Obrázek',
      icon: false,
      onclick: function () {
        editor.insertContent('<img style="display: block; padding: 5px; margin: 0 auto" src="http://farm9.staticflickr.com/8792/28867667215_4507842293_m.jpg">');
      },
    });

    editor.on('init', function () {
      this.getDoc().body.style.fontSize = '14px';
      this.getDoc().body.style.fontFamily = 'Verdana';
    });
  },
});

function toScreenHeight(minus) {
    var height;

    if (typeof(window.innerHeight) == "number") //non-IE
        height = window.innerHeight;
    else if (document.documentElement && document.documentElement.clientHeight) //IE 6+ strict mode
        height = document.documentElement.clientHeight;
    else if (document.body && document.body.clientHeight) //IE 4 compatible / IE quirks mode
        height = document.body.clientHeight;

    return height - minus
}