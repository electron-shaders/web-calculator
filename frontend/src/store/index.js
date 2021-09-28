import { createStore } from "vuex";

const store = createStore({
  state () {
    return {
        token: null,
        origExp: "",
    }
  },
  getters: {
    getOrigExp: (state) => {
      return state.origExp
    },
  },
  mutations: {
    setOrigExp(state, newExp){
        state.origExp = newExp
    },
    appendOrigExp(state, ch){
        state.origExp += ch
    },
    updateRes(state, res){
        state.errorMsg = res.data['error-msg']
        state.correctedExp = res.data['corrected-exp']
        state.answer = res.data['answer']
    },
  },
})

export default store
