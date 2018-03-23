<template>
  <div class="view">
    <grid-layout
            :layout="layout"
            :col-num="12"
            :row-height="30"
            :is-draggable="true"
            :is-resizable="true"
            :is-mirrored="false"
            :vertical-compact="true"
            :margin="[10, 10]"
            :use-css-transforms="true"
    >

        <grid-item class="card" v-for="item in layout"
                  :x="item.x"
                  :y="item.y"
                  :w="item.w"
                  :h="item.h"
                  :i="item.i" v-bind:key="item.i"
                  @resize="changed"
                  @move="changed"
                  @resized="changed"
                  @moved="changed">
            <div class="card-content wrapper">
              <component :is="components[item.i].component" :key="item.i"></component>
            </div>
        </grid-item>
    </grid-layout>
  </div>
</template>

<script>
import VueGridLayout from 'vue-grid-layout'
import NodeCount from '../components/NodeCount'
import NodeAllocatableCPU from '../components/NodeAllocatableCPU'
import NodeAllocatableMemory from '../components/NodeAllocatableMemory'

var GridLayout = VueGridLayout.GridLayout;
var GridItem = VueGridLayout.GridItem;
  
export default {
  name: "Dashboard",
  components: {
    GridLayout,
    GridItem,
  },
  data(){
    return{
      //Default layout 
      layout: [
          
      ],
      components: {
        "node-count": {
          component: NodeCount,
          default: {
            x: 0,
            y: 0,
            w: 3,
            h: 4,
            i: "node-count"
          },
        },
        "node-allocatable-cpu": {
          component: NodeAllocatableCPU,
          default: {
            x: 3,
            y: 0,
            w: 3,
            h: 4,
            i: "node-allocatable-cpu"
          },
        },
        "node-allocatable-memory": {
          component: NodeAllocatableMemory,
          default: {
            x: 6,
            y: 0,
            w: 3,
            h: 4,
            i: "node-allocatable-memory"
          },
        }
      }
    }
  },
  methods: {
    signOut(){
        this.$auth.signOut()
    },
    changed: function(){
      localStorage.setItem("layout", JSON.stringify(this.layout))
    },
  },
  mounted() {
    if(localStorage.getItem("layout") != undefined){
      this.layout = JSON.parse(localStorage.getItem("layout"))
    }else{
      Object.keys(this.components).forEach(k => {
        this.layout.push(this.components[k].default)
      })
    }
    this.$nodes.refresh()
  }
};
</script>

<style scoped lang="scss">
  .view{
    padding: 30px;
  }
  .component{
    height: 100%;
  }
  .wrapper{
    height: 100%;
    overflow: hidden;
  }
</style>
