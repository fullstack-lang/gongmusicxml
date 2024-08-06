package gongmusicxml

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongmusicxml/dist/ng-github.com-fullstack-lang-gongmusicxml
var NgDistNg embed.FS
