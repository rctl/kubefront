<template>
    <ul id="slide-out" class="sidenav sidenav-fixed z-depth-3">
        <li>
            <router-link to="/">
            <i class="material-icons">dns</i>
            Nodes
            </router-link>
        </li>
        <li>
            <router-link to="/pods">
            <i class="material-icons">dvr</i>
            Pods
            </router-link>
        </li>
        <li class="bottom">
            <a href="#" v-on:click="signOut">
                <i class="material-icons">exit_to_app</i>
                Sign out
            </a>
        </li>
    </ul>
</template>

<script>
export default {
  name: 'Menu',
  data() {
    return { 
        
    }
  },
  methods: {
      signOut() {
          this.$auth.signOut()
      }
  },
  mounted() {
    M.Sidenav.init(this.$el.querySelector('.sidenav'), {});
  }
}
</script>

<style scoped lang="scss">
    .bottom{
        position: absolute;
        bottom: 70px;
        width: 100%;
    }
</style>
