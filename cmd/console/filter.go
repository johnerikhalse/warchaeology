package console

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/nlnwa/gowarc"
	"strings"
)

type recordFilter struct {
	error   bool
	recType gowarc.RecordType
}

func (r *recordFilter) filterFunc(rec interface{}) bool {
	if rec == nil {
		return false
	}
	if !r.error && r.recType == 0 {
		return true
	}
	if r.error && rec.(record).hasError {
		return true
	}
	if r.recType&rec.(record).recordType != 0 {
		return true
	}
	return false
}

func (r *recordFilter) toggleErrorFilter(g *gocui.Gui, v *gocui.View) error {
	r.error = !r.error
	v2, err := g.View("records")
	if err != nil {
		return err
	}
	r.refreshHelp(g)
	return state.records.refreshFilter(g, v2)
}

func (r *recordFilter) toggleRecordTypeFilter(g *gocui.Gui, recType gowarc.RecordType) error {
	r.recType = r.recType ^ recType
	v2, err := g.View("records")
	if err != nil {
		return err
	}
	r.refreshHelp(g)
	return state.records.refreshFilter(g, v2)
}

func (r *recordFilter) refreshHelp(g *gocui.Gui) {
	sb := strings.Builder{}
	sb.WriteString("|")
	sb.WriteString(filterString("Error", ErrorColor, r.error))
	sb.WriteString(filterString("warcInfo", WarcInfoColor, r.recType&gowarc.Warcinfo != 0))
	sb.WriteString(filterString("reQuest", RequestColor, r.recType&gowarc.Request != 0))
	sb.WriteString(filterString("Response", ResponseColor, r.recType&gowarc.Response != 0))
	sb.WriteString(filterString("Metadata", MetadataColor, r.recType&gowarc.Metadata != 0))
	sb.WriteString(filterString("reVisit", RevisitColor, r.recType&gowarc.Revisit != 0))
	sb.WriteString(filterString("reSource", ResourceColor, r.recType&gowarc.Resource != 0))
	sb.WriteString(filterString("Continuation", ContinuationColor, r.recType&gowarc.Continuation != 0))
	sb.WriteString(filterString("coNversion", ConversionColor, r.recType&gowarc.Conversion != 0))
	if v, err := g.View("help"); err == nil {
		v.Clear()
		fmt.Fprintf(v, "%s", sb.String())
	}
}

func filterString(s string, color gocui.Attribute, on bool) string {
	fg := escapeFgColor(gocui.NewRGBColor(0, 0, 0))
	bg := escapeBgColor(gocui.NewRGBColor(0, 0, 0))
	if on {
		return fmt.Sprintf("%s%s%s%s|", escapeBgColor(color), fg, s, escapeFgColor(gocui.ColorDefault))
	} else {
		return fmt.Sprintf("%s%s%s%s|", escapeFgColor(color), bg, s, escapeFgColor(gocui.ColorDefault))
	}
}