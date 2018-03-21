<template>
  <div class="grey lighten-5">
    <Loader :active="loading" />
    <div id="content" class="grey lighten-5">
      <router-view />
    </div>
  </div>
</template>

<script>
import Loader from "@/components/Loader.vue";

export default {
  name: "app",
  components: {
    Loader
  },
  data() {
    return {
      signedIn: false,
      loading: false,
    }
  },
  methods: {
    onSignIn(){
      //Fetch user profile and present menu
      this.$router.push("/")
    },
    onSignOut(){
      //Hide menu and clear profile
      this.$router.push("auth")
      localStorage.removeItem("token")
      this.loading = false;
    }
  },
  mounted() {
    //Check if signed in when app is launched
    if(localStorage.getItem("token") != undefined){
      this.onSignIn()
    }else{
      this.onSignOut()
    }
    //Register listener for sign in event
    this.$bus.$on("signedIn", this.onSignIn)
    this.$bus.$on("signOut", this.onSignOut)
    this.$bus.$on("loading", () => { this.loading = true })
    this.$bus.$on("done", () => { this.loading = false })
  }
};
</script>

<style lang="scss">
  #content{
    height: 100vh;
  }
</style>

<!-- <div id="nav">
      <router-link to="/">Home</router-link>
      </div> -->
