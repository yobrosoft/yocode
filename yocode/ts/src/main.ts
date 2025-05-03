// Main TypeScript file for the editor UI

// Editor interface
interface EditorElement extends HTMLTextAreaElement {
  selectionStart: number;
  selectionEnd: number;
  value: string;
}

// Initialize the editor when the DOM is fully loaded
document.addEventListener('DOMContentLoaded', (): void => {
  initEditor();
});

function initEditor(): void {
  // Create a simple editor with Vim-like keybindings
  console.log('Editor initialized');
  
  // Add a simple text area for now
  const editorContainer: HTMLElement | null = document.getElementById('editor');
  if (editorContainer) {
    const textarea: HTMLTextAreaElement = document.createElement('textarea');
    textarea.style.width = '100%';
    textarea.style.height = '100%';
    textarea.style.fontFamily = 'monospace';
    textarea.style.fontSize = '14px';
    textarea.style.padding = '10px';
    textarea.style.border = 'none';
    textarea.style.outline = 'none';
    textarea.style.resize = 'none';
    textarea.placeholder = 'Type here...';
    
    editorContainer.appendChild(textarea);
    
    // Focus the textarea
    textarea.focus();
  }
}
