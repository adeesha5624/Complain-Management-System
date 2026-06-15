<style>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap");

:root {
  --primary: #6366f1;
  --primary-hover: #4f46e5;
  --bg-color: #0f172a;
  --panel-bg: rgba(30, 41, 59, 0.7);
  --panel-border: rgba(255, 255, 255, 0.1);
  --text-main: #f8fafc;
  --text-muted: #94a3b8;
  --success: #10b981;
  --warning: #f59e0b;
  --info: #3b82f6;
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: "Inter", sans-serif;
  background-color: var(--bg-color);
  color: var(--text-main);
  min-height: 100vh;
  overflow-x: hidden;
}

.app-container {
  position: relative;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

/* Dynamic Background Elements */
.background-shapes {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  z-index: -1;
  pointer-events: none;
}
.shape {
  position: absolute;
  filter: blur(80px);
  opacity: 0.5;
  border-radius: 50%;
  animation: float 20s infinite ease-in-out;
}
.shape-1 {
  background: #3b82f6;
  width: 400px;
  height: 400px;
  top: -100px;
  left: -100px;
}
.shape-2 {
  background: #8b5cf6;
  width: 300px;
  height: 300px;
  bottom: -50px;
  right: -50px;
  animation-delay: -5s;
}
.shape-3 {
  background: #06b6d4;
  width: 250px;
  height: 250px;
  top: 40%;
  left: 50%;
  animation-delay: -10s;
}

@keyframes float {
  0%,
  100% {
    transform: translate(0, 0) scale(1);
  }
  33% {
    transform: translate(30px, -50px) scale(1.1);
  }
  66% {
    transform: translate(-20px, 20px) scale(0.9);
  }
}

.main-content {
  width: 100%;
  max-width: 1200px;
  z-index: 1;
  padding: 2rem;
}

.app-header {
  text-align: center;
  margin-bottom: 3rem;
  animation: slideDown 0.8s cubic-bezier(0.16, 1, 0.3, 1);
}
.app-header .logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-bottom: 0.5rem;
}
.app-header svg {
  width: 32px;
  height: 32px;
  color: var(--primary);
}
.app-header h2 {
  font-size: 2.2rem;
  font-weight: 700;
  letter-spacing: -0.5px;
  background: linear-gradient(to right, #fff, #94a3b8);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
.subtitle {
  color: var(--text-muted);
  font-size: 1.1rem;
}

.dashboard-grid {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 2rem;
}

@media (max-width: 900px) {
  .dashboard-grid {
    grid-template-columns: 1fr;
  }
}

/* Glassmorphism Panels */
.glass-panel {
  background: var(--panel-bg);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border: 1px solid var(--panel-border);
  border-radius: 16px;
  padding: 1.5rem;
  box-shadow: 0 20px 40px -15px rgba(0, 0, 0, 0.3);
  animation: fadeIn 1s cubic-bezier(0.16, 1, 0.3, 1);
  transition:
    transform 0.3s ease,
    box-shadow 0.3s ease;
}
.glass-panel:hover {
  transform: translateY(-2px);
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.4);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid var(--panel-border);
}
.section-header h3 {
  font-size: 1.25rem;
  font-weight: 600;
}
.section-header .icon svg,
.refresh-btn svg {
  width: 20px;
  height: 20px;
  color: var(--text-muted);
}
.refresh-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  border-radius: 8px;
  transition: all 0.2s;
}
.refresh-btn:hover {
  background: rgba(255, 255, 255, 0.05);
  color: #fff;
}
.spin {
  animation: spin 1s linear infinite;
}

/* Form Styles */
.form-group {
  margin-bottom: 1.25rem;
}
.form-group label {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  margin-bottom: 0.5rem;
  color: var(--text-muted);
}
.input-wrapper input,
.input-wrapper textarea {
  width: 100%;
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid var(--panel-border);
  color: #fff;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  font-family: inherit;
  font-size: 0.95rem;
  transition: all 0.2s;
}
.input-wrapper input:focus,
.input-wrapper textarea:focus {
  outline: none;
  border-color: var(--primary);
  box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
}
.input-wrapper textarea {
  min-height: 100px;
  resize: vertical;
}

.submit-btn {
  width: 100%;
  background: var(--primary);
  color: white;
  border: none;
  padding: 0.875rem;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition:
    background 0.2s,
    transform 0.1s;
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 48px;
}
.submit-btn:hover:not(:disabled) {
  background: var(--primary-hover);
}
.submit-btn:active:not(:disabled) {
  transform: scale(0.98);
}
.submit-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

/* Table Styles */
.table-container {
  overflow-x: auto;
}
table {
  width: 100%;
  border-collapse: collapse;
}
th {
  text-align: left;
  padding: 1rem;
  font-size: 0.875rem;
  color: var(--text-muted);
  font-weight: 500;
  border-bottom: 1px solid var(--panel-border);
}
td {
  padding: 1rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  vertical-align: top;
}
.table-row {
  transition: background 0.2s;
}
.table-row:hover {
  background: rgba(255, 255, 255, 0.03);
}
.id-col {
  color: var(--text-muted);
  font-family: monospace;
}
.title-cell {
  font-weight: 500;
  margin-bottom: 0.25rem;
}
.desc-cell {
  font-size: 0.875rem;
  color: var(--text-muted);
  line-height: 1.4;
</style>
