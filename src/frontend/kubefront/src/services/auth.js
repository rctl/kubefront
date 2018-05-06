import bus from "../bus"
import api from "../api"
import router from '../router'

//States
let states = {
    signedIn: false,
    token: "",
}

//Broadcasts
let broadcasts = {
    SIGNED_IN: "SIGNED_IN",
    SIGNED_OUT: "SIGNED_OUT"
}

//Init
if(localStorage.getItem("token") != undefined){
    states.signedIn = true
    states.token = localStorage.getItem("token")
    api.defaults.headers.common['token'] = localStorage.getItem("token");
    bus.$emit(broadcasts.SIGNED_IN)
}else{
    states.signedIn = false
    bus.$emit(broadcasts.SIGNED_OUT)
}

//Export
export default {

    //Properties
    states: states,
    broadcasts: broadcasts,

    //Methods
    signIn(username, password) {
        return new Promise((resolve, reject) => {
            var formData = new FormData();
            formData.set("username", username)
            formData.set("password", password)
            api.post("/auth/", formData)
            .then(r => {
                api.defaults.headers.common['token'] = r.data.token;
                states.signedIn = true
                states.token = r.data.token
                localStorage.setItem("token", r.data.token)
                bus.$emit(broadcasts.SIGNED_IN)
                resolve(r)
            }).catch(reject);
        })
    },
    signOut() {
        states.signedIn = false
        localStorage.removeItem("token")
        bus.$emit(broadcasts.SIGNED_OUT)
    },
}