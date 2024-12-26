package templates

const CardRight = `
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
				oninput="handleKey(event)"
			/>
		</form>
	</div>
</div>
`

const CardWrong = `
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
				oninput="handleKey(event)"
			/>
		</form>
	</div>
</div>
`

const DeckSelect = `
<select name="decks" id="decks" required>
	<option value="" disabled selected>Select a deck</option>
	{{range $index, $item := .Decks}}
	<option value="{{$index}}">{{$item.Deck.DeckName}}</option>
	{{end}}
</select>
`

const Start = `
<form action="/cards" method="get">
	<button type="submit">Start</button>
</form>
`

const End = `
<p>{{.Ratio}}</p>
<form action="/" method="get">
	<button type="submit">Home</button>
</form>
`

const Home = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<title>gocards</title>
		<script
			src="https://unpkg.com/htmx.org@2.0.3"
			integrity="sha384-0895/pl2MU10Hqc6jd4RvrthNlDiE9U1tWmX7WRESftEDRosgxNsQG/Ze9YMRzHq"
			crossorigin="anonymous"></script>
		<style>
			@import url("https://fonts.googleapis.com/css2?family=Noto+Sans+JP:wght@100..900&display=swap");
		
			body {
				max-width: 1000px;
				margin: 0 auto;
				padding: 0;
			}

			#home-ui {
				padding: 2vh 5vw 0;
			}

			#header {
				display: flex;
				justify-content: space-between;
				align-items: center;
				margin-bottom: 3rem;
			}

			#gocards-title {
				font-size: 2rem;
				font-family: "Noto Sans JP", sans-serif;
				font-weight: bold;
				margin: 0;
				padding: 0;
			}

			#deck-select {
				display: flex;
				flex-direction: column;
				justify-content: space-between;
				padding: 1rem 1rem 6rem 1rem;
				margin-bottom: 3rem;
				background-color: antiquewhite;
				border-radius: 15px;
			}

			#deck-upload {
				display: flex;
				flex-direction: column;
				justify-content: space-between;
				padding: 1rem 1rem 6rem 1rem;
				margin-bottom: 3rem;
				background-color: antiquewhite;
				border-radius: 15px;
			}

			#deck-upload-form {
			}

			h2 {
				font-size: 1.25rem;
				margin: 0 0 1rem 0;
				padding: 0;
			}
		</style>
	</head>
	<body>
		<div id="home-ui">
			<header id="header">
				<div>
					<h1 id="gocards-title">ゴーカード</h1>
				</div>
				<div><nav></nav></div>
			</header>
			<div id="deck-select">
				<h2>Select a deck</h2>
				<div id="deck-select-form">
					<form
						hx-post="/decks/select"
						hx-trigger="change"
						hx-target="#start"
						hx-swap="innerHTML">
						<select name="decks" id="decks" required>
							<option value="" disabled selected>エンプティー？</option>
							{{range $index, $item := .Decks}}
							<option value="{{$index}}">{{$item.Deck.DeckName}}</option>
							{{end}}
						</select>
					</form>
				</div>
				<div id="start"></div>
			</div>
			<div id="deck-upload">
				<h2>Or, upload a new one</h2>
				<form
					hx-post="/decks/upload"
					hx-trigger="submit"
					hx-target="#decks"
					hx-swap="outerHTML"
					enctype="multipart/form-data"
					id="deck-upload-form">
					<input type="file" name="deck-name" id="deck-name" required />
					<button type="submit">Upload</button>
				</form>
			</div>
		</div>
	</body>
</html>
`

const Cards = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<title>gocards</title>
		<script
			src="https://unpkg.com/htmx.org@2.0.3"
			integrity="sha384-0895/pl2MU10Hqc6jd4RvrthNlDiE9U1tWmX7WRESftEDRosgxNsQG/Ze9YMRzHq"
			crossorigin="anonymous"></script>
		<style>
			@import url("https://fonts.googleapis.com/css2?family=Noto+Sans+JP:wght@100..900&display=swap");

			body {
				padding: 0;
				margin: 0;
			}

			#ui {
				width: 100%;
				display: flex;
				flex-direction: column;
				justify-content: flex-start;
				align-items: center;
			}

			#word-ui {
				width: 100%;
				margin-bottom: 1rem;
				text-align: center;
				font-size: 4rem;
				font-family: "Noto Sans JP", sans-serif;
				color: white;
			}

			#box-ui {
				width: 95%;
			}

			#box-ui form {
				width: 100%;
				display: flex;
				justify-content: center;
			}

			#box-ui input {
				width: 100%;
				font-size: 2.5rem;
				text-align: center;
			}

			.neutral {
				background-color: teal;
			}

			.wrong {
				background-color: red;
			}
		</style>
	</head>
	<script>
		// For translation: romaji -> kana
		//
		// No katakana for now...
		const hiragana = new Map([
			// Main list
			["a", "あ"],
			["i", "い"],
			["u", "う"],
			["e", "え"],
			["o", "お"],
			["ka", "か"],
			["ki", "き"],
			["ku", "く"],
			["ke", "け"],
			["ko", "こ"],
			["sa", "さ"],
			["shi", "し"],
			["su", "す"],
			["se", "せ"],
			["so", "そ"],
			["ta", "た"],
			["chi", "ち"],
			["tsu", "つ"],
			["te", "て"],
			["to", "と"],
			["na", "な"],
			["ni", "に"],
			["nu", "ぬ"],
			["ne", "ね"],
			["no", "の"],
			["ha", "は"],
			["hi", "ひ"],
			["fu", "ふ"],
			["he", "へ"],
			["ho", "ほ"],
			["ma", "ま"],
			["mi", "み"],
			["mu", "む"],
			["me", "め"],
			["mo", "も"],
			["ra", "ら"],
			["ri", "り"],
			["ru", "る"],
			["re", "れ"],
			["ro", "ろ"],
			["ya", "や"],
			["yu", "ゆ"],
			["yo", "よ"],
			["wa", "わ"],
			["wo", "を"],
			["nn", "ん"],

			// Dakuten
			["ga", "が"],
			["gi", "ぎ"],
			["gu", "ぐ"],
			["ge", "げ"],
			["go", "ご"],
			["za", "ざ"],
			["ji", "じ"],
			["zu", "ず"],
			["ze", "ぜ"],
			["zo", "ぞ"],
			["da", "だ"],
			["dji", "ぢ"],
			["dzu", "づ"],
			["de", "で"],
			["do", "ど"],
			["ba", "ば"],
			["bi", "び"],
			["bu", "ぶ"],
			["be", "べ"],
			["bo", "ぼ"],
			["pa", "ぱ"],
			["pi", "ぴ"],
			["pu", "ぷ"],
			["pe", "ぺ"],
			["po", "ぽ"],

			// Combo
			["kya", "きゃ"],
			["kyu", "きゅ"],
			["kyo", "きょ"],
			["gya", "ぎゃ"],
			["gyu", "ぎゅ"],
			["gyo", "ぎょ"],
			["sha", "しゃ"],
			["shu", "しゅ"],
			["sho", "しょ"],
			["ja", "じゃ"],
			["ju", "じゅ"],
			["jo", "じょ"],
			["cha", "ちゃ"],
			["chu", "ちゅ"],
			["cho", "ちょ"],
			["dja", "ぢゃ"],
			["dju", "ぢゅ"],
			["djo", "ぢょ"],
			["nya", "にゃ"],
			["nyu", "にゅ"],
			["nyo", "にょ"],
			["hya", "ひゃ"],
			["hyu", "ひゅ"],
			["hyo", "ひょ"],
			["bya", "びゃ"],
			["byu", "びゅ"],
			["byo", "びょ"],
			["pya", "ぴゃ"],
			["pyu", "ぴゅ"],
			["pyo", "ぴょ"],
			["mya", "みゃ"],
			["myu", "みゅ"],
			["myo", "みょ"],
			["rya", "りゃ"],
			["ryu", "りゅ"],
			["ryo", "りょ"],

			// Misc
			["-", "ー"],
			[",", "、"],
			[".", "。"],
		]);

		// Save a global translation buffer
		//
		// This is really bad in terms of concurrency, but I think
		// since JS uses an event loop with a single thread, you
		// can get away with this kind of thing...
		//
		// Maybe a bug will prove me wrong later or something lol
		let buffer = "";

		// Replaces the text input with appropriate kana
		//
		// Since kana can span anywhere from 1-4 characters (sokukon),
		// we check these sizes from largest-to-smallest order. If it
		// matches anything in the dict, we replace it.
		function handleKey(event) {
			// Check for deletion
			//
			// We can delete from the actual form, but our content
			// will be forever bugged if we don't fix this now.
			if (event.inputType === "deleteContentBackward") {
				if (buffer.length != 0) buffer = buffer.slice(0, -1);
				return; // break!
			}

			// Save the last character
			//
			// Then we look for kana with our updated input;
			// noting that we can also have a repeated consonant
			// (sokukon) which will make the length of that "slice"
			// be succ(n), where n is the length without repeat.
			let input = event.target.value;
			buffer += input.slice(-1);
			if (hiragana.has(buffer)) {
				event.target.value =
					input.slice(0, input.length - buffer.length) + hiragana.get(buffer);
				buffer = ""; // reset
			} else {
				// Check for sokukon
				// e.g., たべちゃtta -> たべちゃった
				//
				// We need:
				// 	- a match for kana at the end, AND
				// 	- the character right _before_ the match needs to be congruent
				// 	  to the first character of the match
				let slice = buffer.slice(1); // shorthand
				if (hiragana.has(slice) && buffer[0] == slice[0]) {
					// Add xtsu
					event.target.value =
						input.slice(0, input.length - buffer.length) +
						"っ" +
						hiragana.get(slice);
					buffer = ""; // still counts as a find; reset
				}
			}
		}
	</script>
	<body>
		<div id="ui">
			<div id="word-ui" class="neutral">
				<p>{{.Word}}</p>
			</div>
			<div id="box-ui">
				<form
					hx-post="/cards/submit"
					hx-target="#ui"
					hx-swap="innerHTML"
					hx-trigger="submit">
					<input
						type="text"
						name="ans"
						id="ans"
						autocomplete="off"
						oninput="handleKey(event)" />
        		</form>
			</div>
		</div>
	</body>
</html>
`
