const editor = ace.edit("editor");
editor.setTheme("ace/theme/github_dark");

const languageSelector = document.querySelector("#language-select");

for (let lang of Object.keys(aceLangs)) {
  languageSelector.innerHTML += `<option value="${aceLangs[lang].mode}">${aceLangs[lang].name}</option>`
}

languageSelector.onchange = () => {
  editor.session.setMode(languageSelector.value);
}
