<template>
    <div class="wrapper">
        <div class="center">Deployments <span class="right"><a href="#" class="grey-text tooltipped" :data-tooltip="lastUpdate|moment('from', 'now')"><i class="material-icons">info</i></a></span></div>
        <div class="wrapper scroll">
            <ul class="collapsible full-width">
                <li v-for="(d, _) in deployments" :key="_">
                    <div class="collapsible-header">
                        <i class="material-icons left" v-if="!extended && !d.job">cloud_circle</i>
                        <div v-if="d.job" class="preloader-wrapper small active left tooltipped" data-tooltip="There is a job running on this deployment" style="height:20px; width: 20px; margin-right:10px;">
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
                        <b>{{d.metadata.name}}</b>
                        <span class="badge">
                            <span v-if="extended">{{d.status.readyReplicas}}/{{d.status.replicas}}</span>
                            <a href="#" class="right">
                                <i v-if="d.status.readyReplicas==d.status.replicas && d.status.replicas != 0" class="material-icons green-text">brightness_1</i>
                                <i v-if="d.status.readyReplicas < d.status.replicas" class="material-icons red-text">brightness_1</i>
                            </a>
                        </span>
                    </div>
                    <div class="collapsible-body">
                        <div class="card view">
                            <PodList :labelSelector="d.spec.selector.matchLabels" extended></PodList>
                        </div>
                        <div class="btn-row" v-if="!d.job">
                            <a class='btn-flat right red-text waves-effect waves-red' href='javascript:void();' @click="deleteDeployment(d)" v-if="extended">
                                <i class="material-icons left">delete</i> Delete Deployment
                            </a>
                        </div>
                    </div>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
import PodList from '../components/PodList'

export default {
  name: 'DeploymentList',
  components: {
    PodList
  },
  data() {
    return { 
       deployments: [],
       lastUpdate: new Date()
    }
  },
  props: {
      extended: {
          default: false,
          type: Boolean
      }
  },
  methods: {
      refresh(){
        this.$deployments.list().then(d => {
            this.deployments = []
            if(d !== undefined){
                d.forEach(x => {
                    this.$workers.get("deployment", x.metadata.namespace, x.metadata.name).then(w => {
                        x.job = w.length != 0;
                    })
                    this.deployments.push(x)
                })
            }
        }).then(_ => {
            this.lastUpdate = new Date()
        })
      },
      deleteDeployment(d){
          this.$deployments.delete(d.metadata.namespace, d.metadata.name).then(this.refresh).catch(() => {
              p.job = true;
              M.toast({html: 'Deployment could not be deleted'})
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
    this.$bus.$on(this.$deployments.broadcasts.UPDATED, _ => {
      this.refresh()
    })
    this.$bus.$on(this.$deployments.broadcasts.ADDED, _ => {
      this.refresh()
    })
    this.$el.querySelectorAll('.tooltipped').forEach(e => {
        M.Tooltip.init(e, {});
    });
    this.$bus.$on("JOB_STARTED", (id, data) => {
      this.refresh()
    })
    this.$bus.$on("JOB_COMPLETED", (id, data) => {
      this.$deployments.refresh().then(this.refresh)
    })
    this.$bus.$on("JOB_FAILED", (id, data) => {
      this.$deployments.refresh().then(this.refresh)
    })
  }
}
</script>

<style scoped lang="scss">
    .wrapper{
        width: 100%;
        height: 100%;
    }
    .scroll{
        overflow:auto;
    }
    .content{
        margin: auto;
    }
    .full-width{
        width: 100%;
    }
    .view{
        padding: 30px;
    }
    .btn-row{
        height: 20px;
    }
</style>
