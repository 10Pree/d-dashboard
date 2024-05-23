<template>
  <!-- Main -->
  <main class="w-100 h-100 p-2 b">
    <div
      data-bs-spy="scroll"
      data-bs-target="#navbar-example"
      data-bs-offset="0"
      tabindex="0"
      class="main-contant w-100 h-100 shadow-sm scrollspy-example"
    >
      <!-- <div>
        <div class="card" style="width: 18rem">
          <div class="card-body">
            <h5 class="card-title px-3">Admin</h5>
            <div class="d-flex justify-content-around align-items-center">
              <i class="fa fa-users h4 icon-user" aria-hidden="true"></i>
              <h2>99 คน</h2>
            </div>
          </div>
        </div>
      </div> -->
      <div class=" w-100 my-3 btn_buttonH">
        <button class="btn btn-success " type="submit">เพิ่มข้อมูล</button>
      </div>
      <table class="table table-bordered table-hover">
        <thead class="">
          <tr class=" text-center">
            <th scope="col">ID</th>
            <th scope="col">UserName</th>
            <th scope="col">Last</th>
            <th scope="col">Handle</th>
            <th scope="col">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <th scope="row">1</th>
            <td>Mark</td>
            <td>Otto</td>
            <td>@mdo</td>
                        <td class=" d-flex justify-content-around align-items-center iconc"><i class="fa fa-edit btn btn-warning w-25" aria-hidden="true"></i>
            <i class="fa fa-trash btn btn-danger w-25" aria-hidden="true"></i>
            </td>
          </tr>
          <tr>
            <th scope="row">2</th>
            <td>Jacob</td>
            <td>Thornton</td>
            <td>@fat</td>
                        <td class=" d-flex justify-content-around align-items-center iconc"><i class="fa fa-edit btn btn-warning w-25" aria-hidden="true"></i>
            <i class="fa fa-trash btn btn-danger w-25" aria-hidden="true"></i>
            </td>
          </tr>
          <tr>
            <th scope="row">3</th>
            <td>Larry</td>
            <td>the Bird</td>
            <td>@twitter</td>
            <td class=" d-flex justify-content-around align-items-center iconc"><i class="fa fa-edit btn btn-warning w-25" aria-hidden="true"></i>
            <i class="fa fa-trash btn btn-danger w-25" aria-hidden="true"></i>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </main>
</template>

<script>
export default {};
</script>

<style>
.main-contant {
  /* background: rgb(241, 241, 241); */
  border-radius: 20px;
  overflow-y: auto;
  padding: 50px;
}
.icon-user {
  color: aqua;
}
.iconc i{
  color: white;
}
/* .btn_buttonH button{
  display: flex;
  justify-content: left;
} */
</style>
