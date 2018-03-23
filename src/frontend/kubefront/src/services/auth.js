import bus from "../bus"
import api from "../api"
import router from '../router'

//States
let states = {
    signedIn: false
}

//Broadcasts
let broadcasts = {
    SIGNED_IN: "SIGNED_IN",
    SIGNED_OUT: "SIGNED_OUT"
}

//Init
if(localStorage.getItem("token") != undefined){
    states.signedIn = true
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
        var formData = new FormData();
        formData.set("username", username)
        formData.set("password", password)
        return api
        .post("/auth/", formData)
        .then(r => {
            states.signedIn = true
            localStorage.setItem("token", r.data.token)
            bus.$emit(broadcasts.SIGNED_IN)
            return r
        });
    },

    signOut() {
        states.signedIn = false
        localStorage.removeItem("token")
        bus.$emit(broadcasts.SIGNED_OUT)
    },
}