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
import auth from "./services/auth"

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

  },
  mounted() {
    this.$bus.$on(this.$auth.broadcasts.SIGNED_IN, (r) => {
      this.signedIn = true
      this.$router.push("/")
    })
    this.$bus.$on(this.$auth.broadcasts.SIGNED_OUT, (r) => {
      this.signedIn = true
      this.$router.push("auth")
    })
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