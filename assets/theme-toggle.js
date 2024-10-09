function themeToggle() {
	let htmlElement = document.documentElement;
	let currentTheme = localStorage.getItem("mode");

	let newTheme = "dark";

	if (currentTheme === "dark") {
		newTheme = "light";
		htmlElement.classList.remove("dark");
	}

	if (currentTheme === "light") {
		newTheme === "dark";
		htmlElement.classList.add("dark");
	}

	localStorage.setItem("mode", newTheme);
}
