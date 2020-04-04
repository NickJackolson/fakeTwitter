<template>
  <div>
    <form class="box" @submit="register">
      <h2>REGISTER</h2>
      <input type="text" v-model='username' name="username" placeholder="Choose a username">
      <input type="password" v-model='password' name="password" placeholder="Choose a password">
      <input type="email" v-model='email' name="email" placeholder="Type in an Email">
      <input type="submit" value="Register">
    </form>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  
  data(){
    return{
      username:'',
      password:'',
      email:''
    }
  },
  methods: {
      register (e){
        e.preventDefault();
        const newUser ={
          username: this.username,
          password: this.password,
          email:this.email,
        }
        axios.post('http://localhost:8081/register',newUser)
        .then(res => {
          console.log(res.data)
          this.$router.push('/login')
          })
        .catch(err => console.log(err));
        this.newUser = NaN;
      }
    }
}
</script>script

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
::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
  color: #DAA520;
  opacity: 0.4; /* Firefox */
}
.box input[type ="text"], .box input[type = "password"], .box input[type = "email"]{
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