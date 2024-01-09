"use client"

import CustomButton from "./CustomButton"
import {useRouter} from "next/navigation"
import useStore from "@/store"

export default function QuizLauncher() {
    const router = useRouter() 
    const { resetAllQuizResult, resetCorrectNum } = useStore()
    const handleQuizStartBtn= () => {
        resetAllQuizResult()
        resetCorrectNum()
        router.push("/quiz")
    }
    return (
        <div>
            <CustomButton
                colorScheme="orange"
                variant="outline"
                onClick={handleQuizStartBtn}
            >
                今すぐ試してみる
            </CustomButton>
        </div>
    )
}