<template>
<ul>
    <li v-for="user in users" :key="user.id">
    <h1>Username {{user.username}}</h1>
  </li>
</ul>
</template>

<script>
export default {
  data() {
    return {
      users: []
    };
  },
  mounted() {
    this.fetchUsers();
  },
  methods: {
    async fetchUsers() {
      try {
        const response = await fetch("http://localhost:8080/api/users");
        if (response.ok) {
          this.users = await response.json();
        } else {
          console.error("Failed to fetch users");
        }
      } catch (error) {
        console.error("Error:", error);
      }
    }
  }
};
</script>
<style>

</style>