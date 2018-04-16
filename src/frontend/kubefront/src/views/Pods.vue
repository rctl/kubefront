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

var GridLayout = VueGridLayout.GridLayout;
var GridItem = VueGridLayout.GridItem;

export default {
  name: "pods",
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
            w: 4,
            h: 5,
            i: "node-count"
          },
        }
      }
    }
  },
  methods:{
    changed: function(){
      localStorage.setItem("layout-pods", JSON.stringify(this.layout))
    },
  },
  mounted(){
    this.$upstream.subscribe("PODS");
    this.$nodes.refresh()
    /*this.$bus.$on(this.$pods.UPDATED, _ => {
        this.lastUpdate = new Date()
    })*/
    let stored = JSON.parse(localStorage.getItem("layout"));
    if(localStorage.getItem("layout") != undefined && Object.keys(this.components).length == stored.length){
      this.layout = stored;
    }else{
      Object.keys(this.components).forEach(k => {
        this.layout.push(this.components[k].default)
      })
    }
  },
  destroyed(){
      this.$upstream.unsubscribe("PODS");
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
