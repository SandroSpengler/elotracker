// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Base() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\" class=\"uk-theme-slate\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"><link rel=\"dns-prefetch\" href=\"https://elotracker.axfert.com/\"><link rel=\"dns-prefetch\" href=\"https://opgg-static.akamaized.net/\"><link href=\"/assets/main.css?version=1\" rel=\"stylesheet\"><link rel=\"stylesheet\" href=\"/assets/frank-ui-core-1.1.0.css\"><script src=\"/assets/window.js\"></script><script src=\"/assets/theme-toggle.js\"></script><script src=\"/assets/htmx.min.js\"></script><script src=\"/assets/franken-ui-core-1.1.0.js\" type=\"module\"></script><script src=\"/assets/franken-ui-core-icon-1.1.0.js\" type=\"module\"></script><script src=\"/assets/param-converter.js\"></script><title>Elotracker</title></head><body id=\"body\" class=\"bg-background text-foreground\"><div class=\"uk-container\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
