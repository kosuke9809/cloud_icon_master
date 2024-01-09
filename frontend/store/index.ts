import { create } from "zustand";

import type { State, QuizResult } from "@/types";



const useStore = create<State> ((set)=>({
    AllQuizResult: [],
    CorrectNum: 0,
    addQuizResult: (payload: QuizResult) => set(state => ({AllQuizResult: [...state.AllQuizResult, payload]})),
    countCorrectNum: () => set(state => ({ CorrectNum: state.CorrectNum + 1})),
    resetAllQuizResult: () => set({
        AllQuizResult: []
    }),
    resetCorrectNum: () => set({
        CorrectNum: 0,
    })
}))

export default useStore