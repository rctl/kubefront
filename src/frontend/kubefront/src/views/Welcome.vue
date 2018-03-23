<template>
  <div class="valign-wrapper placeholder">
    <SignInCard 
      class="col m6 sign-in-card z-depth-2"
      message="Please sign in to access your kubernetes cluster." 
      />
      <footer class="teal white-text z-depth-4">
        Powered by Kubefront, a kubernetes frontend.
        <span class="right">
          <a href="https://github.com/rctl/kubefront" target="_blank" class="white-text">
            Visit Kubefront on GitHub
          </a>
        </span>
      </footer>
  </div>
</template>

<script>
import SignInCard from "@/components/SignInCard.vue";

export default {
  name: "welcome",
  components: {
    SignInCard
  }
};
</script>

<style scoped lang="scss">
  .sign-in-card{
    margin: auto !important;
  }
  .placeholder{
    height: 100vh;
  }
  footer{
    position: absolute;
    bottom: 0;
    height: 50px;
    width: 100%;
    -webkit-box-shadow: 0px 3px 90px -3px rgba(0,0,0,0.31);
    -moz-box-shadow: 0px 3px 90px -3px rgba(0,0,0,0.31);
    box-shadow: 0px 3px 90px -3px rgba(0,0,0,0.31);
    padding: 15px;
  }
</style>
