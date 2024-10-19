window.onload = function () {
	let mode = loadTheme();
	replaceIcons(mode);

	document.body.addEventListener("htmx:afterSwap", () => {
		// idk maybe wrong event but it triggers too early
		setTimeout(() => {
			replaceIcons(mode);
		}, 100);
	});
};

function loadTheme(mode) {
	let mode_storage = localStorage.getItem("mode");
	let htmlElement = document.documentElement;

	if (!mode) {
		htmlElement.classList.add("dark");
		localStorage.setItem("mode", "dark");
		mode = "dark";

		replaceIcons(mode_storage);
	}

	mode === "dark" ? htmlElement.classList.add("dark") : htmlElement.classList.remove("dark");

	return mode;
}
