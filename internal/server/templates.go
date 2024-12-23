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
	<div id="word-ui" class="neutral">
		<p>{{.Word}}</p>
	</div>
	<div id="box-ui">
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
</div>
`

const wrongCardsUI = `
<div id="ui">
	<div id="word-ui" class="wrong">
		<p>{{.Word}}</p>
	</div>
	<div id="box-ui">
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
</div>
`

const deckSelect = `
<select name="decks" id="decks" required>
	<option value="" disabled selected>Select a deck</option>
	{{range $index, $item := .Decks}}
	<option value="{{$index}}">{{$item.Deck.DeckName}}</option>
	{{end}}
</select>
`

const startButton = `
<form action="/cards" method="get">
	<button type="submit">Start</button>
</form>
`

const homeButton = `
<p>{{.Ratio}}</p>
<form action="/" method="get">
	<button type="submit">Home</button>
</form>
`
