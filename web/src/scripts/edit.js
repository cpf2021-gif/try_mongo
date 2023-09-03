const simplemde = new SimpleMDE({
    autoDownloadFontAwesome: true,
    autofocus: true,
    toolbar: ["preview"],
    shortcuts: {
        "togglePreview": "Alt-P",
    }
});

// Clear the editor
const button = document.querySelector('.alert');
    button.addEventListener('click', () => {
    simplemde.value("")
});

// Add markdown-body class to preview
simplemde.codemirror.on("keyHandled", function (editor, name, event) {
    if (name === "Alt-P") {
        const preview = document.querySelector('.editor-preview');
        if (preview) {
            preview.classList.add('markdown-body')
        }
    }
})
