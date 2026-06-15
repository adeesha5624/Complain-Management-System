<template>
  <div class="app-container">
    <div class="background-shapes">
      <div class="shape shape-1"></div>
      <div class="shape shape-2"></div>
      <div class="shape shape-3"></div>
    </div>

    <!-- Main App View -->
    <header v-if="currentUser.token" class="app-header-unified">
      <div class="unified-header-left">
        <div class="unified-logo">◈</div>
        <div class="unified-brand">
          <p class="unified-subtitle">CRM ENTERPRISE</p>
          <p class="unified-title">Complain Management</p>
        </div>
      </div>
      
      <div class="unified-header-center">
      </div>

      <div class="unified-header-right">
        <button @click="goToPortal" class="unified-btn-portal" title="Back to Portal">
          🏠 Portal
        </button>
        <div class="unified-user-info">
          <span class="unified-role">{{ currentUser.role || 'USER' }}</span>
          <span class="unified-username">{{ currentUser.username || 'Agent' }}</span>
        </div>
        <button @click="logout" class="unified-btn-logout">
          Logout
        </button>
      </div>
    </header>

    <main v-if="currentUser.token" class="main-content">
      <p class="subtitle">Streamline your customer support experience</p>

      <div class="dashboard-grid" :style="currentUser.role === 'admin' ? 'grid-template-columns: 1fr;' : ''">
        <section class="glass-panel form-section" v-if="currentUser.role !== 'admin'">
          <div class="section-header">
            <h3>Log New Complaint</h3>
            <span class="icon">
              <svg
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"
                ></path>
                <path
                  d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"
                ></path>
              </svg>
            </span>
          </div>

          <form @submit.prevent="submitComplaint" class="complaint-form">
            <div class="form-group">
              <label>Email Address</label>
              <div class="input-wrapper">
                <input
                  type="email"
                  v-model="form.customer_id"
                  required
                  placeholder="e.g. user@example.com"
                />
              </div>
            </div>

            <div class="form-group">
              <label>Complaint Title</label>
              <div class="input-wrapper">
                <input
                  type="text"
                  v-model="form.title"
                  required
                  placeholder="Brief summary of the issue"
                />
              </div>
            </div>

            <div class="form-group">
              <label>Description Details</label>
              <div class="input-wrapper">
                <textarea
                  v-model="form.description"
                  required
                  placeholder="Provide full details here..."
                ></textarea>
              </div>
            </div>

            <button type="submit" class="submit-btn" :disabled="isSubmitting">
              <span v-if="!isSubmitting">Submit Complaint</span>
              <span v-else class="loader"></span>
            </button>
          </form>
        </section>

        <section class="glass-panel list-section">
          <div class="section-header">
            <h3>Active System Complaints</h3>
            <button
              @click="fetchComplaints"
              class="refresh-btn"
              title="Refresh"
            >
              <svg
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                :class="{ spin: isFetching }"
              >
                <polyline points="23 4 23 10 17 10"></polyline>
                <path d="M20.49 15a9 9 0 1 1-2.12-9.36L23 10"></path>
              </svg>
            </button>
          </div>

          <div class="table-container">
            <table v-if="complaints.length > 0">
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Email Address</th>
                  <th>Title</th>
                  <th>Status</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in complaints" :key="item.id" class="table-row">
                  <td class="id-col">#{{ item.id }}</td>
                  <td>
                    <strong>{{ item.customer_id }}</strong>
                  </td>
                  <td>
                    <div class="title-cell">{{ item.title }}</div>
                    <div class="desc-cell">{{ item.description }}</div>
                  </td>
                  <td>
                    <span
                      class="status-badge"
                      :class="statusClass(item.status)"
                    >
                      {{ item.status }}
                    </span>
                  </td>
                  <td class="actions-cell">
                    <button
                      v-if="currentUser.role === 'admin'"
                      class="action-btn edit-btn"
                      @click="openStatusModal(item)"
                      title="Change Status"
                    >
                      <svg
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                      >
                        <polyline
                          points="12 3 20 7.5 20 16.5 12 21 4 16.5 4 7.5 12 3"
                        ></polyline>
                        <line x1="12" y1="12" x2="20" y2="7.5"></line>
                        <line x1="12" y1="12" x2="12" y2="21"></line>
                        <line x1="12" y1="12" x2="4" y2="7.5"></line>
                      </svg>
                    </button>
                    <button
                      v-if="currentUser.role === 'admin' || currentUser.email === item.customer_id"
                      class="action-btn delete-btn"
                      @click="confirmDelete(item)"
                      title="Delete Complaint"
                    >
                      <svg
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        stroke-width="2"
                      >
                        <polyline points="3 6 5 6 21 6"></polyline>
                        <path
                          d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"
                        ></path>
                        <line x1="10" y1="11" x2="10" y2="17"></line>
                        <line x1="14" y1="11" x2="14" y2="17"></line>
                      </svg>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>

            <div v-else class="empty-state">
              <svg
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="1.5"
              >
                <circle cx="12" cy="12" r="10"></circle>
                <line x1="12" y1="8" x2="12" y2="12"></line>
                <line x1="12" y1="16" x2="12.01" y2="16"></line>
              </svg>
              <p>No complaints found.</p>
            </div>
          </div>
        </section>
      </div>

      <!-- Status Modal -->
      <div
        v-if="showStatusModal"
        class="modal-overlay"
        @click.self="closeStatusModal"
      >
        <div class="modal-content">
          <div class="modal-header">
            <h4>Change Complaint Status</h4>
            <button class="close-btn" @click="closeStatusModal">×</button>
          </div>
          <div class="modal-body">
            <p class="complaint-info">
              Complaint #{{ selectedComplaint?.id }} -
              {{ selectedComplaint?.title }}
            </p>
            <div class="form-group">
              <label>New Status</label>
              <select v-model="newStatus" class="status-select">
                <option value="Pending">Pending</option>
                <option value="In Progress">In Progress</option>
                <option value="Resolved">Resolved</option>
              </select>
            </div>
          </div>
          <div class="modal-footer">
            <button class="btn-secondary" @click="closeStatusModal">
              Cancel
            </button>
            <button class="btn-primary" @click="updateStatus">
              Update Status
            </button>
          </div>
        </div>
      </div>

      <!-- Delete Confirmation Modal -->
      <div
        v-if="showDeleteModal"
        class="modal-overlay"
        @click.self="closeDeleteModal"
      >
        <div class="modal-content danger">
          <div class="modal-header">
            <h4>Delete Complaint</h4>
            <button class="close-btn" @click="closeDeleteModal">×</button>
          </div>
          <div class="modal-body">
            <p class="warning-text">
              Are you sure you want to delete this complaint?
            </p>
            <p class="complaint-info">
              Complaint #{{ selectedComplaint?.id }} -
              {{ selectedComplaint?.title }}
            </p>
            <p class="warning-subtext">This action cannot be undone.</p>
          </div>
          <div class="modal-footer">
            <button class="btn-secondary" @click="closeDeleteModal">
              Cancel
            </button>
            <button class="btn-danger" @click="deleteComplaint">Delete</button>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
