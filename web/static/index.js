import fnv1a from "https://cdn.jsdelivr.net/npm/@sindresorhus/fnv1a@3.1.0/index.min.js";

const language = document.querySelector("#hidden").innerText;

const editor = ace.edit("editor");
editor.setTheme("ace/theme/github_dark");
editor.session.setMode("ace/mode/"+language);

const languageSelector = document.querySelector("#language-select");

for (let lang of aceLangs) {
  languageSelector.innerHTML += `<option value="${lang}">${lang}</option>`
}

languageSelector.onchange = () => {
  editor.session.setMode("ace/mode/"+languageSelector.value);
}

languageSelector.selectedIndex = aceLangs.indexOf(language);

const saveButton = document.querySelector("#save-btn");
saveButton.onclick = () => {
  const content = editor.getValue();
  const valueToHash = content + Date.now().toString();

  fetch("/save", {
    method: "POST",
    body: JSON.stringify(
      {
        id: fnv1a(valueToHash, {size: 64}).toString(),
        content: content,
        language: languageSelector.value,
      },
    ),
    headers: {
      "Content-Type": "application/json",
    },
  }).then(res => console.log(res.json()));
}
