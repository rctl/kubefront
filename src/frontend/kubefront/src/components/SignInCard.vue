<template>
  <div class="card">
    <div class="card-content">
      <span class="card-title" v-if="title.length > 0">{{ title }}</span>
      <p class="card-message grey-text text-darken-1">{{ message }}</p>
      <form v-on:submit.prevent="signIn">
      <div class="input-field">
        <input id="username" type="text" v-on:input="change" v-model="username" autocomplete="off">
        <label for="username">Username</label>
      </div>
      <div class="input-field">
        <input id="password" type="password" v-on:input="change" v-model="password">
        <label for="password">Password</label>
      </div>
      <div class="btn-row">
        <input type="submit" class="btn right" id="signInButton" value="Sign In" />
      </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SignInCard',
  props: {
    title: {
      default: "",
      type: String
    },
    message: {
      default: "",
      type: String
    },
  },
  data() {
    return { 
      username: "",
      password: ""
    }
  },
  methods: {
    signIn() {
      //Package data for sign in
      this.$bus.$emit('loading')
      //Send request
      this.$auth.signIn(this.username, this.password)
      .then(() => {
        this.$bus.$emit('done')
      })
      .catch(r => {
        //Notify failure to user
        this.$bus.$emit('done')
        if(r.response){
          M.toast({html: r.response.data})
          this.password = ""
        }else{
          //No response is a server error
          M.toast({html: "Could not sign in due to server error."})
        }
      });
    },
    change() {
      //Show sign in button when there is data in both fields
      if(this.password.length > 0 && this.username.length > 0){
        signInButton.style.height = "35px";
      }else{
        signInButton.style.height = "0px";
      }
    }
  },
  mounted() {
    M.updateTextFields();
  }
}
</script>

<style scoped lang="scss">
  .card-message {
    margin-bottom: 20px !important;
    text-align: center;
  }
  .btn-row{
    height: 30px;
  }
  #signInButton{
    height: 0px;
    transition: 0.3s ease-in-out;
  }
</style>
