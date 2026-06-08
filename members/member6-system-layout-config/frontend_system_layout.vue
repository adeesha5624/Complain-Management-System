<!-- HTML Template part for App Container & Header Layout -->
<div class="app-container">
  <div class="background-shapes">
    <div class="shape shape-1"></div>
    <div class="shape shape-2"></div>
    <div class="shape shape-3"></div>
  </div>

  <main class="main-content">
    <header class="app-header">
      <div class="logo">
        <h2>CRM - Complain Management Module</h2>
      </div>
      <div class="header-right">
        <div class="user-info">
          <span class="role-badge" :class="{ 'role-admin': currentUser.role === 'admin' }">
            {{ currentUser.role?.toUpperCase() }}
          </span>
          <span class="username">{{ currentUser.username }}</span>
        </div>
        <button @click="logout" class="logout-btn" title="Logout">Logout</button>
      </div>
    </header>

    <p class="subtitle">Streamline your customer support experience</p>

    <div class="dashboard-grid">
      <!-- Injected Form Section from Member 3 -->
      <!-- Injected List Section from Member 5 -->
    </div>
  </main>
</div>

<!-- Global CSS (Sample extract) -->
<style>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap");

:root {
  --primary: #6366f1;
  --bg-color: #0f172a;
  --text-main: #f8fafc;
}

body {
  font-family: "Inter", sans-serif;
  background-color: var(--bg-color);
  color: var(--text-main);
  min-height: 100vh;
}

.glass-panel {
  background: var(--panel-bg);
  backdrop-filter: blur(12px);
  border-radius: 16px;
  padding: 1.5rem;
}
</style>
