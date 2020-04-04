<template>
  <div>
    <form class="box">
      <h2>LOGIN</h2>
      <input type="email" placeholder="E-mail" v-model='email' >
      <input type="password" placeholder="Password" v-model='password'>
      <p>{{error}}</p>
      <input @click="login" type="submit" value ="Login"/>
    </form>
  </div>
</template>


<script>
import axios from 'axios'
export default {
  data(){
    return{
      email:'',
      password:'',
      error:''
    }
  },
  methods: {

    login(e){
      e.preventDefault();
      const User ={
        email: this.email,
        password: this.password,
      }
      axios.post('http://localhost:8081/login',User)
      .then(res => {
        console.log(res.data);
        localStorage.setItem('user', JSON.stringify(res.data.Username));
        localStorage.setItem('token', JSON.stringify(res.data.Token));
        this.$router.push('/');
        },err =>{
          console.log(err.response)
          this.error=err.response.data
        })
    }
  }
}
</script>

<style scoped>
.box{
  width: 300px;
  padding: 40px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%,-50%);
  background: #636363;
  text-align: center;
}
.box h2{
  color: #DAA520;
}
.box p{
  color: #DAA520;
}
::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
  color: #DAA520;
  opacity: 0.4; /* Firefox */
}
.box input[type ="email"], .box input[type = "password"]{
  border:0;
  background: #5c5c5b;
  display: block;
  margin: 20px auto;
  text-align: center;
  border: 2px solid #3f3f3f;
  padding: 14px 10px;
  width: 200px;
  outline: none;
  color: #ffffff;
  border-radius: 24px;
}
.box input[type = "submit"]{
    border:0;
  background: #5c5c5b;
  display: block;
  margin: 20px auto;
  text-align: center;
  border: 2px solid #3f3f3f;
  padding: 14px 10px;
  width: 200px;
  outline: none;
  color: #DAA520;
  border-radius: 24px;
  cursor: pointer;
}
.box input[type = "submit"]:hover{
   color: #000000;
   background: #DAA520;
}
</style>