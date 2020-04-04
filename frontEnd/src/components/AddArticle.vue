<template>
<div class="postArticle">
    <form class="box" @submit="addArticle">
        <h3>Create A Post</h3>
        <input type="text" v-model='title' name="title" placeholder="Title">
        <input type="text" v-model='content' name="content" placeholder="What's on your mind?">
        <input type="submit" value="Submit">
    </form>
</div>
 
</template>

<script>

export default {
    name:"AddArticle",
    data(){
        return{
            title: '',
            content:'',
            author:'',
        }
    },
    methods: {
      addArticle (){
        const user = JSON.parse(localStorage.getItem('user'))
        const newArticle ={
          title: this.title,
          content:this.content,
          author: user
        }
        this.$emit('add-article',newArticle);
        this.newArticle = NaN;
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
