// Code generated by qtc from "dashboard.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/dashboard.qtpl:3
package templates

//line templates/dashboard.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/dashboard.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/dashboard.qtpl:4
type Dashboard struct {
}

//line templates/dashboard.qtpl:8
func (p *Dashboard) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/dashboard.qtpl:8
	qw422016.N().S(`
Tester - Dashboard
`)
//line templates/dashboard.qtpl:10
}

//line templates/dashboard.qtpl:10
func (p *Dashboard) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/dashboard.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/dashboard.qtpl:10
	p.StreamTitle(qw422016)
//line templates/dashboard.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line templates/dashboard.qtpl:10
}

//line templates/dashboard.qtpl:10
func (p *Dashboard) Title() string {
//line templates/dashboard.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/dashboard.qtpl:10
	p.WriteTitle(qb422016)
//line templates/dashboard.qtpl:10
	qs422016 := string(qb422016.B)
//line templates/dashboard.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/dashboard.qtpl:10
	return qs422016
//line templates/dashboard.qtpl:10
}

//line templates/dashboard.qtpl:13
func (p *Dashboard) StreamBody(qw422016 *qt422016.Writer) {
//line templates/dashboard.qtpl:13
	qw422016.N().S(`
<div class="overall">
  <h1 class="h3">Overall Results <small class="text-muted">(last 30d)</small></h1>

  <div class="row">
    <div class="col">
      {{ template "run_summary_month" .OverallMonthlyRunSummary }}
    </div>
  </div>
</div>

<hr>

<div class="packages">
  <h1 class="h3">Results by Package  <small class="text-muted">(last 24h)</small></h1>
  <div class="row row-cols-1 row-cols-md-2 row-cols-lg-3">
    {{ range .Packages }}
    {{ $pkgSummary := index $.DailyPackageRunSummaries .Name }}
    <div class="col mb-2">
      <h2 class="h4"><a href="/packages/{{ .Name }}">{{ .Name }}</a></h2>
      {{ if $pkgSummary }}
      {{ template "package_run_summary_day" $pkgSummary }}
      {{ else }}
      <div class="border p-1">
        <div class="d-flex" style="min-height: 100%;">
          No results yet...
        </div>
      </div>
      {{ end }}
    </div>
    {{else}}
    <div class="col">
      <h5>No results yet</h5>
      <p>Kick off a test run and publish the results...</p>
    </div>
    {{ end }}
  </div>
</div>
`)
//line templates/dashboard.qtpl:51
}

//line templates/dashboard.qtpl:51
func (p *Dashboard) WriteBody(qq422016 qtio422016.Writer) {
//line templates/dashboard.qtpl:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/dashboard.qtpl:51
	p.StreamBody(qw422016)
//line templates/dashboard.qtpl:51
	qt422016.ReleaseWriter(qw422016)
//line templates/dashboard.qtpl:51
}

//line templates/dashboard.qtpl:51
func (p *Dashboard) Body() string {
//line templates/dashboard.qtpl:51
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/dashboard.qtpl:51
	p.WriteBody(qb422016)
//line templates/dashboard.qtpl:51
	qs422016 := string(qb422016.B)
//line templates/dashboard.qtpl:51
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/dashboard.qtpl:51
	return qs422016
//line templates/dashboard.qtpl:51
}