<!-- HTML Template part for Admin Actions & Modals -->
<td v-if="currentUser.role === 'admin'" class="actions-cell">
  <button class="action-btn edit-btn" @click="openStatusModal(item)" title="Change Status">Edit Status</button>
  <button class="action-btn delete-btn" @click="confirmDelete(item)" title="Delete Complaint">Delete</button>
</td>

<!-- Status Modal -->
<div v-if="showStatusModal" class="modal-overlay" @click.self="closeStatusModal">
<div class="modal-content">
    <div class="modal-header">
    <h4>Change Complaint Status</h4>
    <button class="close-btn" @click="closeStatusModal">×</button>
    </div>
    <div class="modal-body">
    <p class="complaint-info">Complaint #{{ selectedComplaint?.id }} - {{ selectedComplaint?.title }}</p>
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
    <button class="btn-secondary" @click="closeStatusModal">Cancel</button>
    <button class="btn-primary" @click="updateStatus">Update Status</button>
    </div>
</div>
</div>

<!-- Delete Confirmation Modal -->
<div v-if="showDeleteModal" class="modal-overlay" @click.self="closeDeleteModal">
<div class="modal-content danger">
    <div class="modal-header">
    <h4>Delete Complaint</h4>
    <button class="close-btn" @click="closeDeleteModal">×</button>
    </div>
    <div class="modal-body">
    <p class="warning-text">Are you sure you want to delete this complaint?</p>
    <p class="complaint-info">Complaint #{{ selectedComplaint?.id }} - {{ selectedComplaint?.title }}</p>
    <p class="warning-subtext">This action cannot be undone.</p>
    </div>
    <div class="modal-footer">
    <button class="btn-secondary" @click="closeDeleteModal">Cancel</button>
    <button class="btn-danger" @click="deleteComplaint">Delete</button>
    </div>
</div>
</div>

<!-- Script part for Admin Manage Complaint -->
<script>
export default {
  data() {
    return {
      showStatusModal: false,
      showDeleteModal: false,
      selectedComplaint: null,
      newStatus: "Pending",
      apiBase: "http://localhost:8080/api/complaints",
    };
  },
  methods: {
    openStatusModal(complaint) {
      this.selectedComplaint = complaint;
      this.newStatus = complaint.status;
      this.showStatusModal = true;
    },
    closeStatusModal() {
      this.showStatusModal = false;
      this.selectedComplaint = null;
      this.newStatus = "Pending";
    },
    async updateStatus() {
      if (!this.selectedComplaint) return;
      try {
        const res = await fetch(`${this.apiBase}?id=${this.selectedComplaint.id}`, {
          method: "PUT",
          headers: this.getAuthHeaders(),
          body: JSON.stringify({ status: this.newStatus }),
        });
        if (res.ok) {
          this.closeStatusModal();
          await this.fetchComplaints(); // From Member 5
        } else if (res.status === 401) {
          alert("Unauthorized - Admin access required");
        }
      } catch (err) {
        console.error("Error updating status:", err);
      }
    },
    confirmDelete(complaint) {
      this.selectedComplaint = complaint;
      this.showDeleteModal = true;
    },
    closeDeleteModal() {
      this.showDeleteModal = false;
      this.selectedComplaint = null;
    },
    async deleteComplaint() {
      if (!this.selectedComplaint) return;
      try {
        const res = await fetch(`${this.apiBase}?id=${this.selectedComplaint.id}`, {
          method: "DELETE",
          headers: this.getAuthHeaders(),
        });
        if (res.ok) {
          this.closeDeleteModal();
          await this.fetchComplaints(); // From Member 5
        } else if (res.status === 401) {
          alert("Unauthorized - Admin access required");
        }
      } catch (err) {
        console.error("Error deleting complaint:", err);
      }
    }
  }
}
</script>
