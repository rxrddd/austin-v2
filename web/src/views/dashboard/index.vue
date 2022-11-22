<template>
  <div class="dashboard-container">
    {{ currentRole }}
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      currentRole: 'adminDashboard'
    }
  },
  created() {
    
  }
}
</script>
