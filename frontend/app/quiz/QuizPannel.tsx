"use client"

import React, { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import useStore from "@/store"
import type { QuizProps , CustomButtonProps, QuizResult} from "@/types"
import CustomButton from "../components/CustomButton"
import { Image, Text } from "@chakra-ui/react"
import { ArrowForwardIcon } from "@chakra-ui/icons"



export default function QuizPannel(quizProps: QuizProps) {
    const [correctAnswer, setCorrectAnswer] = useState<string>(quizProps.answer)
    const [isCorrect, setIsCorrect] = useState<boolean>(false)
    const [resultMessage, setResultMessage] = useState<String>("")
    const [answered, setAnswerd] = useState<boolean>(false)
    const router = useRouter()
    const { AllQuizResult, addQuizResult, countCorrectNum} = useStore()
    
    const updateQuizResultStore = (quizRuslt: QuizResult) => {
        addQuizResult({
            ...quizRuslt
        })
    }

    const handleAnswerSelect: React.MouseEventHandler<HTMLButtonElement> = (event) => {
        event.preventDefault()
        const selectAnswer = event.currentTarget.textContent;
        setAnswerd(true)
        
        console.log(selectAnswer)
        console.log(correctAnswer)
        if (correctAnswer === selectAnswer) {
            setResultMessage("正解です。")
            countCorrectNum()
            setIsCorrect(true)
        } else {
            setResultMessage("ざんねん、不正解です。")
        }

        const quizRuslt: QuizResult = {
            answer: correctAnswer,
            select : selectAnswer,
            result: isCorrect
        }
        updateQuizResultStore(quizRuslt)
    }

    const handleNextQuiz: React.MouseEventHandler<HTMLButtonElement> = () => {
        if (AllQuizResult.length >= 5) {
            router.push("/trial-result")
        }  
        router.refresh()
    }

    // クイズ回答用のボタン
    const quizItemProps: CustomButtonProps = {
        colorScheme: "orange",
        variant: "outline",
        isDisabled: answered,
        onClick: handleAnswerSelect,
        
    }

    const nextQuizButtomProps: CustomButtonProps = {
        colorScheme: "blackAlpha",
        variant: "ghost",
        rightIcon: <ArrowForwardIcon/>,
        onClick: handleNextQuiz,
    }

    useEffect(() => {
        setCorrectAnswer(quizProps.answer)
        setAnswerd(false)
        console.log("更新")
        console.log(AllQuizResult)
    }, [quizProps.answer])


    return (
        <div className="p-8">
            <Text fontSize="2xl"className="text-center font-bold text-gray-700" mb={4} >該当するサービス名を答えよ</Text>
            
            <div className="flex justify-center">
                <Image src={quizProps.iconImgUrl} alt="aws icon" className="w-64 h-64" objectFit="cover"/>
            </div>
            <div className=" justify-center p-8">
                <div className="grid grid-cols-2 gap-4 mx-auto max-w-3xl">
                    {quizProps.quizData?.map((answer) => (
                        <CustomButton key={answer.id} {...quizItemProps}>{answer.name}</CustomButton>
                    ))}
                </div>
            </div>
            <div>
                {answered ? (
                    <div className="text-center">
                        <Text>{resultMessage}正解は{correctAnswer}です。</Text>
                        <CustomButton {...nextQuizButtomProps}>次にすすむ</CustomButton>
                    </div>
                ) : (
                    null
                )}
            </div>
            
        </div>
    )
}
