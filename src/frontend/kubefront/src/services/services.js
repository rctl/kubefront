import bus from "../bus"
import api from "../api"
import router from '../router'

let service = {

    //Properties
    lastUpdate: new Date(),
    services: [],

    //Broadcasts
    broadcasts: {
        UPDATED: "SERVICE_UPDATED",
        ADDED: "SERVICE_ADDED"
    },

    //Methods
    /**
     * Force state sync
     */
    refresh() {
        return new Promise((resolve, reject) => {
            api.get("/services/")
            .then(r => {
                this.lastUpdate = new Date()
                this.services = r.data
                bus.$emit(this.broadcasts.ADDED)
                bus.$emit(this.broadcasts.UPDATED)
            }).catch(reject);
        })
    },
    /**
     * Returns a Promise which yeilds the number of services in existance
     */
    count() {
        return new Promise((resolve, reject) => {
           resolve(this.services.length)
        })
    },
    /**
     * Returns a Promise which yeilds the list of services in existance
     */
    list() {
        return new Promise((resolve, reject) => {
           resolve(this.services)
        })
    },
    /**
     * Start a deletion job. Returns a request Promise
     * @param  {string} namespace - The namespace of the resource to be deleted
     * @param  {string} name - The name of the resource to be deleted
     */
    delete(namespace, name) {
        return api.delete("/services/" + namespace + "/" + name)
    }
}

//Handle to events informing that a service has been changed
bus.$on("SERVICE_CHANGED", (entityID, data) => {
    if(data.status.replicas == 0){
        service.services = service.services.filter(x => x.metadata.name != entityID)
        bus.$emit(service.broadcasts.ADDED)
        return
    }
    let i = service.services.findIndex(x => x.metadata.name == entityID)
    data.lastUpdate = new Date()
    if(i != -1){
        service.services[i] = data
        bus.$emit(service.broadcasts.UPDATED)
    }else{
        service.services.push(data)
        bus.$emit(service.broadcasts.ADDED)
    }
})

//Export
export default service