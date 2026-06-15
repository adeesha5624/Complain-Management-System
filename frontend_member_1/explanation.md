# Frontend Member 1: UI Layout & Components

## Responsibilities
Your part in the viva is to explain the visual layout, HTML structure, and styling of the application. You will focus on the `<template>` and `<style>` sections of the `App.vue` file.

## Key Code Parts to Explain

### 1. App Structure & Conditional Rendering
- **Main Container:** The app is wrapped in `.app-container` and uses a `.main-content` area.
- **Conditional Views:** Explain how `<header v-if="currentUser.token">` and `<main v-if="currentUser.token">` ensure that the app content is only visible if the user is authenticated. 
- **Role-based Rendering:** Explain how the layout changes based on the role. For example: `v-if="currentUser.role !== 'admin'"` hides the "Log New Complaint" form for admins, as they only manage complaints. 

### 2. Form and Table UI
- **Form UI:** Explain the `v-model` bindings in the complaint form (Email, Title, Description) which automatically sync user input to the Vue data object.
- **Table UI:** Explain the `v-for="item in complaints"` loop which iterates over the complaints array to dynamically render rows in the table. 
- **Action Buttons:** Point out the buttons in the table (Edit, Delete) and how they also conditionally render based on the user's role (`v-if="currentUser.role === 'admin'"`).

### 3. Modals and Interactions
- Explain the HTML structure for the **Status Modal** and **Delete Confirmation Modal**. 
- Highlight the `@click.self="closeStatusModal"` directive which allows closing the modal by clicking outside the modal content.

### 4. CSS and Theming
- **Glassmorphism:** Explain the `.glass-panel` CSS class which uses `backdrop-filter: blur(12px)` and semi-transparent backgrounds to create a modern "glass" effect.
- **Dynamic Background:** Explain the `.background-shapes` and `.shape` classes which use CSS animations (`@keyframes float`) to animate background elements.
- **CSS Variables:** Mention the use of `:root` variables for consistent theming (colors like `--primary`, `--bg-color`).
