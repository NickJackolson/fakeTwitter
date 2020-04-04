<template>
    <div id="app">
      <Header />
      <router-view/>
    </div>
</template>

<script>
import Header from './components/layout/Header'
export default {
  components:{
    Header
  }
}
</script>

<style>
body{
  background: #91A196;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  text-align: center;
}

#nav {
  padding: 10px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #DAA520;
}

#nav a:hover {
  font-weight: bold;
  color: #bdbdbd;
}
</style>
