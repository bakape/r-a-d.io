// Landing page JS

// Main R/a/dio API data structure
type API = {
	np: string
	listeners: number
	current: number
	start_time: number
	end_time: number
	thread: string
	requesting: true
	dj: {
		id: number
		djname: string
		djimage: number
	}
	queue: Song[]
	lp: Song[]
}

// Song-related data
type Song = {
	meta: string
	timestamp: number
}

(() => {
	// Cache elements to modify
	const nowPlaying = document.getElementById("now-playing");
	const listeners = document.getElementById("listeners");
	const progress = {
		time: document.getElementById("time-progress"),
		bar: document.getElementById("progress-bar") as HTMLProgressElement,
	};
	const dj = {
		image: document.getElementById("dj-image") as HTMLImageElement,
		name: document.getElementById("dj-name"),
	};
	const lastPlayed = document.getElementById("last-played");
	const queue = document.getElementById("queue");

	// Last fetched API data
	// TODO: Initial data should be extracted from prerendered page as JSON
	let data: API = null;

	// So we can make the next fetch happen quicker, when the last song ends
	let nextFetch = 0;

	// Correction for time skew between API server and client
	let skew = 0;

	// Reduce default volume
	(document.getElementById("stream") as HTMLAudioElement).volume = 0.2;

	fetchData();
	setInterval(renderProgress, 1000);

	function fetchData() {
		const xhr = new XMLHttpRequest()
		xhr.onload = () => {
			if (xhr.status === 200) {
				data = xhr.response.main;
				skew = now() - data.current;
				render();
			}
			nextFetch = setTimeout(fetchData, 10000);
		}
		xhr.onerror = () =>
			nextFetch = setTimeout(fetchData, 10000);
		xhr.responseType = "json";
		xhr.open("GET", "https://r-a-d.io/api");
		xhr.send();
	}

	function render() {
		nowPlaying.textContent = data.np;
		listeners.textContent = `Listeners: ${data.listeners}`;
		dj.image.src = "https://r-a-d.io/api/dj-image/" + data.dj.djimage;
		dj.name.textContent = data.dj.djname;
		renderProgress();

		let html = "";
		for (let { timestamp, meta } of data.lp) {
			html += `<tr>`
				+ `<td>${renderTime(timestamp)}</td><td>${meta}</td>`
				+ `</tr>`;
		}
		lastPlayed.innerHTML = html;

		html = "";
		for (let { timestamp, meta } of data.queue) {
			html += `<tr>`
				+ `<td>${meta}</td><td>${renderTime(timestamp)}</td>`
				+ `</tr>`;
		}
		queue.innerHTML = html;
	}

	function renderProgress() {
		// Not yet fetched
		if (!data) {
			return;
		}

		const delta = data.end_time - data.start_time;
		const prog = now() - skew - data.start_time;
		progress.time.textContent
			= `${formatDuration(prog)} / ${formatDuration(delta)}`;
		progress.bar.value = delta ? prog / delta : 0;
		if (prog > delta) { // Data is stale
			clearTimeout(nextFetch);
			nextFetch = 0;
			fetchData();
		}
	}

	// Return current Unix time
	function now(): number {
		return Math.floor(Date.now() / 1000);
	}

	// Pad with zero, if needed
	function pad(n: number): string {
		return n < 10 ? `0${n}` : n.toString();
	}

	// Format duration string, such as "03:55", from second number
	function formatDuration(n: number): string {
		return `${pad(Math.floor(n / 60))}:${pad(n % 60)}`;
	}

	// Renders readable elapsed/remaining time
	function renderTime(t: number): string {
		t = Math.floor(now() - t);
		let isFuture = false;
		if (t < 1) {
			isFuture = true;
			t = -t;
		}

		if (t < 60) {
			return isFuture ? "soon™" : "just now";
		}
		t = Math.floor(t / 60);
		if (t < 60) {
			return ago(t, "minute", isFuture);
		}
		return ago(Math.floor(t / 60), "hour", isFuture);
	}

	// Return either the singular or plural form of a word, depending on n
	function pluralize(n: number, word: string): string {
		let s = `${n} ${word}`;
		if (n !== 1 && n !== -1) {
			s += "s";
		}
		return s;
	}

	// Renders "56 minutes ago" or "in 56 minutes" like relative time text
	function ago(time: number, word: string, isFuture: boolean, ): string {
		const count = pluralize(time, word)
		if (isFuture) {
			return `in ${count}`
		}
		return `${count} ago`
	}
})();
