<template>
  <div class="w-100 d-flex justify-content-center align-items-center" style="height: 100vh">
    <form class=" container-fluid  w-25 border border-1 rounded shadow-sm align-items-center p-6" style="height: 400px;" action="">
      <div class=" d-flex justify-content-center w-100 mt-1" key="">
        <img  src="../../assets/images/logo_1.png" style=" width: 100px;" alt="">
      </div>
      <div class=" m-2">
        <label class=" form-label d-block" for="">Username</label>
        <input class=" form-control" type="text" />
      </div>
      <div class=" m-2">
        <label class=" form-label d-block" for="">Password</label>
        <input class=" form-control" type="password" name="password" />
      </div>
      <div class=" m-2">
        <button class=" btn btn-primary w-100" type="submit">เข้าสู่ระบบ</button>
      </div>
      <div class="d-flex justify-content-between align-items-center m-2">
        <div>
          <input type="checkbox" />
          <small>จำรหัสไว้</small>
        </div>
        <div>
          <a href="#">ลืมรหัส</a>
        </div>
      </div>
      <div class=" m-2 mt-4 ">
        <a class=" btn btn-outline-primary w-100" >สมัคร สมาชิก</a>
      </div>
    </form>
  </div>
</template>

<script></script>

<style>
img{
  background-position: center;
  background-size: cover;
  object-fit: cover;

}
</style>
