<!-- HTML Template part for User Add Complaint -->
<section v-if="currentUser.role !== 'admin'" class="glass-panel form-section">
<div class="section-header">
    <h3>Log New Complaint</h3>
</div>

<form @submit.prevent="submitComplaint" class="complaint-form">
    <div class="form-group">
    <label>Customer ID (Shared Link)</label>
    <div class="input-wrapper">
        <input type="number" v-model.number="form.customer_id" required placeholder="e.g. 101" />
    </div>
    </div>

    <div class="form-group">
    <label>Complaint Title</label>
    <div class="input-wrapper">
        <input type="text" v-model="form.title" required placeholder="Brief summary of the issue" />
    </div>
    </div>

    <div class="form-group">
    <label>Description Details</label>
    <div class="input-wrapper">
        <textarea v-model="form.description" required placeholder="Provide full details here..."></textarea>
    </div>
    </div>

    <button type="submit" class="submit-btn" :disabled="isSubmitting">
    <span v-if="!isSubmitting">Submit Complaint</span>
    <span v-else class="loader"></span>
    </button>
</form>
</section>

<!-- Script part for Add Complaint -->
<script>
export default {
  data() {
    return {
      form: { customer_id: null, title: "", description: "" },
      isSubmitting: false,
      apiBase: "http://localhost:8080/api/complaints",
    };
  },
  methods: {
    async submitComplaint() {
      if (this.isSubmitting) return;
      this.isSubmitting = true;
      try {
        const res = await fetch(this.apiBase, {
          method: "POST",
          headers: this.getAuthHeaders(),
          body: JSON.stringify(this.form),
        });
        if (res.ok) {
          this.form = { customer_id: null, title: "", description: "" };
          await this.fetchComplaints(); // From member 5
        } else if (res.status === 401) {
          alert("Unauthorized - please log in again.");
          this.logout();
        }
      } catch (err) {
        console.error("Error posting complaint:", err);
      } finally {
        this.isSubmitting = false;
      }
    }
  }
}
</script>
