<script setup>
import MainContent from "./MainContent.vue";

</script>
<template>
  <div class="w-100 d-flex gap-2" style="height: 100vh">
    <aside id="navbar-example">
      <div class="sidebar shadow-sm">
        <div class="d-flex justify-content-center align-items-center logo">
          <img
            src="../../assets/images/logo_1.png"
            style="width: 80px"
            alt=""
          />
        </div>
        <ul class="sidebar-nav list-unstyled">
          <li class="nav-item">
            <a class="btn btn-outline-primary border-0" href="#">
              <i class="fa fa-dashboard px-2" aria-hidden="true"></i>
              <span>Dashboard</span>
            </a>
          </li>
          <li class="nav-item">
            <a class="btn btn-outline-primary border-0" href="#">
              <i class="fa-solid fa-cart-shopping px-2"></i>
              <span>HOME</span>
            </a>
          </li>
          <li class="nav-item">
            <a class="btn btn-outline-primary border-0" href="#">
              <i class="fa-solid fa-users px-2"></i>
              <span>Admin</span>
            </a>
          </li>
          <li class="nav-item last-menu-out">
            <a
              class="btn btn-outline-primary border-0"
              href="./Login/login.vue"
            >
              <i class="fa fa-sign-out px-2" aria-hidden="true"></i>
              <span>ออกจากระบบ</span>
            </a>
          </li>
        </ul>
      </div>
    </aside>

    <!-- Main -->
    <MainContent/>
  </div>
</template>

<style>
.sidebar {
  width: 250px;
  height: 100vh;
  text-decoration: none;
}
.sidebar-nav {
  width: 95%;
  padding-left: 1.5rem;
}
.nav-item {
  margin-top: 18px;
}
.sidebar a {
  padding: 0.625rem 2.625rem;
  display: block;
  text-decoration: none;
  border-radius: 20px;
  text-align: left;
  font-weight: 500;
}
.last-menu-out {
  bottom: 50px;
  width: 100%;
  color: brown;
  font-size: 1.3rem;
  margin-top: 100%;
}

.logo {
  height: 100px;
}
</style>
