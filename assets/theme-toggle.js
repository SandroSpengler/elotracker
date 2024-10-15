function themeToggle() {
	let htmlElement = document.documentElement;
	let currentTheme = localStorage.getItem("mode");

	let newTheme = "dark";

	if (currentTheme === "dark") {
		newTheme = "light";
		htmlElement.classList.remove("dark");

		replaceIcons(newTheme);
	}

	if (currentTheme === "light") {
		newTheme === "dark";
		htmlElement.classList.add("dark");

		replaceIcons(newTheme);
	}

	localStorage.setItem("mode", newTheme);
}

/**
 * close an open modal dialog
 * @param {string} theme
 */
function replaceIcons(theme) {
	if (!theme) {
		throw new Error("theme is required");
	}

	const socialIconHTML = document.querySelectorAll("#socialIcon");

	if (socialIconHTML && socialIconHTML.length > 0) {
		if (theme === "dark") {
			for (const icon of socialIconHTML) {
				icon.style.filter = "invert(0.75)";
			}
		} else {
			for (const icon of socialIconHTML) {
				icon.style.filter = "";
			}
		}
	}
}
