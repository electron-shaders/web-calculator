import { createStore } from "vuex";

const store = createStore({
  state() {
    return {
      token: null,
      origExp: "",
      isLoading: false,
      ansHistory: [],
    }
  },
  mutations: {
    setOrigExp(state, newValue) {
      state.origExp = newValue
    },
    setLoading(state, newValue) {
      state.isLoading = newValue
    },
    updateAnsHistory(state, newValue) {
      state.ansHistory.unshift({
        correctedExp: newValue.correctedExp,
        answer: newValue.answer,
        index: 0,
      });
      for (let i = 1; i < state.ansHistory.length; i++) {
        state.ansHistory[i].index = i;
      }
    },
    deleteAnsHistory(state, index) {
      state.ansHistory.splice(index, 1);
      for (let i = 0; i < state.ansHistory.length; i++) {
        state.ansHistory[i].index = i;
      }
    }
  },
})

export default store
