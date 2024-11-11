import fnv1a from "https://cdn.jsdelivr.net/npm/@sindresorhus/fnv1a@3.1.0/index.min.js";

const editor = ace.edit("editor");
editor.setTheme("ace/theme/github_dark");

const languageSelector = document.querySelector("#language-select");

for (let lang of Object.keys(aceLangs)) {
  languageSelector.innerHTML += `<option value="${aceLangs[lang].mode}">${aceLangs[lang].name}</option>`
}

languageSelector.onchange = () => {
  editor.session.setMode(languageSelector.value);
}

const saveButton = document.querySelector("#save-btn");
saveButton.onclick = () => {
  const content = document.querySelector("#editor").innerText;
  const valueToHash = content + Date.now().toString();

  fetch("/save", {
    method: "POST",
    body: JSON.stringify(
      {
        id: fnv1a(valueToHash, {size: 64}).toString(),
        content: content,
      },
    ),
    headers: {
      "Content-Type": "application/json",
    },
  }).then(res => console.log(res.json()));
}
