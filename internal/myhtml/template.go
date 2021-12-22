package myhtml

const (
	htmlTmp = `
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>Directory listing for {{.Path}}</title>
  </head>
  <body>
    <h1>Directory listing for {{.Path}}</h1>
    <hr />
	<ul>
	  {{range .Content}}
		<li><a href="/{{.URI}}">{{.Name}}</a></li>
	  {{ end }}
	</ul>
    <hr />
  </body>
</html>
`

	notFoundTmp = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
        "http://www.w3.org/TR/html4/strict.dtd">
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html;charset=utf-8">
        <title>Error response</title>
    </head>
    <body>
        <h1>Error response</h1>
        <p>Error code: 404</p>
        <p>Message: File not found.</p>
        <p>Error code explanation: HTTPStatus.NOT_FOUND - Nothing matches the given URI.</p>
    </body>
</html>
`

	noAccessTmp = `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
        "http://www.w3.org/TR/html4/strict.dtd">
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html;charset=utf-8">
        <title>Error response</title>
    </head>
    <body>
        <h1>Error response</h1>
        <p>Error code: 500</p>
        <p>Message: No permission to access the file.</p>
        <p>Error code explanation: HTTPStatus.PERMISSION_DENIED - You have no permission to access the file.</p>
    </body>
</html>
`
)

func GetTemplate(name string) string {
	switch name {
	case "display":
		return htmlTmp
	case "404":
		return notFoundTmp
	case "500":
		return noAccessTmp
	}
	return ""
}
