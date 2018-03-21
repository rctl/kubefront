<template>
    <div class="wrapper" id="loader">
        <div class="overlay"></div>
        <div class="wrapper-outer valign-wrapper">
            <div class="card white z-depth-0">
                <div class="card-overlay">
                    <div class="wrapper-inner valign-wrapper">
                        <div class="preloader-wrapper active">
                            <div class="spinner-layer spinner-red-only">
                            <div class="circle-clipper left">
                                <div class="circle"></div>
                            </div><div class="gap-patch">
                                <div class="circle"></div>
                            </div><div class="circle-clipper right">
                                <div class="circle"></div>
                            </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'Loader',
  props: {
      active: {
          default: false,
          type: Boolean
      }
  },
  watch: {
    active () {
        if (this.active && !this.visible) this.show()
        else if (!this.active && this.visible) this.hide()
    },
  },
  data() {
    return { 
        visible: false,
    }
  },
  methods: {
    hide () {
        this.visible = false
        loader.style.visibility = "hidden";
    },
    show () {
        this.visible = true
        loader.style.visibility = "visible";
    },
  },
  mounted() {
      if(this.active) this.show()
  }
}
</script>

<style scoped lang="scss">
    #loader{
        visibility: hidden;
    }
    .wrapper{
        width: 100vw;
        height: 100vh;
        position: absolute;
    }
    .overlay{
        width: 100vw;
        height: 100vh;
        position: absolute;
        top: 0;
        background-color: white;
        opacity: 1;
        z-index: 999999990 !important;
    }
    .wrapper-outer{
        margin: auto;
        width: 100%;
        height: 100%;
    }
    .wrapper-inner{
        margin: auto;
        width: 50px;
        height: 100%;
    }
    .card-overlay{
        width: 50vh;
        height: 50vh;
        max-width: 200px;
        min-width: 80px;
        max-height: 200px;
        min-height: 80px;
        z-index: 999999999 !important;
    }
    .card{
       z-index: 9999999999 !important;
       margin: auto !important;
    }
</style>
