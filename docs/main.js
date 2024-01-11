const loadTheme = () => {
	const theme = localStorage.getItem('theme');
	if (theme) {
		document.querySelector("theme").href = `/themes/${theme}.css`;
	}
}

fetch("http://localhost:8001").then(t => t.text()).then(t => console.log(t))
