<!-- HTML Template part for View Complaints List -->
<section :class="['glass-panel', 'list-section', { 'full-width': currentUser.role === 'admin' }]">
<div class="section-header">
    <h3>Active System Complaints</h3>
    <button @click="fetchComplaints" class="refresh-btn" title="Refresh">
    Refresh
    </button>
</div>

<div class="table-container">
    <table v-if="complaints.length > 0">
    <thead>
        <tr>
        <th>ID</th>
        <th>Customer ID</th>
        <th>Title</th>
        <th>Status</th>
        <th v-if="currentUser.role === 'admin'">Actions</th>
        </tr>
    </thead>
    <tbody>
        <tr v-for="item in complaints" :key="item.id" class="table-row">
        <td class="id-col">#{{ item.id }}</td>
        <td><strong>{{ item.customer_id }}</strong></td>
        <td>
            <div class="title-cell">{{ item.title }}</div>
            <div class="desc-cell">{{ item.description }}</div>
        </td>
        <td>
            <span class="status-badge" :class="statusClass(item.status)">{{ item.status }}</span>
        </td>
        <!-- Actions cell implemented by Member 4 -->
        </tr>
    </tbody>
    </table>

    <div v-else class="empty-state">
    <p>No complaints found.</p>
    </div>
</div>
</section>

<!-- Script part for View Complaints -->
<script>
export default {
  data() {
    return {
      complaints: [],
      isFetching: false,
      apiBase: "http://localhost:8080/api/complaints",
    };
  },
  methods: {
    async fetchComplaints() {
      this.isFetching = true;
      try {
        const res = await fetch(this.apiBase);
        this.complaints = await res.json();
      } catch (err) {
        console.error("Error fetching integration data:", err);
      } finally {
        setTimeout(() => (this.isFetching = false), 500);
      }
    },
    statusClass(status) {
      if (status === "Pending") return "status-pending";
      if (status === "In Progress") return "status-progress";
      return "status-resolved";
    }
  }
}
</script>
