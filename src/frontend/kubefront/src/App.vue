<template>
  <div>
    <Loader :active="loading" />
    <div v-bind:class="{ menu: signedIn }">
      <Menu v-if="signedIn" />
      <div id="content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
//Components
import Loader from "@/components/Loader.vue";
import Menu from "@/components/Menu.vue";

export default {
  name: "app",
  components: {
    Loader,
    Menu
  },
  data() {
    return {
      signedIn: false,
      loading: false,
    }
  },
  methods: {

  },
  updated(){
    if(!this.$auth.states.signedIn){
      this.$router.push("auth")
    }
  },
  mounted() {
    

    //Authentication states
    this.$bus.$on(this.$auth.broadcasts.SIGNED_IN, (r) => {
      this.signedIn = true
      this.$router.push("/")
    })
    this.$bus.$on(this.$auth.broadcasts.SIGNED_OUT, (r) => {
      this.signedIn = false
      this.$router.push("auth")
    })
    if(!this.$auth.states.signedIn){
      this.$router.push("auth")
    }
    this.signedIn = this.$auth.states.signedIn;

    //Loading spinner
    this.$bus.$on("loading", () => { this.loading = true })
    this.$bus.$on("done", () => { this.loading = false })

    //App global data load
    this.$workers.update();
    this.$upstream.connect()
  }
};
</script>

<style lang="scss">
  #content{
    height: 100vh;
  }
  .menu{
    padding-left: 300px;
    width: 100%;
  }
  @media only screen and (max-width: 992px) {
    .menu{
      padding-left: 0px;
      width: 100%;
    }
  }
</style>