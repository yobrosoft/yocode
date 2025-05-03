// TypeScript implementation of Vim-like keybindings

// Define the possible editor modes
type EditorMode = 'normal' | 'insert' | 'visual';

// Editor interface
interface EditorElement extends HTMLTextAreaElement {
  selectionStart: number;
  selectionEnd: number;
  value: string;
}

class VimMode {
  private editor: EditorElement;
  private mode: EditorMode;
  private indicator: HTMLElement | null;

  constructor(editor: EditorElement) {
    this.editor = editor;
    this.mode = 'normal'; // normal, insert, visual
    this.indicator = document.querySelector('.vim-mode-indicator');
    this.setupKeyHandlers();
  }

  private setupKeyHandlers(): void {
    this.editor.addEventListener('keydown', (e: KeyboardEvent) => {
      if (this.mode === 'normal') {
        this.handleNormalMode(e);
      } else if (this.mode === 'insert') {
        this.handleInsertMode(e);
      }
    });
  }

  private handleNormalMode(e: KeyboardEvent): void {
    // Handle normal mode key bindings
    switch (e.key) {
      case 'i':
        this.setMode('insert');
        e.preventDefault();
        break;
      case 'h':
        // Move cursor left
        const pos = this.editor.selectionStart;
        if (pos > 0) {
          this.editor.selectionStart = this.editor.selectionEnd = pos - 1;
        }
        e.preventDefault();
        break;
      case 'l':
        // Move cursor right
        const posRight = this.editor.selectionStart;
        if (posRight < this.editor.value.length) {
          this.editor.selectionStart = this.editor.selectionEnd = posRight + 1;
        }
        e.preventDefault();
        break;
      case 'j':
        // Move cursor down (simplified)
        this.moveCursorVertically(1);
        e.preventDefault();
        break;
      case 'k':
        // Move cursor up (simplified)
        this.moveCursorVertically(-1);
        e.preventDefault();
        break;
    }
  }

  private handleInsertMode(e: KeyboardEvent): void {
    // Handle insert mode key bindings
    if (e.key === 'Escape') {
      this.setMode('normal');
      e.preventDefault();
    }
  }

  private setMode(mode: EditorMode): void {
    this.mode = mode;
    if (this.indicator) {
      this.indicator.textContent = mode.toUpperCase();
    }
  }

  private moveCursorVertically(direction: number): void {
    // This is a simplified implementation
    // A real implementation would need to calculate line positions
    const text = this.editor.value;
    const pos = this.editor.selectionStart;
    
    // Find the current line start and end
    let lineStart = text.lastIndexOf('\n', pos - 1) + 1;
    let lineEnd = text.indexOf('\n', pos);
    if (lineEnd === -1) lineEnd = text.length;
    
    // Calculate column position
    const column = pos - lineStart;
    
    if (direction > 0) {
      // Move down
      if (lineEnd < text.length) {
        const nextLineStart = lineEnd + 1;
        const nextLineEnd = text.indexOf('\n', nextLineStart);
        if (nextLineEnd === -1) {
          const nextLineLength = text.length - nextLineStart;
          const newColumn = Math.min(column, nextLineLength);
          this.editor.selectionStart = this.editor.selectionEnd = nextLineStart + newColumn;
        } else {
          const nextLineLength = nextLineEnd - nextLineStart;
          const newColumn = Math.min(column, nextLineLength);
          this.editor.selectionStart = this.editor.selectionEnd = nextLineStart + newColumn;
        }
      }
    } else {
      // Move up
      if (lineStart > 0) {
        const prevLineEnd = lineStart - 1;
        const prevLineStart = text.lastIndexOf('\n', prevLineEnd - 1) + 1;
        
        const prevLineLength = prevLineEnd - prevLineStart;
        const newColumn = Math.min(column, prevLineLength);
        
        this.editor.selectionStart = this.editor.selectionEnd = prevLineStart + newColumn;
      }
    }
  }
}

// Initialize Vim mode when the editor is ready
function initVimMode(): void {
  const editorElement = document.querySelector('#editor textarea') as EditorElement;
  if (editorElement) {
    new VimMode(editorElement);
  } else {
    // If editor isn't ready yet, wait a bit and try again
    setTimeout(initVimMode, 100);
  }
}

// Start initialization when the page loads
document.addEventListener('DOMContentLoaded', (): void => {
  setTimeout(initVimMode, 500); // Give the main.ts time to create the textarea
});
