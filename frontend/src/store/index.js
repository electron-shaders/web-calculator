import { createStore } from "vuex";

const store = createStore({
  state () {
    return {
        token: null,
        /*
        origExp: "",
        errorMsg: "",
        correctedExp: "",
        answer: NaN,
        */
    }
  },
  mutations: {
    /*
    updateOrigExp(state, newExp){
        state.origExp = newExp
    },
    updateRes(state, res){
        state.errorMsg = res.data['error-msg']
        state.correctedExp = res.data['corrected-exp']
        state.answer = res.data['answer']
    }
    */
  },
  actions: {

  },
})

export default store
