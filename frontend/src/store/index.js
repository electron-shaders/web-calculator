import { createStore } from "vuex";

const store = createStore({
  state() {
    return {
      token: null,
      origExp: "",
      isLoading: false,
    }
  },
  mutations: {
    setOrigExp(state, newValue) {
      state.origExp = newValue
    },
    setLoading(state, newValue) {
      state.isLoading = newValue
    }
  },
})

export default store
