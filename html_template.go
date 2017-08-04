package main

const issuetmpl = `<html>
<head>
<style>
table, th, td {
  border: 1px solid black;
  border-collapse: collapse;
}
u {
text-decoration: underline;
}
</style>
</head>
<body>
<h3><u>Total Count</u> {{.TotalCount}}</h3>
<h3><u>Organization</u> {{.Organization}}</h3>
<table style="width:100%">
<tr>
  <th>Repository URL</th>
  <th>Title</th>
  <th>Status</th>
  <th>Author</th>
  <th>Milestone</th>
  <th>Assignees</th>
  <th>Labels</th>
</tr>
{{range .Issues}}
<tr>
  <td>{{.RepositoryURL}}</td>
  <td>{{.Title}}</td>
  <td>{{.Status}}</td>
  <td>{{.Author}}</td>
  <td>{{.Milestone}}</td>
  <td>{{.Assignees}}</td>
  <td>{{.Labels}}</td>
</tr>
{{ end }}
</table>

</body>
</html>
`
