<template>
  <!-- Main -->
  <main class="w-100 h-100 p-2">
    <div
      data-bs-spy="scroll"
      data-bs-target="#navbar-example"
      data-bs-offset="0"
      tabindex="0"
      class="main-contant w-100 h-100 shadow-sm scrollspy-example"
    >
      <div>
        <div class="card" style="width: 18rem">
          <div class="card-body">
            <h5 class="card-title px-3">Admin</h5>
            <div class="d-flex justify-content-around align-items-center">
              <i class="fa fa-users h4 icon-user" aria-hidden="true"></i>
              <h2>{{ userCount }} คน</h2>
            </div>
          </div>
        </div>
      </div>
      
      <table class="table my-3 shadow-sm table-bordered">
        <thead class="">
          <tr>
            <th scope="col">#</th>
            <th scope="col">UserName</th>
            <th scope="col">Email</th>
            <th scope="col">Password</th>
          </tr>
        </thead>
        <tbody v-for="user in users" :key="user.id">
          <tr>
            <th scope="row">{{ user.id }}</th>
            <td>{{ user.username }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.password }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </main>
</template>

<script>
export default {
  data() {
    return {
      users: [],
    };
  },
    computed: {
    userCount() {
      return this.users.length;
    },
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
          // this.fetchUsers();
        } else {
          console.error("Failed to fetch users");
        }
      } catch (error) {
        console.error("Error".error);
      }
    },
  },
};
</script>

<style>
.main-contant {
  background: rgb(241, 241, 241);
  border-radius: 20px;
  overflow-y: auto;
  padding: 50px;
}
.icon-user {
  color: black;
}

</style>
