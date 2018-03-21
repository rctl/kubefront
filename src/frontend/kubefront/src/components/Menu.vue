<template>
    <div id="menu">
        
    </div>
</template>

<script>
export default {
  name: 'Menu',
  props: {
      visible: {
          default: true,
          type: Boolean
      }
  },
  watch: {
    visible () {
        if (this.visible && !this.visible) this.show()
        else if (!this.visible && this.visible) this.hide()
    },
  },
  data() {
    return { 
        isVisible: false,
    }
  },
  methods: {
    hide () {
        this.isVisible = false
        menu.style.display = "none";
    },
    show () {
        this.isVisible = true
        menu.style.display = "block";
    },
  },
  mounted() {
      if(this.visible){
        this.show()
      }else{
        this.hide()
      } 
  }
}
</script>

<style scoped lang="scss">

</style>
