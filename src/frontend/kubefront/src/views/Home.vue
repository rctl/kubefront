<template>
  <div>
    Home sweet home!
    <button v-on:click="signOut">Sign out</button>
  </div>
</template>

<script>
export default {
  name: "home",
  components: {
    
  },
  methods: {
    signOut(){
        this.$auth.signOut()
    }
  }
};
</script>

<style scoped lang="scss">

</style>
