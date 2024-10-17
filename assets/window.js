window.onload = function() {
    let htmlElement = document.documentElement;

    let mode = localStorage.getItem("mode");


    if (!mode) {
        htmlElement.classList.add("dark");
        localStorage.setItem("mode", "dark");
        mode = "dark";

        replaceIcons(mode)
    }

    mode === "dark" ? htmlElement.classList.add("dark") : htmlElement.classList.remove("dark");

    replaceIcons(mode)

}
