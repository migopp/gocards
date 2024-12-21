package server

// I can't figure out a better way to do this, so we are literally
// just going to re-template the same code every time...
//
// I could put this in it's own file and open it each time, but
// the performance there clearly sucks.
//
// We will just do the slightly unsanitary thing for now!

const rightCardsUI = `
<div id="ui">
	<p>{{.Word}}</p>
	<form
		hx-post="/cards/submit"
		hx-target="#ui"
		hx-swap="innerHTML"
		hx-trigger="submit"
	>
		<input
			type="text"
			name="ans"
			id="ans"
			autocomplete="off"
			hx-on:keyup="handleKey(event)"
		/>
	</form>
</div>
`

const wrongCardsUI = `
<div id="ui">
	<p style="background-color: red">{{.Word}}</p>
	<form
		hx-post="/cards/submit"
		hx-target="#ui"
		hx-swap="innerHTML"
		hx-trigger="submit"
	>
		<input
			type="text"
			name="ans"
			id="ans"
			autocomplete="off"
			hx-on:keyup="handleKey(event)"
		/>
	</form>
</div>
`
