import bus from "../bus"
import api from "../api"
import router from '../router'

let service = {

    //Properties
    lastUpdate: new Date(),
    deployments: [],

    //Broadcasts
    broadcasts: {
        UPDATED: "DEPLOYMENT_UPDATED",
        ADDED: "DEPLOYMENT_ADDED"
    },

    //Methods
    refresh() {
        return new Promise((resolve, reject) => {
            api.get("/deployments/")
            .then(r => {
                this.lastUpdate = new Date()
                this.deployments = r.data
                bus.$emit(this.broadcasts.ADDED)
                bus.$emit(this.broadcasts.UPDATED)
            }).catch(reject);
        })
    },
    count() {
        return new Promise((resolve, reject) => {
           resolve(this.deployments.length)
        })
    },
    list() {
        return new Promise((resolve, reject) => {
           resolve(this.deployments)
        })
    },
    delete(namespace, name) {
        return api.delete("/deployments/" + namespace + "/" + name)
    }
}

//Init and listeners
bus.$on("DEPLOYMENT_CHANGED", (entityID, data) => {
    if(data.status.replicas == 0){
        service.deployments = service.deployments.filter(x => x.metadata.name != entityID)
        bus.$emit(service.broadcasts.ADDED)
        return
    }
    let i = service.deployments.findIndex(x => x.metadata.name == entityID)
    data.lastUpdate = new Date()
    if(i != -1){
        service.deployments[i] = data
        bus.$emit(service.broadcasts.UPDATED)
    }else{
        service.deployments.push(data)
        bus.$emit(service.broadcasts.ADDED)
    }
})

//Export
export default service