<template>
    <div class="wrapper">
        <div class="center">Services <span class="right"><a href="#" class="grey-text tooltipped" :data-tooltip="lastUpdate|moment('from', 'now')"><i class="material-icons">info</i></a></span></div>
        <div class="wrapper scroll">
            <div class="collection-item" v-for="(l, _) in services" :key="_">
                <b class="namespace">{{l.namespace}}</b>
                <ul class="collapsible full-width">
                    <li class="item" v-for="(s, _) in l.services" :key="_">
                        <div class="collapsible-header">
                            {{s.metadata.name}}
                        </div>
                        <a class='dropdown-trigger right grey-text extension' href='#' v-if="!s.job && extended" :data-target='s.metadata.name'>
                            <i class="material-icons">more_vert</i>
                        </a>
                        <div v-if="s.job" class="preloader-wrapper small active right tooltipped extension" data-tooltip="There is a job running on this service" style="height:20px; width: 20px; margin-right:10px;">
                            <div class="spinner-layer spinner-green-only">
                            <div class="circle-clipper left">
                                <div class="circle"></div>
                            </div><div class="gap-patch">
                                <div class="circle"></div>
                            </div><div class="circle-clipper right">
                                <div class="circle"></div>
                            </div>
                            </div>
                        </div>
                        <ul :id='s.metadata.name' class='dropdown-content dropper'>
                            <li><a href="javascript:void();" class="red-text" @click="deleteService(s)"><i class="material-icons">delete</i> Delete Service</a></li>
                        </ul>
                        <div class="collapsible-body">
                            <b>Cluster IP:</b> {{s.spec.clusterIP}}<br>
                            <div class="card view">
                                <PodList :labelSelector="s.spec.selector" extended></PodList>
                            </div>
                        </div>
                    </li>
                </ul>  
            </div>
        </div>
    </div>
</template>

<script>
import PodList from '../components/PodList'
export default {
  name: 'ServiceList',
  data() {
    return { 
       services: [],
       lastUpdate: new Date()
    }
  },
  components: {
    PodList
  },
  props: {
      extended: {
          default: false,
          type: Boolean
      }
  },
  methods: {
      deleteService(s){
          s.job = true;
          this.$services.delete(s.metadata.namespace, s.metadata.name).then(this.refresh).catch(() => {
              s.job = false;
              M.toast({html: 'Service could not be deleted'})
          })
      },
      refresh(){
        this.$services.list().then(d =>{
            this.services = []
            let namespaces = {}
            d.forEach(x => {
                if(namespaces[x.metadata.namespace] === undefined){
                    namespaces[x.metadata.namespace] = []
                }
                this.$workers.get("service", x.metadata.namespace, x.metadata.name).then(w => {
                    x.job = w.length != 0;
                })
                namespaces[x.metadata.namespace].push(x);
            })
            Object.keys(namespaces).forEach(x => {
                this.services.push({
                    namespace: x,
                    services: namespaces[x]
                })
            });
        }).then(_ => {
            this.lastUpdate = new Date()
        })
      }
  },
  updated() {
    this.$el.querySelectorAll('.tooltipped').forEach(e => {
        if(!e.getAttribute("init")){
            e.setAttribute("init", true)
            M.Tooltip.init(e, {});
        }
    });
    this.$el.querySelectorAll('.dropdown-trigger').forEach(e => {
        if(!e.getAttribute("init")){
            e.setAttribute("init", true)
            M.Dropdown.init(e, {
                constrainWidth: false,
            });
        }
    });
    if(this.extended){
        this.$el.querySelectorAll('.collapsible').forEach(e => {
            if(!e.getAttribute("init")){
                e.setAttribute("init", true)
                M.Collapsible.init(e, {});
            }
        });
    }
  },
  mounted() {
    this.refresh()
    this.$bus.$on(this.$services.broadcasts.UPDATED, _ => {
      this.refresh()
    })
    this.$bus.$on(this.$services.broadcasts.ADDED, _ => {
      this.refresh()
    })
    this.$bus.$on("JOB_STARTED", (id, data) => {
      this.refresh()
    })
    this.$bus.$on("JOB_COMPLETED", (id, data) => {
      this.$services.refresh().then(this.refresh)
    })
    this.$bus.$on("JOB_FAILED", (id, data) => {
      this.$services.refresh().then(this.refresh)
    })
    this.$el.querySelectorAll('.tooltipped').forEach(e => {
        M.Tooltip.init(e, {});
    });
  }
}
</script>

<style scoped lang="scss">
    .wrapper{
        width: 100%;
        height: 100%;
        padding-bottom: 15px;
    }
    .scroll{
        overflow:auto;
    }
    .content{
        margin: auto;
    }
    .full-width{
        width: 100%;
        overflow: visible;
    }
    .indicator{
        margin-top: 8px;
    }
    .group{
        padding-right: 20px;
        padding-top: 20px;
        overflow: visible;
    }
    .dropper{
        overflow: visible;
    }
    .item{
        min-height: 30px;
    }
    .pods{
        margin-top: 10px;
    }
    .details{
        margin-left: 30px;
    }
    .extension{
        top:-40px; 
        position:relative;
    }
    .view{
        padding: 30px;
    }
</style>

