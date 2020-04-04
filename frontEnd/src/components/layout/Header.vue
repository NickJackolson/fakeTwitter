<template>
    <header class="header">
        <h1>FakeTwitter</h1>
            <div id="nav">
                <router-link to="/login">Login</router-link> |
                <router-link to="/register">Register</router-link> |
                <router-link to="/">Home</router-link> |
                <router-link to="/about">About</router-link>
            </div>
    </header>
</template>

<script>
export default {
    name:"Header",
}
</script>

<style scoped>
.header{
    background: rgb(92, 92, 92);
    color:#fff;
    text-align: center;
    padding: 10px;
}
.header a {
    color:#fff;
    padding-right: 5px;
    text-decoration: none;
}
.header h1 {
    color: #DAA520;
    padding-right: 5px;
    text-decoration: none;
}
</style>
