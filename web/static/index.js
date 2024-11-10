const editor = ace.edit("editor");
editor.setTheme("ace/theme/cloud_editor_dark");
editor.session.setMode("ace/mode/golang");

const languageSelector = document.querySelector("#language-select");
languageSelector.onchange = () => {
  editor.session.setMode("ace/mode/" + languageSelector.value);
}
