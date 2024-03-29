package godataagent

import (
	"io"
	"strings"

	"github.com/efjoubert/goforit/goblet/embed"

	"github.com/efjoubert/goforit/goblet"
)

//DataAgentWidget - Data Agent Widget
type DataAgentWidget struct {
	*goblet.BaseWidget
}

//RegisterAgentSession - session
func (dtaAgntwdgt *DataAgentWidget) RegisterAgentSession() {

}

func initDataAgentWidget() {
	goblet.RegisterEmbededReaders(
		"/fontawesome/fontawesome.js", embed.FontAwesomeJS,
		"/moment/moment.js", embed.MomentJS,
		"/moment/moment-with-locals.js", embed.MomentLocalJS,
		"/jquery.js", embed.JQueryJS,
		"/blockui.js", embed.BlockuiJS,
		"/webactions.js", func() io.Reader { return embed.WebactionsJS(false) },
		"/goldenlayout/css/goldenlayoutlight.css", func() io.Reader { return embed.GoldenlayoutCSS() },
		"/goldenlayout/js/goldenlayout.js", embed.GoldenlayoutJS,
		"/bootstrap/css/bootstrap.css", embed.BootstrapCSS,
		"/bootstrap/js/bootstrap.js", embed.BootstrapJS,
		"/bootstrapdatetime/css/bootstrapdatetime.css", embed.BootstrapDatetimeCSS,
		"/bootstrapdatetime/js/bootstrapdatetime.js", embed.BootstrapDatetimeJS,
		"/datatables/js/datatables.js", embed.DatatableJS,
		"/datatables/css/datatables.css", embed.DatatableCSS,
		"/hcoffcanvas/css/hcoffcanvas.css", embed.HCOffCanvasNavCSS,
		"/hcoffcanvas/js/hcoffcanvas.js", embed.HCOffCanvasNavJS,
		"/hcsticky/js/hcsticky.js", embed.HCStickyJS,
		"/knockout/knockout.js", embed.KnockoutJS,
		"/mmenu/css/mmenu.css", func() io.Reader { return embed.MMenuCSS(true) },
		"/mmenu/js/mmenu.js", func() io.Reader { return embed.MMenuJS(true) })

	goblet.RegisterServletContextPath("/", "./data").RegisterServlet("/", nil)
	goblet.RegisterWidgetPath("/data", "agent.html", newDataAgentWidget)
}

const dataagentwidgethtml string = `<!doctype html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
<link rel="stylesheet" href="/bootstrap/css/bootstrap.css">
<link rel="stylesheet" href="/bootstrapdatetime/css/bootstrapdatetime.css">
<link rel="stylesheet" href="/datatables/css/datatables.css">
<link rel="stylesheet" href="/hcoffcanvas/css/hcoffcanvas.css">
<script type="text/javascript" src="/jquery.js"></script>
<script type="text/javascript" src="/blockui.js"></script>
<script type="text/javascript" src="/webactions.js"></script>
<script type="text/javascript" src="/boootstrap/js/bootstrap.js"></script>
<script type="text/javascript" src="/boootstrapdatetime/js/bootstrapdatetime.js"></script>
<script type="text/javascript" src="/datatables/js/datatables.js"></script>
<script type="text/javascript" src="/hcoffcanvas/js/hcoffcanvas.js"></script>
</head>
<body><span>DATA AGENT</span>
<body></html>`

func dataAgentWidgetHTML() io.Reader {
	return strings.NewReader(dataagentwidgethtml)
}

func newDataAgentWidget(wdgtbrkr *goblet.WidgetBroker) goblet.Widget {
	var dtaagntwdgt = &DataAgentWidget{BaseWidget: goblet.NewBaseWidget(wdgtbrkr)}
	dtaagntwdgt.DefaultWidgetMarkupHandle(dataAgentWidgetHTML)
	return (goblet.Widget)(dtaagntwdgt)
}

func init() {
	initDataAgentWidget()
}
