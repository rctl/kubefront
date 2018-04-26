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
import NodeList from '../components/NodeList'
import PodCount from '../components/PodCount'
import PodList from '../components/PodList'
import DeploymentList from '../components/DeploymentList';
import ServiceList from '../components/ServiceList'

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
            w: 2,
            h: 4,
            i: "node-count"
          },
        },
        "node-allocatable-cpu": {
          component: NodeAllocatableCPU,
          default: {
            x: 4,
            y: 0,
            w: 2,
            h: 4,
            i: "node-allocatable-cpu"
          },
        },
        "node-allocatable-memory": {
          component: NodeAllocatableMemory,
          default: {
            x: 0,
            y: 4,
            w: 2,
            h: 5,
            i: "node-allocatable-memory"
          },
        },
        "node-list": {
          component: NodeList,
          default: {
            x: 2,
            y: 4,
            w: 4,
            h: 5,
            i: "node-list"
          },
        },
        "pod-count": {
          component: PodCount,
          default: {
            x: 2,
            y: 0,
            w: 2,
            h: 4,
            i: "pod-count"
          },
        },
        "pod-list": {
          component: PodList,
          default: {
            x: 0,
            y: 9,
            w: 6,
            h: 19,
            i: "pod-list"
          },
        },
        "deployment-list": {
          component: DeploymentList,
          default: {
            x: 8,
            y: 9,
            w: 4,
            h: 28,
            i: "deployment-list"
          },
        },
        "service-list": {
          component: ServiceList,
          default: {
            x: 6,
            y: 0,
            w: 2,
            h: 28,
            i: "service-list"
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
    let stored = JSON.parse(localStorage.getItem("layout"));
    if(localStorage.getItem("layout") != undefined && Object.keys(this.components).length == stored.length){
      this.layout = stored;
    }else{
      Object.keys(this.components).forEach(k => {
        this.layout.push(this.components[k].default)
      })
    }
    this.$upstream.subscribe("NODES");
    this.$upstream.subscribe("PODS");
    this.$upstream.subscribe("DEPLOYMENTS");
    this.$upstream.subscribe("SERVICES");
    this.$nodes.refresh()
    this.$pods.refresh()
    this.$deployments.refresh()
    this.$services.refresh()
  },
  updated(){

  },
  destroyed(){
    this.$upstream.unsubscribe("NODES");
    this.$upstream.unsubscribe("DEPLOYMENTS");
    this.$upstream.unsubscribe("PODS");
    this.$upstream.unsubscribe("SERVICES");
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
