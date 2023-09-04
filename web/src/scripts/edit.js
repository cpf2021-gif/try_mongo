import EasyMDE from "easymde";
import "highlight.js/styles/github.css";


const simplemde = new EasyMDE({
    autoDownloadFontAwesome: true,
    maxHeight: "600px",
    toolbar: ["preview"],
    shortcuts: {
        "togglePreview": "Alt-P",
        "toggleFullScreen": "Alt-O",
    },
    renderingConfig: {
        codeSyntaxHighlighting: true,
    }
});

// Load the editor content from localStorage
const content = localStorage.getItem('md-content');
if (content) {
    simplemde.value(content);
}

// Load the editor title from localStorage
const title = localStorage.getItem('md-title');
const titleInput = document.querySelector('.title-input');
if (title) {
    titleInput.value = title;
}

// Clear the editor
const button = document.querySelector('.alert');
    button.addEventListener('click', () => {
    simplemde.value("")
});

// Save the editor content
const saveButton = document.querySelector('.save');
saveButton.addEventListener('click', () => {
    // Save the content to localStorage
    const content = simplemde.value();
    localStorage.setItem('md-content', content);
    // Save the title to localStorage
    const title = titleInput.value;
    localStorage.setItem('md-title', title); 
});

// Create a new file
const createButton = document.querySelector('.create');
createButton.addEventListener('click', () => {
    console.log('create');
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
