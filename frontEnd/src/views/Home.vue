<template>
  <div class="home">
    <button @click="logout" name="logout">Logout</button>
    <AddArticle v-on:add-article="addArticle"/>
    <hr>
    <Articles v-bind:articles="articles"/>
  </div>
</template>



<script>
import Articles from '@/components/Articles.vue'
import AddArticle from '@/components/AddArticle.vue'
import axios from 'axios'

export default {
  name: 'Home',
  components: {
    Articles,
    AddArticle
  },
  data(){
    return{
      articles: []
    }
  },
  methods:{
    logout(){
      localStorage.removeItem('user');
      localStorage.removeItem('token');
      this.$router.push('/login');
    },
    addArticle(newArticle){
      const {title,content,author} = newArticle;
      axios.post('http://localhost:8081/articles',{
        title,
        content,
        author
      })
      .then(res=>this.articles.unshift(res.data))
      .catch(err=>console.log(err));
    }
  },
  created(){
    axios.get('http://localhost:8081/articles')
    .then(res => {
      this.articles = res.data
      this.articles.ptime = res.data.ptime.to_String()
      console.log(this.articles)})
    .catch(err => console.log(err));
  }
}
</script>

<style scoped>
</style>