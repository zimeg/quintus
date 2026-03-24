package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

// Shop offers printed quintus calendars for purchase
func Shop(w http.ResponseWriter, r *http.Request) {
	content := `<ul>
<li>Order a pressing for this current year.</li>
<li>$20.26</li>
<li><a href="https://buy.stripe.com/cNiaEZ0G56p5eQFgiB9EI00" target="_blank">Checkout</a>. <a href="mailto:calendar@quintus.sh">Contact</a>.</li>
</ul>
<p><img src="/shop.png" alt="A stamped marking of upcoming month with leaf"></p>`
	if r.Header.Get("HX-Request") == "" {
		index(w, r, template.HTML(content))
		return
	}
	fmt.Fprint(w, content)
}
