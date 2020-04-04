<template>
  <div class="article-item">
      <p>-{{article.author}}</p>
      <p>{{article.pdate}}</p>
      <p>{{article.title}}</p>
      <p>{{article.content}}</p>
      <hr>
  </div>
</template>

<script>
export default {
  name: 'ArticleItem',
  props: ["article"],
  methods:{
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
