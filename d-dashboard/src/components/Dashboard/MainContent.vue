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
          <div class="card" style="width: 18rem" >
            <div class="card-body">
              <h5 class="card-title px-3 ">Admin</h5>
              <div class="d-flex justify-content-around align-items-center">
                <i class="fa fa-users h4 icon-user" aria-hidden="true"></i>
                <h2>99 คน</h2>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
</template>

<script>
export default {

}
</script>

<style>
.main-contant {
  background: rgb(241, 241, 241);
  border-radius: 20px;
  overflow-y: auto;
  padding: 50px;
}
.icon-user {
  color: aqua;
}
</style>